package proxy

import (
	"errors"
	"fmt"
	"os"

	"github.com/PoW/v2/crypto"
	"github.com/PoW/v2/proxy/sender"
	"github.com/PoW/v2/server/websocket_server"
)

type ProxySender struct {
	ws             *websocket_server.WebSocketServer
	callback       func(*websocket_server.WebSocketServer, []byte) error
	crypto         crypto.Crypto
	authorized_key string
}

func (s *ProxySender) Run(address []int, port int) {
	wsHandle := websocket_server.New()
	wsHandle.SetCallback(s.callback)
	wsHandle.Run(address, port)
}

func (s *ProxySender) loadAuthorizedKey() {
	if _, err := os.Stat("/cert/authorized_key"); errors.Is(err, os.ErrNotExist) {
		os.Mkdir("/cert", 0755)
		os.WriteFile("/cert/authorized_key", []byte(""), 0644)
	}

	content, err := os.ReadFile("/cert/authorized_key")
	if err != nil {
		fmt.Println("authorized_key cannot be empty")
		panic(fmt.Sprintf("Error reading authorized key: %s", err))
	}

	s.authorized_key = string(content)
}

func (s *ProxySender) New() *ProxyService {
	wsHandle := websocket_server.New()
	crypto_service := crypto.New()
	s.loadAuthorizedKey()

	proxy_sender := &ProxySender{
		ws:       wsHandle,
		callback: sender.Callback,
		crypto:   crypto_service,
	}
	var proxy ProxyService = proxy_sender
	return &proxy
}
