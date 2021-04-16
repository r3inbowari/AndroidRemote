package touch

import (
	bilicoin "RemoteServer/utils"
	"github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	UnimplementedTouchServer
}

type Touch struct {
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
	cs.Ch2 = make(chan *Touch)
	TapSessionsMap.Store(id, cs)
	return &cs
}

func (s Server) TouchReq(request *TouchRequest, stream Touch_TouchReqServer) error {
	var ts *TapSession
	defer func() {
		// release session
		ts.CancelTapSession()
		bilicoin.Info("[TAP] session canceled", logrus.Fields{"id": request.GetId()})
	}()

	ctx := stream.Context()

	switch request.GetType() {
	case bilicoin.REG:
		// reg here
		bilicoin.Info("[TAP] a touch device reg")
		ts = RegTapSession(request.GetId(), stream) // reg
	default:
		bilicoin.Fatal("[TAP] unsupported request")
	}

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[TAP] a close signal send from client-side", logrus.Fields{"id": request.GetId()})
			return ctx.Err()
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
