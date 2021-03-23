package main

import (
	"CloudGameServer"
	bilicoin "CloudGameServer/utils"
	"time"
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("Main")

	//CloudGameServer.MQTTMapping("hello", func(client mqtt.Client, message mqtt.Message) {
	//	println(string(message.Payload()))
	//})

	CloudGameServer.BCApplication()

	time.Sleep(time.Hour)
}
