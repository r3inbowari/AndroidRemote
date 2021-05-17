package manager

import (
	"CloudGameServer/db"
	"CloudGameServer/service/rtmsg"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"net/http"
	"runtime"
)

func GetSessions(c *gin.Context) {

	var ret []rtmsg.UserSession

	rtmsg.WSSessionMap.Range(func(key, value interface{}) bool {
		ret = append(ret, value.(rtmsg.UserSession))
		return true
	})

	c.JSON(200, bilicoin.ResponseWrapWithDate(200, "", ret))
}

func KillSession(c *gin.Context) {
	stub := c.Param("stub")

	ret, ok := rtmsg.WSSessionMap.Load(stub)

	if ok {
		p := ret.(rtmsg.UserSession)
		err := p.Ws.Close()
		if err != nil {
			c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
			return
		}
	}
	c.JSON(200, bilicoin.ResponseOKWrap("ok"))
}

func SystemHeap(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	c.JSON(200, &m)
}

func GetDevicesCount(c *gin.Context) {
	results, err := db.Rdb.Keys(context.TODO(), "reg:info:*").Result()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, len(results))
}

// GetDevices ..temp
func GetDevices(c *gin.Context) {
	//results, err := db.Rdb.Keys(context.TODO(), "reg:info:*").Result()
	//if err != nil {
	//	return
	//}
	results, err := db.Rdb.Keys(context.TODO(), "reg:info:*").Result()
	if err != nil {
		return
	}

	var rets string
	rets = "["

	vl := len(results)
	for k, v := range results {
		// 事务
		result := db.Rdb.Get(context.TODO(), v)
		//rets = append(rets, result.Val())
		rets += result.Val()

		if k+1 == vl {
			rets += "]"
		} else {
			rets += ","
		}
	}
	c.String(200, rets)
}

func GetPlayLog(c *gin.Context) {
	var pl []rtmsg.PlayLog
	err := db.MDB().FindAll(bson.M{}, &pl)
	if err != nil {
		bilicoin.Warn("error find all")
	}
	c.JSON(200, bilicoin.ResponseWrapWithDate(200, "ok", pl))
}
