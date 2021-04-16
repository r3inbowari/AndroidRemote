package event

import (
	bilicoin "RemoteServer/utils"
	"github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	UnimplementedEventServer
}

// Event event struct
type Event struct {
	Id      string `json:"id"`
	Type    int    `json:"type"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Contact int    `json:"contact"`
	Ts      int64  `json:"ts"`
}

// Event 执行
func (t *Event) exec() {

}

type TapSession struct {
	Id          string
	TouchStream Event_EventReqServer
	Ch2         chan *EventResponse // 命令传输通道
	// 持有者 CAS..?
}

var TapSessionsMap sync.Map

func (cs *TapSession) CancelTapSession() {
	TapSessionsMap.Delete(cs.Id)
}

func RegTapSession(id string, stream Event_EventReqServer) *TapSession {
	var cs TapSession
	cs.TouchStream = stream
	cs.Id = id
	cs.Ch2 = make(chan *EventResponse, 100) // cap = 100
	TapSessionsMap.Store(id, cs)
	return &cs
}

func (s Server) EventReq(request *EventRequest, stream Event_EventReqServer) error {
	var ts *TapSession = nil
	defer func() {
		// release session
		ts.CancelTapSession()
		bilicoin.Info("[EVT] session canceled", logrus.Fields{"id": request.GetId()})
	}()

	ctx := stream.Context()

	switch request.GetType() {
	case bilicoin.REG:
		// reg here
		bilicoin.Info("[EVT] a event device reg", logrus.Fields{"id": request.GetId()})
		ts = RegTapSession(request.GetId(), stream) // reg
	default:
		bilicoin.Fatal("[EVT] unsupported request")
	}

	// ts make no sense
	if ts == nil {
		bilicoin.Fatal("[EVT] unsupported request")
		return ctx.Err()
	}

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[EVT] a close signal send from client-side", logrus.Fields{"id": request.GetId()})
			return ctx.Err()
		case event := <-ts.Ch2:
			bilicoin.Info("[EVT] send event", logrus.Fields{"id": request.GetId(), "x": event.X, "y": event.Y, "contact": event.Contact, "type": event.Type, "ts": event.Ts})
			err := stream.Send(event)
			if err != nil {
				println(err.Error())
				return err
			}
		}
	}

	//if request.Message == "attach event" {
	//	println("start event")
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
