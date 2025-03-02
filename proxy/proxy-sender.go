package proxy

type ProxySender struct{}

func (s *ProxySender) Run(address []int, port int) {
	// Implementation for receiving functionality goes here
}

func (s *ProxySender) New() *ProxyService {
	proxy_sender := &ProxySender{}
	var proxy ProxyService = proxy_sender
	return &proxy
}
