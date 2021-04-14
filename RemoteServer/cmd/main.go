package main

import (
	"RemoteServer/db"
	"RemoteServer/queue"
	"RemoteServer/rpc"
	bilicoin "RemoteServer/utils"
	"RemoteServer/ws"
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("streamer")

	db.InitRDB()
	queue.InitRMQ()
	w := ws.NewWsServer()
	go w.Start()
	rpc.StartRPC()
}
