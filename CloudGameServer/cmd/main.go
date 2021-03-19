package main

import (
	"CloudGameServer"
	bilicoin "CloudGameServer/utils"
	"github.com/eclipse/paho.mqtt.golang"
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("Main")

	CloudGameServer.MQTTMapping("hello", func(client mqtt.Client, message mqtt.Message) {
		println(string(message.Payload()))
	})

	select {

	}
}
