package main

import (
	"CloudGameServer"
	"CloudGameServer/db"
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

	CloudGameServer.BCApplication()

	time.Sleep(time.Hour * 24)

}
