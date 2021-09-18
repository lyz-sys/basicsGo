package main

import (
	"fmt"
	"time"
)

type intnum uint64

func main() {
	fmt.Println(time.Now())
	var n intnum = 50
	fmt.Println(n.t())
	fmt.Println(time.Now())
}

func (n intnum) t() uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return (n - 1).t() + (n - 2).t()
}
