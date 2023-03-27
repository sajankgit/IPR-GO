package main

import (
	"fmt"
	"net"
	"os"
)

func validate(argument string) (net.IP, *net.IPNet) {
	ip, subnet, err := net.ParseCIDR(argument)
	if err != nil {
		fmt.Println("Error parsing IP range:", err)
		os.Exit(1)
	}
	return ip, subnet
}
