package proxy

import (
	"fmt"
)

type ProxyService interface {
	Run(address []int, port int)
	New() *ProxyService
}

type ProxyFactory struct{}

func (p ProxyFactory) New(proxy_type string) (*ProxyService, error) {
	switch proxy_type {
	case "sender":
		proxy_sender := &ProxySender{}
		proxy := proxy_sender.New()
		return proxy, nil
	case "receiver":
		proxy_sender := &ProxyReceiver{}
		proxy := proxy_sender.New()
		return proxy, nil
	case "connector":
		proxy_sender := &ProxyConnector{}
		proxy := proxy_sender.New()
		return proxy, nil
	default:
		return nil, fmt.Errorf("invalid proxy type")
	}
}
