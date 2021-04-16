package pics

import (
	bilicoin "RemoteServer/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
	"unsafe"
)

type Server struct {
	UnimplementedJPEGServer
}

func (s *Server) SendJPEG(stream JPEG_SendJPEGServer) error {
	ctx := stream.Context()

	id := RegScreenProcess(stream)
	ss := RegScreenSession(id, stream)
	go ss.ScreenProcess()

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[JPEG] a close signal send from client-side", logrus.Fields{"id": ss.Id})
			return ctx.Err()
		case o := <-ss.Ch2:
			ss.sc.Lock()
			if ss.sc.Conn != nil {
				bilicoin.Info("[JPEG] device session switch to a new ws", logrus.Fields{"oldWsSessionID": ss.sc.WsSessionID, "newWsSession": o.WsSessionID})
			}
			ss.sc.Conn = o.Conn
			ss.sc.WsSessionID = o.WsSessionID
			ss.sc.Unlock()
			bilicoin.Info("[JPEG] ws was successfully attached to the device", logrus.Fields{"addr": ss.sc.Conn.RemoteAddr().String(), "sessionID": ss.sc.WsSessionID, "deviceID": id, "wsPtr": fmt.Sprintf("%p", unsafe.Pointer(ss.sc.Conn))})
		}
	}
	//exp.Mux.Lock()
	//if exp.Conn != nil {
	//	exp.Conn.WriteMessage(0x02, in.Data)
	//}
	//exp.Mux.Unlock()
	//
	// return &JPEGReply{Length: 12}, nil
}

// ScreenProcess 转移
func (cs *ScreenSession) ScreenProcess() {
	defer func() {
		// 清除 session
		cs.CancelScreenSession()
		bilicoin.Info("[JPEG] session canceled", logrus.Fields{"id": cs.Id})
	}()
	// recv can not called(blocking) more than two on the same time
	// it will cause the thread race...
	for {
		result, err := cs.ScreenStream.Recv()
		if cs.sc.Conn == nil {
			// println("skip")
			continue
		}
		// println("1")
		if err == io.EOF {
			bilicoin.Info("[JPEG] stream EOF", logrus.Fields{"id": cs.Id})
			return
		} else if err != nil {
			bilicoin.Info("[JPEG] stream error", logrus.Fields{"id": cs.Id})
			return
		}

		switch result.GetType() {
		case bilicoin.REQ_SEND_JPEG:
			// println("!23")
			cs.sc.Lock()
			if err = cs.sc.Conn.WriteMessage(websocket.BinaryMessage, result.Data); err != nil {
				bilicoin.Warn("[JPEG] exception: the ws has been detached by device implicitly", logrus.Fields{"addr": cs.sc.Conn.RemoteAddr().String(), "device": cs.Id})
				cs.sc.Conn = nil
			}
			cs.sc.Unlock()
			// start screen share
			// @param output the response of the order
			// bilicoin.Info("[JPEG] device start screen share ok", logrus.Fields{"id": result.GetId()})
		default:
			bilicoin.Fatal("[JPEG] unsupported request")
		}
	}
}

var ScreenSessionsMap sync.Map

type ScreenConn struct {
	Conn        *websocket.Conn
	WsSessionID string
	sync.Mutex
}

type ScreenSession struct {
	Id           string
	ScreenStream JPEG_SendJPEGServer
	Ch2          chan *ScreenConn
	sc           *ScreenConn // 线程被持有者 这里指的是 websocket 的 conn
}

func (cs *ScreenSession) CancelScreenSession() {
	ScreenSessionsMap.Delete(cs.Id)
}

func RegScreenSession(id string, stream JPEG_SendJPEGServer) *ScreenSession {
	var cs ScreenSession
	cs.ScreenStream = stream
	cs.Id = id
	cs.Ch2 = make(chan *ScreenConn)
	cs.sc = &ScreenConn{}
	ScreenSessionsMap.Store(id, cs)
	return &cs
}

func RegScreenProcess(stream JPEG_SendJPEGServer) string {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			bilicoin.Info("[JPEG] a close signal send from client-side", logrus.Fields{"id": "undefined"})
			return ""
		default:
			// recv msg from client
			result, err := stream.Recv()
			if err == io.EOF {
				bilicoin.Info("[JPEG] stream EOF")
				return ""
			} else if err != nil {
				bilicoin.Info("[JPEG] stream error")
				return ""
			}

			switch result.GetType() {
			case bilicoin.REG:
				bilicoin.Info("[JPEG] reg", logrus.Fields{"id": result.GetId()})
			default:
				bilicoin.Fatal("[JPEG] unsupported request")
			}
			bilicoin.Info("[JPEG] reg done", logrus.Fields{"id": result.Id})
			return result.Id
		}
	}
}
