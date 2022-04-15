package main

import (
	_ "demo/config"
	"errors"
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

const (
	c string = "常量字符"
	d
	e = iota
	f
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

func formal(nums int, len int) (int, int, error) {
	var i int
out:
	for ; i < len; i++ {
		if nums += 10; len == 10 {
			continue out
		}
	}
	return nums, i, nil
}

func int2b(num int) int {
	if num > 0 {
		var a int = num % 2
		return int2b(num/2) + a
	}
	return 0
}

func bubbleSort(arr []int) ([]int, error) {
	if arr == nil {
		return nil, errors.New("slice is nil")
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr, nil
}

func travelFee(oncePay float64, day int) float64 {
	var total float64
	day *= 2
	decimal.DivisionPrecision = 2
	for i := 0; i < day; i++ {
		if total < 100 {
			total += oncePay
		} else if total < 150 {
			total, _ = decimal.NewFromFloat(total).Add(decimal.NewFromFloat(oncePay * 0.8)).Float64()
		} else {
			total, _ = decimal.NewFromFloat(total).Add(decimal.NewFromFloat(oncePay * 0.5)).Float64()
		}
	}

	return total
}

func main() {
	//monthtravelFee := travelFee(7.00, 22)
	//fmt.Println(monthtravelFee)

	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)
	//fmt.Println("-------------------------------------")
	//var name string = "name"
	//name = "123123aaa"
	//var intc int = 123
	//fmt.Println(intc)
	//fmt.Println(c)
	//fmt.Println(name)

	//var oldArr = []int{9, 5, 8, 4, 2, 3, 1, 99}
	////array := [...] int{9,5,8,4,2,3,1,99}
	//arr, _ := bubbleSort(oldArr)
	//
	//fmt.Println(arr)

	//var map1 = map[string]int{"go":1}
	//fmt.Println(map1)
	//
	//var arr = [3]int {1,2,3}
	//var selice = arr[:]
	//selice = append(selice,4,5)
	//selice[0] = 9
	//fmt.Println(selice)
	//fmt.Println(arr)

	//fmt.Println(array)
	//for index,value := range array {
	//	fmt.Printf("下标是：%d,value:%d\n", index, value)
	//}

	//for i := 0; i < 100; i++ {
	//	fmt.Println("sqrt", int(math.Sqrt(float64(i))))
	//}
	//
	//a, b, _ := formal(100, 200)
	//fmt.Println(a, b)
	//
	//p := int2b(14)
	//fmt.Println(p)
	//fmt.Printf("%T,常量c=%s\n",c,c)
	//fmt.Printf("%T,常量e=%d\n",e,e)
	//fmt.Printf("%T,常量f=%d\n",f,f)
	//if a > 10 {
	//	fmt.Printf("%v", a)
	//}
	//

	//var f float64
	//fmt.Println("请输入一个整数")
	//通过io判断键盘输入
	//fmt.Println("请输入一个字符串：")
	//reader := bufio.NewReader(os.Stdin)
	//sq,_ := reader.ReadString('\n')
	//fmt.Println(sq)
}
