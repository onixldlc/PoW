package websocket_server

import (
	"log"

	"github.com/gorilla/websocket"
)

func (server *WebSocketServer) BroadcastMessage(message []byte) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for client := range server.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error broadcasting message:", err)
			client.Close()
			delete(server.clients, client)
		}
	}
}

func (server *WebSocketServer) BroadcastToClients(message []byte) {
	server.broadcast <- message
}
