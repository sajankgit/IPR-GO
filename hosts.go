package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

func hosts(subnet *net.IPNet) int {
	cidr := strings.Split(subnet.String(), "/")
	pow, err := strconv.Atoi(cidr[1])
	if err != nil {
		fmt.Println("Error !!", err)
	}
	n := float64(32 - pow)
	res := math.Pow(2, n)
	return int(res)
}
