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
			ss.sc.conn = o
			ss.sc.Unlock()
			bilicoin.Info("[JPEG] ws was successfully bound to the device", logrus.Fields{"deviceID": id, "wsPtr": fmt.Sprintf("%p", unsafe.Pointer(o))})
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
		if cs.sc.conn == nil {
			println("skip")
			continue
		}
		if err == io.EOF {
			bilicoin.Info("[JPEG] stream EOF", logrus.Fields{"id": cs.Id})
			return
		} else if err != nil {
			bilicoin.Info("[JPEG] stream error", logrus.Fields{"id": cs.Id})
			return
		}

		switch result.GetType() {
		case bilicoin.REQ_SEND_JPEG:
			cs.sc.Lock()
			if err = cs.sc.conn.WriteMessage(websocket.BinaryMessage, result.Data); err != nil {
				println("[JPEG] ws disconnected", logrus.Fields{"addr": cs.sc.conn.RemoteAddr().String()})
				cs.sc.conn = nil
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
	conn *websocket.Conn
	sync.Mutex
}

type ScreenSession struct {
	Id           string
	ScreenStream JPEG_SendJPEGServer
	Ch2          chan *websocket.Conn
	sc           ScreenConn // 线程被持有者
}

func (cs *ScreenSession) CancelScreenSession() {
	ScreenSessionsMap.Delete(cs.Id)
}

func RegScreenSession(id string, stream JPEG_SendJPEGServer) *ScreenSession {
	var cs ScreenSession
	cs.ScreenStream = stream
	cs.Id = id
	cs.Ch2 = make(chan *websocket.Conn)
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
