package hello

import (
	"golang.org/x/net/context"
	"log"
)

//SayHello 实现了 hello.GreeterServer 中的方法, 对内部的 SyaHello ‘not implemented’ 进行屏蔽
//远程调用过程函数 SayHello
func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

// server is used to implement helloworld.GreeterServer.
type Server struct {
	UnimplementedHelloServer
}
