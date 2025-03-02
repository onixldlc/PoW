package cmd

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/PoW/v2/proxy"
)

type CLI struct {
	Sender     bool
	Receiver   bool
	Middleware bool
	BindAddr   string
	BindPort   int
}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) ParseFlags() error {
	flag.Usage = c.customUsage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return errors.New("no arguments provided")
	}

	// Identify the subcommand first
	switch args[0] {
	case "s", "sender":
		c.Sender = true
	case "r", "receiver":
		c.Receiver = true
	case "m", "middleware":
		c.Middleware = true
	default:
		flag.Usage()
		return fmt.Errorf("unknown command: %s", args[0])
	}

	c.BindAddr = "0.0.0.0"

	var subArgs []string
	if len(args) > 1 && !strings.HasPrefix(args[1], "-") {
		c.BindAddr = args[1]
		subArgs = args[2:]
	} else {
		subArgs = args[1:]
	}

	subFlags := flag.NewFlagSet(args[0], flag.ExitOnError)
	subFlags.IntVar(&c.BindPort, "p", 8080, "Bind port (optional)")
	subFlags.Parse(subArgs)
	return nil
}

func (c *CLI) ConfigureProxy() (*proxy.ProxyService, error) {
	proxy_factory := proxy.ProxyFactory{}

	if c.Sender {
		proxy_service, err := proxy_factory.New("sender")
		if err != nil {
			return nil, fmt.Errorf("error creating sender service: %v", err)
		}
		return proxy_service, nil
	} else if c.Receiver {
		proxy_service, err := proxy_factory.New("receiver")
		if err != nil {
			return nil, fmt.Errorf("error creating receiver service: %v", err)
		}
		return proxy_service, nil
	} else if c.Middleware {
		proxy_service, err := proxy_factory.New("connector")
		if err != nil {
			return nil, fmt.Errorf("error creating connector service: %v", err)
		}
		return proxy_service, nil
	}
	return nil, errors.New("please specify one of the following: --sender (-s), --receiver (-r), --middleware (-m)")
}

func (c *CLI) Run() {

	err := c.ParseFlags()
	if err != nil {
		println(err.Error())
		return
	}

	proxy_service, err := c.ConfigureProxy()
	if err != nil {
		panic(fmt.Errorf("error configuring proxy: %v", err))
	}

	ipAddr, err := toSimpleIp(c.BindAddr)
	if err != nil {
		panic(fmt.Errorf("error parsing ip address: %v", err))
	}

	if !isPort(c.BindPort) {
		panic(fmt.Errorf("invalid port number: %d", c.BindPort))
	}

	fmt.Printf("Parsed flags: %+v\n", c)

	service := *proxy_service
	service.Run(ipAddr, c.BindPort)
}
