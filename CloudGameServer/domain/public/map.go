package public

import (
	"CloudGameServer/service"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
)

func MappingGameService(s *gin.RouterGroup) {
	bilicoin.Info("[GS] public service mapping")
	// banner 数据
	s.GET("/home/banner", service.GetBanner)
}

