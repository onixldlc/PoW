package websocket_server

import (
	"log"

	"github.com/gorilla/websocket"
)

func (server *WebSocketServer) registerClient(client *websocket.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.clients[client] = true
	log.Println("New client connected")
}

func (server *WebSocketServer) unregisterClient(client *websocket.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)
		client.Close()
		log.Println("Client disconnected")
	}
}
