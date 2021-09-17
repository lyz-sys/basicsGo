package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Println("sum值", sum)
		fmt.Println("x值", x)
		sum += x
		return sum
	}
}

func main() {
	_, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			//pos(i),
			neg(-2 * i),
		)
		fmt.Println("")
	}
}
