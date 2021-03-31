package test

import (
	"CloudGameServer/db"
	"CloudGameServer/service/user"
	bilicoin "CloudGameServer/utils"
	"testing"
)

func TestCreateToken(t *testing.T) {
	user1 := user.User{
		Username: "",
		Mobile:   "15598870762",
		Avatar:   "",
		Password: "",
		Uid:      bilicoin.CreateMD5("12"),
	}

	println(user1.CreateToken())
}

func TestCheckToken(t *testing.T) {
	println(user.CheckToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTcxODk1NDgsImp0aSI6ImMyMGFkNGQ3NmZlOTc3NTlhYTI3YTBjOTliZmY2NzEwIiwiaXNzIjoicjNpbmIiLCJuYmYiOjE2MTcxMDMxNDh9.LPb20xTbGBtIu7CXi9g8m_eq8LY4-ZFVhIaJKZSN1OE"))
}

func TestExist(t *testing.T) {
	db.InitMongo()
	user1 := user.User{
		Username: "",
		Mobile:   "r3inbowari",
		Avatar:   "",
		Password: "",
	}
	println(user1.IsExist())
}

	func TestLogin(t *testing.T) {
	db.InitMongo()

	var user = user.User{
		Mobile:   "15598870762",
		Password: "15598870762",
	}

	user.Login()
}
