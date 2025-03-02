package http_server

import (
	"fmt"
	"net/http"
)

type HttpServer struct{}

func New() *HttpServer {
	http_server := &HttpServer{}
	return http_server
}

func (hs HttpServer) NewPath(path string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})
}

func (hs HttpServer) Run(ipBind []int, port int) {
	server_bind := fmt.Sprintf("%d.%d.%d.%d", ipBind[0], ipBind[1], ipBind[2], ipBind[3])
	server_port := fmt.Sprintf("%d", port)
	fmt.Printf("INFO: listening on %s:%s\n", server_bind, server_port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", server_bind, server_port), nil)
}
