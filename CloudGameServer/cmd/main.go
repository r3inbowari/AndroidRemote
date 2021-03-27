package main

import (
	"CloudGameServer"
	"CloudGameServer/db"
	"CloudGameServer/service/user"
	bilicoin "CloudGameServer/utils"
	"time"
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("Main")

	db.InitMongo()




	println(user.Exist("xiao"))


	//CloudGameServer.MQTTMapping("hello", func(client mqtt.Client, message mqtt.Message) {
	//	println(string(message.Payload()))
	//})

	CloudGameServer.BCApplication()


	time.Sleep(time.Hour)

}
