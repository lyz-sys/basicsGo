package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Perter", "Giana", "Adriano", "Aarno", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	var goldNum int
	var name string
	for _, val := range users {
		fmt.Println(val)
		name = val
		goldNum = 0
		val = strings.ToLower(val)
		goldNum += strings.Count(val, "e") * 1
		goldNum += strings.Count(val, "i") * 2
		goldNum += strings.Count(val, "o") * 3
		goldNum += strings.Count(val, "u") * 4
		coins -= goldNum
		distribution[name] = goldNum
	}
	fmt.Println(distribution)
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩:", left)
}
