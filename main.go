package main

import (
	"fmt"
	"os"
)

func main() {
	argument := os.Args[1]
	ip, subnet := validate(argument)
	fmt.Printf("IP - %s\n", ip)
	fmt.Printf("Subnet - %s\n", subnet)
	num_hosts := hosts(subnet)
	fmt.Printf("Number of hosts - %d\n", num_hosts)
	ip_range := ip_range(argument)
	fmt.Println(ip_range)
}
