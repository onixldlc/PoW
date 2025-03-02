package proxy

type ProxyReceiver struct{}

func (r *ProxyReceiver) Run(address []int, port int) {
	// Implementation for receiving functionality goes here
}

func (r *ProxyReceiver) New() *ProxyService {
	proxy_receiver := &ProxyReceiver{}
	var proxy ProxyService = proxy_receiver
	return &proxy
}
