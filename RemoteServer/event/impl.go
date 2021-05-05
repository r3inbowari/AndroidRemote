package event

import (
	bilicoin "RemoteServer/utils"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	UnimplementedEventServer
}

type TapSession struct {
	Id          string
	TouchStream Event_EventReqServer
	Ch2         chan *EventResponse // 命令传输通道
	// 持有者 CAS..?
}

var TapSessionsMap sync.Map

func (s Server) mustEmbedUnimplementedTouchServer() {
	// panic("implement me")
}

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
			bilicoin.Info("[EVT] send event", logrus.Fields{"id": request.GetId(), "x": event.X, "y": event.Y, "contact": event.Contact, "type": int32(event.Type), "ts": event.Ts})
			err := stream.Send(event)
			if err != nil {
				println(err.Error())
				return err
			}
		}
	}
}

// ParseEvent
// Type:    int(bs[3] & 0x0F),
// X:       int(binary.BigEndian.Uint16(bs[7:9])),
// Y:       int(binary.BigEndian.Uint16(bs[9:11])),
// Contact: int((bs[6] & 0xF0) >> 4),
// Ts:      int64(binary.BigEndian.Uint64(bs[11:19])),
func ParseEvent(bs []byte) *EventResponse {
	if len(bs) != 12 {
		return nil
	}
	t := EventResponse{
		Type:    EventResponse_EventType(bs[3] & 0x0F),
		Contact: int32((bs[3] & 0xF0) >> 4),
		X:       int32(binary.BigEndian.Uint16(bs[4:6])),
		Y:       int32(binary.BigEndian.Uint16(bs[6:8])),
		Ts:      int64(binary.BigEndian.Uint32(bs[8:])),
	}
	return &t
}
