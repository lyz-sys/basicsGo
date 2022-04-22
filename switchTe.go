package main

import "fmt"

func main() {
	var a int
	fmt.Println("请输入一个整数")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("无效的选项")
		return
	}
	//switch 默认true
	switch {
	case a < 10, a > 1000:
		fmt.Printf("%v\n", a-1)
		break
	case a == 10, a == 100:
		fmt.Printf("%v\n", a)
		fallthrough //继续执行下个，不判断条件
	case a > 10 && 1000 > a:
		fmt.Printf("%v\n", a+1)
		break
	}
}
