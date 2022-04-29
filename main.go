package main

import (
	_ "demo/config"
	"math"
	"strconv"
)

// sushu 素数
func sushu(start int, end int) (res []int) {
	if start <= end {
		return
	}
	for i := start; i < end; i++ {
		flag := true
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i&j == 0 {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}
	return
}

// int2b int转2进制
func int2b(num int) string {
	if num > 0 {
		a := num % 2
		return int2b(num/2) + strconv.Itoa(a)
	}
	return ""
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
