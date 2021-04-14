package control

import (
	"RemoteServer/device"
	bilicoin "RemoteServer/utils"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

// Streamer 服务端
type Streamer struct {
	UnimplementedChatServer
}

// BidStream 实现了 ChatServer 接口中定义的 BidStream 方法
// 写信道传递的注册方式
func (s *Streamer) BidStream(stream Chat_BidStreamServer) error {
	ctx := stream.Context()

	id := RegProcess(stream)
	cs := RegSession(id, stream)
	go cs.ChatProcess()

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[CHAT] a close signal send from client-side")
			// if pc here, the chat process will throw error and
			// return defer func and then release session resource.
			return ctx.Err()
		}
	}
}

func (cs *ChatSession) ChatProcess() {
	defer func() {
		// 清除 session
		cs.CancelSession()
		bilicoin.Info("[CHAT] session canceled", logrus.Fields{"id": cs.Id})
	}()
	for {
		result, err := cs.Stream.Recv()
		if err == io.EOF {
			bilicoin.Info("[CHAT] stream EOF")
			return
		} else if err != nil {
			bilicoin.Info("[CHAT] stream error")
			return
		}

		switch result.GetType() {
		case bilicoin.REG:
			device.Renew(result.GetId(), result.GetInput())
			bilicoin.Info("[CHAT] device service renewal", logrus.Fields{"id": result.GetId()})
		default:
			bilicoin.Fatal("[CHAT] error request")
		}
		if err := cs.Stream.Send(&ChatResponse{Output: "return: " + result.Input}); err != nil {
			return
		}
	}
}

var SessionsMap sync.Map

type ChatSession struct {
	Id     string
	Stream Chat_BidStreamServer
	// sync.RWMutex not need
}

func RegSession(id string, stream Chat_BidStreamServer) *ChatSession {
	var cs ChatSession
	cs.Stream = stream
	cs.Id = id
	SessionsMap.Store(id, cs)
	return &cs
}

func CancelSession(id string) {
	SessionsMap.Delete(id)
}

func (cs *ChatSession) CancelSession() {
	SessionsMap.Delete(cs.Id)
}

func RegProcess(stream Chat_BidStreamServer) string {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[CHAT] a close signal send from client-side")
			return ""
		default:
			// recv msg from client
			result, err := stream.Recv()
			if err == io.EOF {
				bilicoin.Info("[CHAT] stream EOF")
				return ""
			} else if err != nil {
				bilicoin.Info("[CHAT] stream error")
				return ""
			}

			switch result.GetType() {
			case bilicoin.REG:
				device.Renew(result.GetId(), result.GetInput())
				bilicoin.Info("[CHAT] reg", logrus.Fields{"id": result.GetId()})
			default:
				bilicoin.Fatal("[CHAT] error request")
			}

			if err := stream.Send(&ChatResponse{Output: "return: " + result.Input}); err != nil {
				return ""
			}
			bilicoin.Info("[Chat] reg done", logrus.Fields{"id": result.Id})

			//go func() {
			//	bilicoin.Warn("[TEST] go start")
			//	cs, err := SessionsMap.Load(result.Id)
			//	if err {
			//		bilicoin.Warn("[TEST] error")
			//	}
			//	a := cs.(ChatSession).Stream
			//	for  {
			//		if err := a.Send(&ChatResponse{Output: "return: " + result.Input}); err != nil {
			//			bilicoin.Warn("[TEST] error")
			//			return
			//		}
			//		bilicoin.Warn("[TEST] send ok")
			//		time.Sleep(time.Second * 1)
			//	}
			//}()
			return result.Id
		}
	}
}
