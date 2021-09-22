package main

import (
	"fmt"
	"time"
)

type intnum uint64

func (n intnum) t() intnum {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return (n - 1).t() + (n - 2).t()
}

func main() {
	var choice intnum
	fmt.Print("请输入台阶数量:")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("无效的选项")
		return
	}
	fmt.Printf("你选择了:%d\n", choice)
	fmt.Println(time.Now())
	fmt.Println(choice.t())
	fmt.Println(time.Now())
}
