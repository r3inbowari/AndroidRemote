package manager

import (
	"CloudGameServer/service/manager"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/load"
	"net/http"
)

func MappingManager(s *gin.RouterGroup) {
	bilicoin.Info("[M] manager service mapping")

	s.GET("/version", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"version": "v0.1.2"})
	})

	s.GET("/sessions", manager.GetSessions)
	s.GET("/kill/:stub", manager.KillSession)

	s.GET("/sys/heap", manager.SystemHeap)
	s.GET("/sys/load", func(context *gin.Context) {
		info, _ := load.Avg()
		context.JSON(http.StatusOK, info)
	})

	// 要分页的
	s.GET("/device/count", manager.GetDevicesCount)
	s.GET("/device/all", manager.GetDevices)

	s.GET("/play/log", manager.GetPlayLog)


}
