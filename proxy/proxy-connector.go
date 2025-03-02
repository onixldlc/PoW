package proxy

type ProxyConnector struct{}

func (c *ProxyConnector) Run(address []int, port int) {
	// Implementation for middleware functionality goes here
}

func (c *ProxyConnector) New() *ProxyService {
	proxy_receiver := &ProxyConnector{}
	var proxy ProxyService = proxy_receiver
	return &proxy
}
