package user

import (
	"github.com/gin-gonic/gin"
)

func HandlerReg(context *gin.Context) {
	// post
	username := context.PostForm("username")
	password := context.PostForm("password")

	println(username)
	println(password)

}
