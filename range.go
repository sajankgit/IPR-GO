package main

import (
	"fmt"
	"net"
)

func ip_range(argument string) string {
	_, ipnet, err := net.ParseCIDR(argument)
	if err != nil {
		fmt.Println("Error parsing IP range:", err)
	}
	firstIP := ipnet.IP
	lastIP := net.IP{0, 0, 0, 0}
	for i := range ipnet.Mask {
		lastIP[i] = ipnet.IP[i] | ^ipnet.Mask[i]
	}
	return "IP range: " + firstIP.String() + " - " + lastIP.String()

}
