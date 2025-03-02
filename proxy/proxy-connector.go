package proxy

import (
	"fmt"
	"net/http"

	"github.com/PoW/v2/proxy/connector"
	"github.com/PoW/v2/server/http_server"
)

type ProxyConnector struct {
	http_server *http_server.HttpServer
	connector   connector.ConnectorPage
}

func (c *ProxyConnector) Run(address []int, port int) {
	fmt.Println("INFO: setting up proxy connector")
	c.http_server = http_server.New()
	c.http_server.NewPath("/", func(w http.ResponseWriter, r *http.Request) {
		page, err := c.connector.GetPage("index")
		if err != nil {
			w.Write([]byte("Error"))
			return
		}
		w.Write([]byte(page))
	})

	fmt.Println("INFO: starting server")
	c.http_server.Run(address, port)
}

func (c *ProxyConnector) New() *ProxyService {
	proxy_receiver := &ProxyConnector{
		http_server: http_server.New(),
		connector:   connector.New(),
	}
	var proxy ProxyService = proxy_receiver
	return &proxy
}
