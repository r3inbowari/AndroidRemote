package rpc

import (
	"RemoteServer/control"
	"RemoteServer/hello"
	pics "RemoteServer/jpeg"
	"RemoteServer/touch"
	bilicoin "RemoteServer/utils"
	"google.golang.org/grpc"
	"net"
)

func RegRPCServer(s *grpc.Server) {
	hello.RegisterHelloServer(s, &hello.Server{})
	pics.RegisterJPEGServer(s, &pics.Server{})
	touch.RegisterTouchServer(s, &touch.Server{})
	control.RegisterChatServer(s, &control.Streamer{})
}

func StartRPC() {
	bilicoin.Info("service -> gRPC start listening on " + bilicoin.GetConfig().APIAddr)
	lis, err := net.Listen("tcp", bilicoin.GetConfig().APIAddr)
	if err != nil {
		bilicoin.Fatal("failed to listen: " + err.Error())
	}

	s := grpc.NewServer()
	RegRPCServer(s)

	if err := s.Serve(lis); err != nil {
		bilicoin.Fatal("failed to serve: " + err.Error())
	}
}
