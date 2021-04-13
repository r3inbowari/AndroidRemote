package main

import (
	"RemoteServer/db"
	"RemoteServer/rpc"
	bilicoin "RemoteServer/utils"
	"RemoteServer/ws"
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("streamer")

	db.InitRDB()

	w := ws.NewWsServer()
	go w.Start()
	rpc.StartRPC()
}
