package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var str string
	for _, v := range ip {
		if str != "" {
			str += "."
		}
		str += fmt.Sprintf("%d", v)
	}
	return str
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
