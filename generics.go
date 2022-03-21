package main

import "fmt"

func gen[K comparable, V bool | int](P1 K) V {
	var res V
	//switch res := P1.(type) {
	//case int:
	//	fmt.Println(123, res)
	//default:
	//	fmt.Println(456, res)
	//}
	res = true
	return res
}

func main() {
	fmt.Println(gen(1))
}
