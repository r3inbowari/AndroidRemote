package CloudGameServer

import (
	"CloudGameServer/domain/manager"
	"CloudGameServer/domain/public"
	"CloudGameServer/domain/rtmsg"
	"CloudGameServer/domain/user"
	"CloudGameServer/service"
	user2 "CloudGameServer/service/user"
	bilicoin "CloudGameServer/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var s *gin.Engine

func BCApplication() {
	bilicoin.Info("[WS] Cloud Game Server running")
	bilicoin.Info("[BCS] Listened on " + bilicoin.GetConfig().APIAddr)

	s = gin.Default()
	s.Use(CorsSimple())

	s.Use(gin.Recovery())

	// 公共内容分发
	gameService := s.Group("/public")
	gameService.Use(CorsSimple())
	public.MappingGameService(gameService)


	// 用户服务
	userService := s.Group("/v1")
	// 校验中间件
	userService.Use(UserAuth())
	user.MappingUser(userService)

	rtmsg.MappingRTMsg(s)

	// 管理服务
	managerService := s.Group("/m")
	manager.MappingManager(managerService)

	//go s.Run(bilicoin.GetConfig().APIAddr)
	s.Run(bilicoin.GetConfig().APIAddr)
}

func CorsComplex() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}

func CorsSimple() gin.HandlerFunc {
	return func(c *gin.Context) {
		// origin filter
		origin := c.Request.Header.Get("origin")
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, DELETE")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// UserAuth 一个简单的验证中间件
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() == "/v1/login" || c.FullPath() == "/v1/reg" {
			bilicoin.Info("[HR] login")
			c.Next()
			return
		} else if c.FullPath() == "/push" {
			bilicoin.Info("[HR] push")
			c.Next()
			return
		} else {
			// token get
			auth := c.GetHeader("Authorization")
			if auth == "" {
				c.JSON(http.StatusUnauthorized, service.FailedResponse(", unauthorized", bilicoin.UserUnauthorized))
				c.Abort()
				return
			}
			token := bilicoin.ParseToken(auth)
			if user2.CheckToken(token) {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, service.FailedResponse(", unauthorized", bilicoin.UserUnauthorized))
				c.Abort()
			}
			return
		}
	}
}
