package user

import "github.com/gin-gonic/gin"

type User struct {
	Username string
	Mobile   string
	Avatar   string
	Password string

}

func HandlerReg(context *gin.Context) {
	// post
	username := context.PostForm("username")
	password := context.PostForm("password")


	println(username)
	println(password)

}
