package main

import (
	"CloudGameServer"
	"CloudGameServer/db"
	"CloudGameServer/rmq"
	bilicoin "CloudGameServer/utils"
	"time"
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

	//CloudGameServer.MQTTMapping("hello", func(client mqtt.Client, message mqtt.Message) {
	//	println(string(message.Payload()))
	//})
	db.InitRDB()

	rmq.InitRMQPub()

	CloudGameServer.BCApplication()

	time.Sleep(time.Hour * 24)
	// token 获取
	// 进入

}
