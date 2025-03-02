package cmd

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func isIP(ip string) bool {
	ip = strings.TrimSpace(ip)
	return net.ParseIP(ip) != nil
}

func toSimpleIp(ip string) ([]int, error) {
	if !isIP(ip) {
		return nil, fmt.Errorf("invalid ip address `%s`", ip)
	}

	ipAddr := []int{0, 0, 0, 0}
	for i, v := range strings.Split(ip, ".") {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to int", v)
		}
		ipAddr[i] = num
	}
	return ipAddr, nil
}

func isPort(port int) bool {
	return port > 0 && port < 65536
}
