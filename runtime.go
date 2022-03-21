package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//获取goroot目录：
	fmt.Println("GOROOT-->", runtime.GOROOT())

	//获取操作系统
	fmt.Println("os/platform-->", runtime.GOOS) // GOOS--> darwin，mac系统

	//1.获取逻辑cpu的数量
	fmt.Println("逻辑CPU的核数：", runtime.NumCPU())
	//2.设置go程序执行的最大的：[1,256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine。。。")
		}
	}()

	for i := 0; i < 4; i++ {
		//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("main。。")
	}

	go func() {
		fmt.Println("goroutine开始。。。")

		//调用了别的函数
		fun()

		fmt.Println("goroutine结束。。")
	}()

	time.Sleep(3 * time.Second)
}

func fun() {
	defer fmt.Println("defer。。。")

	//return           //终止此函数
	runtime.Goexit() //终止所在的协程

	fmt.Println("fun函数。。。")
}
