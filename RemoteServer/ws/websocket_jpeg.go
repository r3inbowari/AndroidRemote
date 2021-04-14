package ws

import (
	bilicoin "RemoteServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var SessionMap sync.Map

func ping(c *gin.Context) {
	addr, _ := c.RemoteIP()
	sessionID := c.Query("session")

	if sessionID == "" {
		bilicoin.Fatal("[WS] unavailable params", logrus.Fields{"addr": addr.String()})
		// c.JSON(http.StatusNotFound, FailedResponse(errors.New("unavailable params").Error(), bilicoin.InternalServerError))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	println(sessionID)

	val, ok := SessionMap.Load(sessionID)

	if !ok {
		bilicoin.Fatal("[WS] unavailable session", logrus.Fields{"addr": addr.String()})
		// c.JSON(http.StatusNotFound, FailedResponse(errors.New("unavailable params").Error(), bilicoin.InternalServerError))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	val

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		bilicoin.Fatal("[WS] upgrade connection failed", logrus.Fields{"addr": addr.String()})
		return
	}
	bilicoin.Info("[WS] upgraded connection", logrus.Fields{"addr": addr.String(), "sessionID": sessionID})

	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func JPEGWebsocketServer() {
	SessionMap.Store("a8f5f167f44f4964e6c998dee827110c", "deab9dbaaa74541d")

	bindAddress := bilicoin.GetConfig().WsAddr
	r := gin.Default()
	r.Use(SessionMiddleware)
	r.GET("/screen", ping)
	_ = r.Run(bindAddress)
	bilicoin.Info("[WS] running -> screen share service")
}

// SessionMiddleware 中间件
func SessionMiddleware(c *gin.Context) {

}
