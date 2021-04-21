package main

import (
	"RemoteServer/db"
	"RemoteServer/queue"
	"RemoteServer/rpc"
	bilicoin "RemoteServer/utils"
	"RemoteServer/ws"
)

var (
	gitHash   string
	buildTime string
	goVersion string
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("streamer", gitHash, buildTime, goVersion)

	db.InitRDB()
	queue.InitRMQ()
	//w := ws.NewWsServer()
	//go w.Start()

	go ws.JPEGWebsocketServer()
	rpc.StartRPC()
}
