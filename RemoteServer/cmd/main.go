package main

import (
	"RemoteServer/db"
	"RemoteServer/queue"
	"RemoteServer/rpc"
	bilicoin "RemoteServer/utils"
	"RemoteServer/ws"
	"fmt"
)

var (
	gitHash   string
	buildTime string
	goVersion string
)

func main() {
	//args := os.Args
	//if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Commit Hash: %s \n", gitHash)
		fmt.Printf("Build TimeStamp: %s \n", buildTime)
		fmt.Printf("GoLang Version: %s \n", goVersion)
	//}

	bilicoin.InitLogger()
	bilicoin.AppInfo("streamer")

	db.InitRDB()
	queue.InitRMQ()
	//w := ws.NewWsServer()
	//go w.Start()

	go ws.JPEGWebsocketServer()
	rpc.StartRPC()
}
