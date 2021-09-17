package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(t(50))
	fmt.Println(time.Now())
}

func t(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return t(n-1) + t(n-2)
}
