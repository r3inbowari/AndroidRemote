package touch

import (
	bilicoin "RemoteServer/utils"
	"github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	UnimplementedTouchServer
}

// Touch event struct
type Touch struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Contact int    `json:"contact"`
	Ts      int64  `json:"ts"`
}

type TapSession struct {
	Id          string
	TouchStream Touch_TouchReqServer
	Ch2         chan *Touch // 命令传输通道
	// 持有者 CAS..?
}

var TapSessionsMap sync.Map

func (cs *TapSession) CancelTapSession() {
	TapSessionsMap.Delete(cs.Id)
}

func RegTapSession(id string, stream Touch_TouchReqServer) *TapSession {
	var cs TapSession
	cs.TouchStream = stream
	cs.Id = id
	cs.Ch2 = make(chan *Touch, 100) // cap = 100
	TapSessionsMap.Store(id, cs)
	return &cs
}

func (s Server) TouchReq(request *TouchRequest, stream Touch_TouchReqServer) error {
	var ts *TapSession = nil
	defer func() {
		// release session
		ts.CancelTapSession()
		bilicoin.Info("[TAP] session canceled", logrus.Fields{"id": request.GetId()})
	}()

	ctx := stream.Context()

	switch request.GetType() {
	case bilicoin.REG:
		// reg here
		bilicoin.Info("[TAP] a touch device reg", logrus.Fields{"id": request.GetId()})
		ts = RegTapSession(request.GetId(), stream) // reg
	default:
		bilicoin.Fatal("[TAP] unsupported request")
	}

	// ts make no sense
	if ts == nil {
		bilicoin.Fatal("[TAP] unsupported request")
		return ctx.Err()
	}

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[TAP] a close signal send from client-side", logrus.Fields{"id": request.GetId()})
			return ctx.Err()
		case event := <-ts.Ch2:
			bilicoin.Info("[TAP] send event", logrus.Fields{"id": event.Id, "x": event.X, "y": event.Y, "contact": event.Contact, "type": event.Type, "ts": event.Ts})
		}
	}

	//if request.Message == "attach touch" {
	//	println("start touch")
	//	time.Sleep(time.Second * 10)
	//	// ctx := server.Context()
	//	server.Send(&TouchReply{
	//		Type:    0,
	//		X:       300,
	//		Y:       300,
	//		Contact: 0,
	//		Ts:      10,
	//	})
	//
	//	time.Sleep(time.Millisecond * 20)
	//	server.Send(&TouchReply{
	//		Type:    TouchReply_RELEASE,
	//		Contact: 0,
	//		Ts:      10,
	//	})
	//}
}

func (s Server) mustEmbedUnimplementedTouchServer() {
	// panic("implement me")
}
