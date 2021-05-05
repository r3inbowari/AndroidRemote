package main

import (
	"CloudGameServer"
	"CloudGameServer/db"
	"CloudGameServer/rmq"
	bilicoin "CloudGameServer/utils"
	"time"
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("Main")

	db.InitMongo()

	//CloudGameServer.MQTTMapping("hello", func(client mqtt.Client, message mqtt.Message) {
	//	println(string(message.Payload()))
	//})

	rmq.InitRMQPub()

	CloudGameServer.BCApplication()

	time.Sleep(time.Hour * 24)
	// token 获取
	// 进入

}
