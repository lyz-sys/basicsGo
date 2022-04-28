package main

import (
	"fmt"
)

// CountHowManyWays
// There are n steps
// you can take one or two steps at a time
// how many ways to finish them
func CountHowManyWays(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return CountHowManyWays(n-1) + CountHowManyWays(n-2)
}

func main() {
	var choice uint64
	fmt.Print("Please enter the number of steps:")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid option")
		return
	}
	fmt.Printf("You chose:%d\n", choice)
	fmt.Println("The result is:", CountHowManyWays(choice))
}
