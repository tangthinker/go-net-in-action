package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
			if ipnet.IP.To16() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
