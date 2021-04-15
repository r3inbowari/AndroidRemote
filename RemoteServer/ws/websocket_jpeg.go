package ws

import (
	pics "RemoteServer/jpeg"
	bilicoin "RemoteServer/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"unsafe"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var SessionMap sync.Map

func screen(c *gin.Context) {
	addr, _ := c.RemoteIP()
	sessionID := c.Query("session")

	// session 检查是否存在
	if sessionID == "" {
		bilicoin.Fatal("[WS] unavailable params", logrus.Fields{"addr": addr.String()})
		// c.JSON(http.StatusNotFound, FailedResponse(errors.New("unavailable params").Error(), bilicoin.InternalServerError))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// session 检查是否有效
	// val 为设备 did
	// 数据源 请求队列
	deviceID, ok := SessionMap.Load(sessionID)
	if !ok {
		bilicoin.Fatal("[WS] unavailable session", logrus.Fields{"addr": addr.String()})
		// c.JSON(http.StatusNotFound, FailedResponse(errors.New("unavailable params").Error(), bilicoin.InternalServerError))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 协议升级
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		bilicoin.Fatal("[WS] upgrade connection failed", logrus.Fields{"addr": addr.String()})
		return
	}
	bilicoin.Info("[WS] upgraded connection", logrus.Fields{"addr": addr.String(), "sessionID": sessionID})

	// 处理断开
	defer func() {
		ws.Close()
		bilicoin.Info("[WS] close connection", logrus.Fields{"addr": addr.String(), "sessionID": sessionID, "deviceID": deviceID})
	}()

	// 获取设备屏幕实例
	// @param did 设备id
	entry, ok := pics.ScreenSessionsMap.Load(deviceID)
	if !ok {
		bilicoin.Info("[WS] illegal or a non-existent device accessed", logrus.Fields{"addr": addr.String(), "sessionID": sessionID, "deviceID": deviceID})
		return
	}
	bilicoin.Info("[WS] try to bind ws with device", logrus.Fields{"addr": addr.String(), "sessionID": sessionID, "deviceID": deviceID, "wsPtr": fmt.Sprintf("%p", unsafe.Pointer(ws))})
	wsChannel := entry.(pics.ScreenSession).Ch2
	wsChannel <- ws
	for {

	}

	//scr := entry.(pics.ScreenSession)
	//for {
	//	input, err := scr.ScreenStream.Recv()
	//	if err == io.EOF {
	//		log.Println("客户端发送的数据流结束")
	//		return
	//	}
	//	if err != nil {
	//		log.Println("接收数据出错:", err)
	//		return
	//	}
	//	err = ws.WriteMessage(websocket.BinaryMessage, input.Data)
	//	if err != nil {
	//		log.Println("接收数据出错:", err)
	//		return
	//	}
	//}
	//for {
	//	mt, message, err := ws.ReadMessage()
	//	if err != nil {
	//		break
	//	}
	//	if string(message) == "ping" {
	//		message = []byte("pong")
	//	}
	//	err = ws.WriteMessage(mt, message)
	//	if err != nil {
	//		break
	//	}
	//}
}

func JPEGWebsocketServer() {
	SessionMap.Store("a8f5f167f44f4964e6c998dee827110c", "deab9dbaaa74541d")

	bindAddress := bilicoin.GetConfig().WsAddr
	r := gin.Default()
	r.Use(SessionMiddleware)
	r.GET("/screen", screen)
	_ = r.Run(bindAddress)
	bilicoin.Info("[WS] running -> screen share service")
}

// SessionMiddleware 中间件
func SessionMiddleware(c *gin.Context) {

}
