package pics

import (
	bilicoin "RemoteServer/utils"
	"RemoteServer/ws"
	"io"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	UnimplementedJPEGServer
}

var ak = 0

func (s *Server) SendJPEG(stream JPEG_SendJPEGServer) error {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[CHAT] a close signal send from client-side")
			return ctx.Err()
		default:
			// 接收从客户端发来的消息
			input, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				return err
			}
			if len(input.Data) > 2 {
				print(input.Data[0])
				println(input.Data[1])
			}
			ws.Exp.Mux.Lock()
			if ws.Exp.Conn != nil {
				ws.Exp.Conn.WriteMessage(0x02, input.Data)
			}
			ws.Exp.Mux.Unlock()
			// log.Printf("get -> " + strconv.Itoa(ak))

		}
	}
	//ak++
	//exp.Mux.Lock()
	//if exp.Conn != nil {
	//	exp.Conn.WriteMessage(0x02, in.Data)
	//}
	//exp.Mux.Unlock()
	//
	// return &JPEGReply{Length: 12}, nil
}
