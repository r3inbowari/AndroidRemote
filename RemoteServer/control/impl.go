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

	id := RegChatProcess(stream)
	cs := RegChatSession(id, stream)
	go cs.ChatProcess()

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[CHAT] a close signal send from client-side", logrus.Fields{"id": cs.Id})
			// if pc here, the chat process will throw error and
			// return defer func and then release session resource.
			return ctx.Err()
		}
	}
}

func (cs *ChatSession) ChatProcess() {
	defer func() {
		// 清除 session
		cs.CancelChatSession()
		bilicoin.Info("[CHAT] session canceled", logrus.Fields{"id": cs.Id})
	}()
	for {
		result, err := cs.ChatStream.Recv()
		if err == io.EOF {
			bilicoin.Info("[CHAT] stream EOF", logrus.Fields{"id": cs.Id})
			return
		} else if err != nil {
			bilicoin.Info("[CHAT] stream error", logrus.Fields{"id": cs.Id})
			return
		}

		switch result.GetType() {
		case bilicoin.REG:
			device.Renew(result.GetId(), result.GetInput())
			bilicoin.Info("[CHAT] device service renewal", logrus.Fields{"id": result.GetId()})
		case bilicoin.REQ_START_SENDER:
			// start screen share
			// @param output the response of the order
			bilicoin.Info("[CHAT] device start screen share ok", logrus.Fields{"id": result.GetId()})
		case bilicoin.REQ_PAUSE_SENDER:
			// pause screen share
			// @param output the response of the order
			bilicoin.Info("[CHAT] device pause screen share ok", logrus.Fields{"id": result.GetId()})
		default:
			bilicoin.Fatal("[CHAT] unsupported request")
		}
		if err := cs.ChatStream.Send(&ChatResponse{Output: "return: " + result.Input}); err != nil {
			return
		}
	}
}

var ChatSessionsMap sync.Map

// ChatSession RegChatSession 暂时不改了
type ChatSession struct {
	Id         string
	ChatStream Chat_BidStreamServer
	// sync.RWMutex not need
}

func RegChatSession(id string, stream Chat_BidStreamServer) *ChatSession {
	var cs ChatSession
	cs.ChatStream = stream
	cs.Id = id
	ChatSessionsMap.Store(id, cs)
	return &cs
}

func CancelSession(id string) {
	ChatSessionsMap.Delete(id)
}

func (cs *ChatSession) CancelChatSession() {
	ChatSessionsMap.Delete(cs.Id)
}

func GetSession(id string) *ChatSession {
	ret, ok := ChatSessionsMap.Load(id)
	if !ok {
		return nil
	}
	cs := ret.(ChatSession)
	return &cs
}

func RegChatProcess(stream Chat_BidStreamServer) string {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[CHAT] a close signal send from client-side", logrus.Fields{"id": "undefined"})
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
				bilicoin.Fatal("[CHAT] unsupported request")
			}

			if err := stream.Send(&ChatResponse{Output: "return: " + result.Input}); err != nil {
				return ""
			}
			bilicoin.Info("[CHAT] reg done", logrus.Fields{"id": result.Id})

			//go func() {
			//	bilicoin.Warn("[TEST] go start")
			//	cs, err := ChatSessionsMap.Load(result.Id)
			//	if err {
			//		bilicoin.Warn("[TEST] error")
			//	}
			//	a := cs.(ChatSession).ChatStream
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
