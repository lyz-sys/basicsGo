package main

import (
	_ "demo/config"
	"fmt"
	"math"
)

func shuiXianHua(start int, end int) {
	for i := start; i < end; i++ {
		x := i / 100
		y := i / 10 % 10
		z := i & 10
		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}
}

func sushu(start int, end int) {
	for i := start; i < end; i++ {
		flag := true
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i&j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}

func int2b(num int) int {
	if num > 0 {
		var a int = num % 2
		return int2b(num/2) + a
	}
	return 0
}

func main() {
	//var s string
	//for _, arg := range os.Args[1:] {
	//	if s == "" {
	//		s = arg
	//		continue
	//	}
	//	s += " " + arg
	//}
	//fmt.Println(s)

	//for i := 0; i < 100; i++ {
	//	fmt.Println("sqrt", math.Sqrt(float64(i)))
	//}

	//var f float64
	//fmt.Println("请输入一个整数")
	//通过io判断键盘输入
	//fmt.Println("请输入一个字符串：")
	//reader := bufio.NewReader(os.Stdin)
	//sq,_ := reader.ReadString('\n')
	//fmt.Println(sq)
}
