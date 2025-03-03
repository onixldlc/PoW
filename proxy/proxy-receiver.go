package proxy

import (
	"fmt"
	"os"

	"github.com/PoW/v2/crypto"
	"github.com/PoW/v2/proxy/receiver"
	"github.com/PoW/v2/server/websocket_server"
)

type ProxyReceiver struct {
	ws       *websocket_server.WebSocketServer
	callback func(*websocket_server.WebSocketServer, []byte) error
	crypto   crypto.Crypto
}

func (r *ProxyReceiver) Run(address []int, port int) {
	wsHandle := websocket_server.New()
	wsHandle.SetCallback(r.callback)
	wsHandle.Run(address, port)
}

func (r *ProxyReceiver) exportkey() {
	public_cert, err := r.crypto.RSA.GetPublicKey()
	if err != nil {
		panic(fmt.Sprintf("Error getting public key: %s", err))
	}

	os.WriteFile("/cert/authorized_key", []byte(public_cert), 0644)
}

func (r *ProxyReceiver) New() *ProxyService {
	wsHandle := websocket_server.New()
	crypto_service := crypto.New()
	r.exportkey()

	proxy_receiver := &ProxyReceiver{
		ws:       wsHandle,
		callback: receiver.Callback,
		crypto:   crypto_service,
	}
	var proxy ProxyService = proxy_receiver
	return &proxy
}
