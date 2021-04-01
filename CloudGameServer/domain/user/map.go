package user

import (
	"CloudGameServer/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MappingUser(s *gin.Engine) {
	s.GET("v1/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"version": "v0.1.2"})
	})

	s.POST("v1/reg", user.HandlerReg)
	s.POST("v1/login", user.HandlerLogin) // bearer token
}
