package rpc

import (
	"RemoteServer/control"
	"RemoteServer/event"
	"RemoteServer/hello"
	pics "RemoteServer/jpeg"
	bilicoin "RemoteServer/utils"
	"google.golang.org/grpc"
	"net"
)

func RegRPCServer(s *grpc.Server) {
	hello.RegisterHelloServer(s, &hello.Server{})
	pics.RegisterJPEGServer(s, &pics.Server{})
	event.RegisterEventServer(s, &event.Server{})
	control.RegisterChatServer(s, &control.Streamer{})
}

func StartRPC() {
	bilicoin.Info("[RPC] gRPC start listened on " + bilicoin.GetConfig().APIAddr)
	lis, err := net.Listen("tcp", bilicoin.GetConfig().APIAddr)
	if err != nil {
		bilicoin.Fatal("[RPC] failed to listen: " + err.Error())
	}

	s := grpc.NewServer()
	RegRPCServer(s)

	if err := s.Serve(lis); err != nil {
		bilicoin.Fatal("[RPC] failed to serve: " + err.Error())
	}
}
