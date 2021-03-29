package test

import (
	"CloudGameServer/service/user"
	"testing"
)

//func TestFound(t *testing.T) {
//	db.InitMongo()
//	user1 := user.User{
//		Username: "",
//		Mobile:   "15598870762",
//		Avatar:   "",
//		Password: "",
//	}
//	println(user1.FoundUser())
//}

func TestCreateToken(t *testing.T) {
	user1 := user.User{
		Username: "",
		Mobile:   "15598870762",
		Avatar:   "",
		Password: "",
	}

	user1.CreateToken("12")
}


