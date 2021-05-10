package public

import (
	"CloudGameServer/service/public"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
)

func MappingGameService(s *gin.RouterGroup) {
	bilicoin.Info("[GS] public service mapping")
	// banner 数据
	s.GET("/home/banner", public.GetBanner)
	s.DELETE("/home/banner/:id", public.DelBanner)
	s.POST("/home/banner", public.AddBanner)

	s.GET("/home/hot", public.GetHot)
	s.DELETE("/home/hot/:id", public.DelHot)
	s.POST("/home/hot", public.AddHot)

	s.GET("/home/update", public.GetUpdate)
	s.DELETE("/home/update/:id", public.DelUpdate)
	s.POST("/home/update", public.AddUpdate)
}
