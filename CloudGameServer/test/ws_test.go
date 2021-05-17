package test

import (
	"CloudGameServer/db"
	"CloudGameServer/service/rtmsg"
	"testing"
)

func TestGetGame(t *testing.T) {
	db.InitMongo()

	info, ok := rtmsg.GetGameInfoByAid("5a1db08d50cff06b19e873c86919e1c")
	if !ok {
		println("12")
		return
	}

	println(info.Pack)
}
