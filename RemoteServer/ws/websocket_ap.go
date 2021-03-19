package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"sync"
	"time"
)

type WsServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}

type WsConn struct {
	*websocket.Conn
	Mux sync.RWMutex
}

var Exp WsConn

func NewWsServer() *WsServer {
	ws := new(WsServer)
	ws.addr = "0.0.0.0:8080"
	ws.upgrade = &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				fmt.Println("path error")
				return false
			}
			return true
		},
	}
	return ws
}

func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ws" {
		httpCode := http.StatusInternalServerError
		reasePhrase := http.StatusText(httpCode)
		fmt.Println("path error ", reasePhrase)
		http.Error(w, reasePhrase, httpCode)
		return
	}

	var err error

	Exp.Conn, err = ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}

	fmt.Println("client connect :", Exp.Conn.RemoteAddr())
	go ws.connHandle(Exp.Conn)
}

func (ws *WsServer) connHandle(conn *websocket.Conn) {
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
					fmt.Printf("ReadMessage timeout remote: %v\n", conn.RemoteAddr())
					return
				}
			}
			// 其他错误，如果是 1001 和 1000 就不打印日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("ReadMessage other remote:%v error: %v \n", conn.RemoteAddr(), err)
			}
			return
		}
		fmt.Println("收到消息：", string(msg))
	}
}

//测试一次性发送 10万条数据给 client, 如果不使用 time.Sleep browser 过了超时时间会断开
func (ws *WsServer) send10(conn *websocket.Conn) {
	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("hello websocket test from server %v", time.Now().UnixNano())
		err := conn.WriteMessage(1, []byte(data))
		if err != nil {
			fmt.Println("send msg faild ", err)
			return
		}
		// time.Sleep(time.Millisecond * 1)
	}
}

func (ws *WsServer) send(conn *websocket.Conn, stopCh chan int) {
	ws.send10(conn)
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

func (ws *WsServer) Start() (err error) {
	ws.listener, err = net.Listen("tcp", ws.addr)
	if err != nil {
		fmt.Println("net listen error:", err)
		return
	}
	err = http.Serve(ws.listener, ws)
	if err != nil {
		fmt.Println("http serve error:", err)
		return
	}
	return nil
}
