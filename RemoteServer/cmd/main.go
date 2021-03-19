package main

import (
	"RemoteServer/rpc"
	bilicoin "RemoteServer/utils"
	"RemoteServer/ws"
)

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("streamer")

	w := ws.NewWsServer()
	go w.Start()
	rpc.StartRPC()
}
