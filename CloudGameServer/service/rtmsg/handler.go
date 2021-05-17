package rtmsg

import (
	"CloudGameServer/db"
	"CloudGameServer/device"
	"CloudGameServer/rmq"
	"CloudGameServer/service/public"
	"CloudGameServer/service/user"
	bilicoin "CloudGameServer/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"net/http"
	"sync"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSSessionMap 记录Map
var WSSessionMap sync.Map

type UserSession struct {
	UID           string          `json:"uid"`
	Stub          string          `json:"stub"`
	Aid           string          `json:"aid"`
	State         string          `json:"state"`
	RemoteAddress string          `json:"remote_address"`
	Created       int64           `json:"created"`
	DeviceID      string          `json:"device_id"`
	Ws            *websocket.Conn `json:"-"`
	PackageName   string          `json:"-"`
	Game          interface{}     `json:"game"` // 游戏信息
}

func PushHandler(c *gin.Context) {

	addr, _ := c.RemoteIP()

	// 协议升级
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		bilicoin.Fatal("[WS] upgrade connection failed", logrus.Fields{"addr": addr.String(), "err": err.Error()})
		return
	}

	var isLogin = false
	var info *user.User
	var running = false
	var applied = false

	// 存根生成
	spStub := bilicoin.CreateMD5(time.Now().String())

	// 处理断开
	defer func() {
		ws.Close()
		bilicoin.Info("[WS] close ws connection", logrus.Fields{"addr": addr.String()})

		if running {

			// 关闭推流与运行状态
			deviceID, ok := getDevice(spStub)
			if ok {
				ins, ok1 := WSSessionMap.Load(spStub)
				if ok1 {
					in := ins.(UserSession)
					gi := in.Game.(*public.Update)

					rmq.CloseSender(spStub, deviceID)
					rmq.CloseApp(spStub, deviceID, gi.Pack)
				}
			}
		}

		// 清除本地存根
		WSSessionMap.Delete(spStub)

		// 清除redis
		device.Try2Detach(spStub)

	}()

	// 添加存根到map
	WSSessionMap.Store(spStub, UserSession{
		UID:           "未绑定",
		Aid:           "未绑定",
		DeviceID:      "未分配",
		Stub:          spStub,
		State:         "未登录",
		RemoteAddress: addr.String(),
		Created:       time.Now().Unix(),
		Ws:            ws,
	})

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			bilicoin.Warn("[WS] websocket reset due to some inner problems")
			return
		}

		var res Push
		err = json.Unmarshal(message, &res)
		if err != nil {
			bilicoin.Warn("[WS] websocket reset due to some inner problems")
			return
		}

		switch res.Op {
		case REQ_LOGIN:
			used := user.CheckToken(res.Data)

			if used {
				// 获取用户信息
				info, err = user.GetInfoByToken(res.Data)

				// track
				res.Stub = spStub //...
				res.Op = SU_LOGIN
				res.Data = "login ok"

				//ret, err := json.Marshal(res)
				//if err != nil {
				//	bilicoin.Warn("[WS] websocket reset due to some inner problems")
				//	return
				//}

				//ws.WriteMessage(mt, ret)

				err := ws.WriteJSON(res)
				if err != nil {
					bilicoin.Warn("[WS] websocket reset due to some inner problems")
					return
				}

				if err != nil {
					bilicoin.Warn("[WS] websocket reset due to some inner problems")
					return
				}
				isLogin = true
				bilicoin.Info("[WS] login succeed", logrus.Fields{"userID": info.Uid})

				infojson, err := json.Marshal(info)
				if err != nil {
					bilicoin.Warn("[WS] websocket reset due to some inner problems")
					return
				}

				// online ask
				db.Rdb.Set(context.TODO(), "user:online:"+info.Uid, infojson, time.Second*120)

				// session
				ins, ok := WSSessionMap.Load(spStub)
				if ok {
					p := ins.(UserSession)
					p.UID = info.Uid
					p.State = "已登录"
					WSSessionMap.Store(spStub, p)
				}

			} else {
				err := ws.WriteJSON(Push{
					Stub: "",
					Op:   UNAUTH,
					Data: "oops unauth",
				})
				if err != nil {
					bilicoin.Warn("[WS] websocket reset due to some inner problems")
					return
				}
			}
		case REQ_HB:
			if isLogin {
				db.Rdb.Expire(context.TODO(), "user:online:"+info.Uid, time.Second*120)
			}
		case REQ_EXIT:
			// 退出
			if running {
				deviceID, ok := getDevice(spStub)
				if ok {
					ins, ok1 := WSSessionMap.Load(spStub)
					if ok1 {
						in := ins.(UserSession)
						gi := in.Game.(*public.Update)

						rmq.CloseSender(spStub, deviceID)
						rmq.CloseApp(spStub, deviceID, gi.Pack)
						running = false
					}
				}
			}
			return
		case REQ_APPLY:
			bilicoin.Info("[WS] try to apply a device", logrus.Fields{"stub": spStub})
			// 尝试申请容器 ok 则申请成功
			ok := device.Try2Apply(spStub)
			if ok {
				ws.WriteJSON(Push{
					Stub: spStub,
					Op:   ASK_APPLY_SUCCEED,
					Data: "申请成功",
				})
				applied = true
				bilicoin.Info("[WS] apply succeed", logrus.Fields{"stub": spStub})
			} else {
				// 可以扩展排队系统
				ws.WriteJSON(Push{
					Stub: spStub,
					Op:   ASK_APPLY_FAILED,
					Data: "等待5s后尝试",
				})
				bilicoin.Info("[WS] apply failed", logrus.Fields{"stub": spStub})
			}
		case REQ_START:
			if applied {
				// 使用会话获取已分配设备
				deviceID, ok := getDevice(spStub)

				// 查询游戏信息 res.Data里面放aid的
				game, ok1 := GetGameInfoByAid(res.Data)
				if !ok1 {
					ws.WriteJSON(Push{
						Stub: spStub,
						Op:   ASK_START_FAILED,
						Data: "服务器内部错误",
					})
				} else {
					if ok {
						ins, ok := WSSessionMap.Load(spStub)
						if ok {
							p := ins.(UserSession)
							p.Aid = "app"
							p.DeviceID = deviceID
							p.State = "游戏中"
							p.Game = game // 添加游戏信息
							WSSessionMap.Store(spStub, p)

							rmq.OpenSender(spStub, deviceID)
							rmq.OpenApp(spStub, deviceID, game.Title)
							running = true
						}
					}
				}

			}
		}
	}
}

func getDevice(sessionID string) (string, bool) {
	ret := db.Rdb.HGet(context.TODO(), "sessionMap", sessionID).Val()
	if ret == "" {
		return "", false
	}
	return ret, true
}

func GetGameInfoByAid(aid string) (*public.Update, bool) {
	u := public.Update{}
	if err := db.MDB().FindOne(bson.M{"aid": aid}, &u); err != nil {
		return nil, false
	}
	return &u, true
}

type Push struct {
	Stub string `json:"stub"`
	Op   int    `json:"op"`
	Data string `json:"data"`
}

const (
	UNAUTH            = -1
	REQ_LOGIN         = 0
	SU_LOGIN          = 1
	REQ_HB            = 2
	REQ_EXIT          = 3
	REQ_APPLY         = 4
	REQ_START         = 5
	ASK_APPLY_SUCCEED = 6
	ASK_APPLY_FAILED  = 7
	ASK_START_SUCCEED = 8
	ASK_START_FAILED  = 9
)
