package manager

import (
	"CloudGameServer/service/manager"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MappingManager(s *gin.RouterGroup) {
	bilicoin.Info("[M] manager service mapping")



	s.GET("/version", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"version": "v0.1.2"})
	})

	s.GET("/sessions", manager.GetSessions)
	s.GET("/kill/:stub", manager.KillSession)
}
