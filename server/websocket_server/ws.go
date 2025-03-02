package websocket_server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PoW/v2/server/http_server"
	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mutex      sync.Mutex
	upgrader   websocket.Upgrader
	callback   func(*WebSocketServer, []byte) error
}

func New() *WebSocketServer {
	return &WebSocketServer{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (server *WebSocketServer) SetCallback(callback func(*WebSocketServer, []byte) error) {
	server.callback = callback
}

func (server *WebSocketServer) StartEventHandler() {
	for {
		select {
		case client := <-server.register:
			server.registerClient(client)
		case client := <-server.unregister:
			server.unregisterClient(client)
		case message := <-server.broadcast:
			err := server.callback(server, message)
			if err != nil {
				fmt.Println("Error callback message:", err)
			}
		}
	}
}

func (server *WebSocketServer) Run(addr []int, port int) {
	go server.StartEventHandler()

	http_server := http_server.New()
	http_server.NewPath("/ws", server.HandleWebSocket)
	http_server.Run(addr, port)

	fmt.Printf("INFO: WebSocket server listening at %v:%v\n", addr, port)
}
