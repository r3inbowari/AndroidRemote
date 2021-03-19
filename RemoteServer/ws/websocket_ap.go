package ws

import (
	bilicoin "RemoteServer/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sync"
	"time"
)

type WebsocketServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}

type WebsocketConn struct {
	*websocket.Conn
	Mux sync.RWMutex
}

var Exp WebsocketConn

func NewWsServer() *WebsocketServer {
	ws := new(WebsocketServer)
	ws.addr = bilicoin.GetConfig().WsAddr
	ws.upgrade = &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				bilicoin.Warn("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				bilicoin.Warn("path error")
				return false
			}
			return true
		},
	}
	bilicoin.Info("service -> a websocket server instance listened on " + ws.addr)
	return ws
}

func (ws *WebsocketServer) Start() (err error) {
	ws.listener, err = net.Listen("tcp", ws.addr)
	if err != nil {
		bilicoin.Fatal("net listen error: " + err.Error())
		return
	}
	err = http.Serve(ws.listener, ws)
	if err != nil {
		bilicoin.Fatal("http serve error: " + err.Error())
		return
	}
	return nil
}

// overwrite
func (ws *WebsocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ws" {
		httpCode := http.StatusInternalServerError
		codeString := http.StatusText(httpCode)
		bilicoin.Warn("path error " + codeString)
		http.Error(w, codeString, httpCode)
		return
	}

	var err error

	Exp.Conn, err = ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		bilicoin.Warn("websocket updated error:" + err.Error())
		return
	}

	bilicoin.Info("client connect: " + Exp.Conn.RemoteAddr().String())

	// 由前级服务器提供一个token表面要控制哪个设备

	go ws.connHandle(Exp.Conn)
	// 进入时应该时等待的状态

	// 等待后端串流
}

func (ws *WebsocketServer) connHandle(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	// stopCh := make(chan int)
	// go ws.send(conn, stopCh)
	for {
		// conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(5000)))
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// close(stopCh)
			// 判断是不是超时
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					bilicoin.Fatal("ReadMessage timeout", logrus.Fields{"remote": conn.RemoteAddr(), "err": err.Error()})
					return
				}
			}
			// 其他错误，如果是 1001 和 1000 就不打印日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				bilicoin.Fatal("ReadMessage error", logrus.Fields{"remote": conn.RemoteAddr(), "err": err.Error()})
			}
			return
		}
		fmt.Println("收到消息：", string(msg))
	}
}
//
////测试一次性发送 10万条数据给 client, 如果不使用 time.Sleep browser 过了超时时间会断开
//func (ws *WebsocketServer) send10(conn *websocket.Conn) {
//	for i := 0; i < 10; i++ {
//		data := fmt.Sprintf("hello websocket test from server %v", time.Now().UnixNano())
//		err := conn.WriteMessage(1, []byte(data))
//		if err != nil {
//			fmt.Println("send msg faild ", err)
//			return
//		}
//		// time.Sleep(time.Millisecond * 1)
//	}
//}

func (ws *WebsocketServer) send(conn *websocket.Conn, stopCh chan int) {
	// ws.send10(conn)
	for {
		select {
		case <-stopCh:
			fmt.Println("connect closed")
			return
		case <-time.After(time.Second * 1):
			data := fmt.Sprintf("hello websocket test from server %v", time.Now().UnixNano())
			err := conn.WriteMessage(1, []byte(data))
			fmt.Println("sending....")
			if err != nil {
				fmt.Println("send msg faild ", err)
				return
			}
		}
	}
}