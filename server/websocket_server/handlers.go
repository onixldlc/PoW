package websocket_server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (server *WebSocketServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	server.register <- conn

	go server.handleMessages(conn)
}

func (server *WebSocketServer) handleMessages(conn *websocket.Conn) {
	defer func() {
		server.unregister <- conn
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Printf("Received message: %s", message)

		server.broadcast <- message
	}
}
