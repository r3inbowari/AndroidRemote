package user

import (
	"CloudGameServer/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MappingUser(s *gin.RouterGroup) {
	s.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"version": "v0.1.2"})
	})

	s.POST("/reg", user.HandlerReg)
	s.POST("/login", user.HandlerLogin) // bearer token

	s.GET("/info", user.HandlerInfo)

}
