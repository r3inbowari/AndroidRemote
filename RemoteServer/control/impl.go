package control

import (
	"RemoteServer/device"
	bilicoin "RemoteServer/utils"
	"io"
)

// Streamer 服务端
type Streamer struct {
	UnimplementedChatServer
}

// BidStream 实现了 ChatServer 接口中定义的 BidStream 方法
func (s *Streamer) BidStream(stream Chat_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[CHAT] a close signal send from client-side")
			return ctx.Err()
		default:
			// recv msg from client
			result, err := stream.Recv()
			if err == io.EOF {
				bilicoin.Info("[CHAT] stream EOF")
				return nil
			} else if err != nil {
				bilicoin.Info("[CHAT] stream error")
				return err
			}

			switch result.GetType() {
			case bilicoin.REG:
				device.Renew(result.GetId(), result.GetInput())
			}
			bilicoin.Info("[CHAT] get")
			if err := stream.Send(&ChatResponse{Output: "服务端返回: " + result.Input}); err != nil {
				return err
			}

		}
	}
}
