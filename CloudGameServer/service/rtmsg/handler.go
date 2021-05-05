package rtmsg

import (
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
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

	for {
		mt, _, err := ws.ReadMessage()
		if err != nil {
			bilicoin.Warn("[WS] websocket reset due to some inner problems")
			return
		}

		// println(string(message))
		err = ws.WriteMessage(mt, []byte("pong"))
		if err != nil {
			return
		}
	}

}
