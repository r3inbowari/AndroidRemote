package main

import (
	"RemoteServer/rpc"
	bilicoin "RemoteServer/utils"
	"RemoteServer/ws"
)

const ()

var ak = 0

func main() {
	bilicoin.InitLogger()
	bilicoin.AppInfo("streamer")

	w := ws.NewWsServer()
	go w.Start()
	rpc.StartRPC()
}
