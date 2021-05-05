package rtmsg

import (
	"CloudGameServer/service/rtmsg"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
)

func MappingRTMsg(s *gin.Engine) {
	bilicoin.Info("[WS] rtmsg mapping")
	s.GET("/push", rtmsg.PushHandler)
}

