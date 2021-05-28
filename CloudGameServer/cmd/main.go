package main

import (
	"CloudGameServer"
	"CloudGameServer/db"
	"CloudGameServer/device"
	"CloudGameServer/rmq"
	bilicoin "CloudGameServer/utils"
)

var (
	gitHash   string
	buildTime string
	goVersion string
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("Main", gitHash, buildTime, goVersion)

	db.InitMongo()

	db.InitRDB()

	rmq.InitRMQPub()

	device.InitDeviceDispatcher()

	CloudGameServer.BCApplication()

	// time.Sleep(time.Hour * 24)
	// token 获取
	// 进入

}
