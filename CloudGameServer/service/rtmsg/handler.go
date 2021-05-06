package rtmsg

import (
	"CloudGameServer/db"
	"CloudGameServer/service/user"
	bilicoin "CloudGameServer/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func PushHandler(c *gin.Context) {

	addr, _ := c.RemoteIP()

	// 协议升级
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		bilicoin.Fatal("[WS] upgrade connection failed", logrus.Fields{"addr": addr.String(), "err": err.Error()})
		return
	}

	// 处理断开
	defer func() {
		ws.Close()
		bilicoin.Info("[WS] close ws connection", logrus.Fields{"addr": addr.String()})
	}()

	var isLogin = false
	var info *user.User

	for {
		mt, message, err := ws.ReadMessage()
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
				res.Stub = bilicoin.CreateMD5(time.Now().String()) //...
				res.Op = SU_LOGIN
				res.Data = "login ok"

				ret, err := json.Marshal(res)
				if err != nil {
					bilicoin.Warn("[WS] websocket reset due to some inner problems")
					return
				}

				ws.WriteMessage(mt, ret)

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

				db.Rdb.Set(context.TODO(), "user:online:"+info.Uid, infojson, time.Second*120)
			}
		case REQ_HB:
			if isLogin {
				db.Rdb.Expire(context.TODO(), "user:online:"+info.Uid, time.Second*120)
			}
		case REQ_EXIT:
			// 退出
			return
		case REQ_APPLY:
			println("apply")
			// 查询可用机器

		}

	}
}

type Push struct {
	Stub string `json:"stub"`
	Op   int    `json:"op"`
	Data string `json:"data"`
}

const (
	REQ_LOGIN = 0
	SU_LOGIN  = 1
	REQ_HB    = 2
	REQ_EXIT  = 3
	REQ_APPLY = 4
)
