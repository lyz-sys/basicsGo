package tools

import (
	"math"
	"strconv"
)

// Narcises get narcises
func Narcises(start uint, end uint) (res []uint) {
	if start < 100 {
		start = 100
	}
	if start >= end {
		return
	}
	var tempLen int
	var temp float64
	for i := start; i < end; i++ {
		tempLen = len(strconv.Itoa(int(i)))
		for c := 0; c < tempLen; c++ {
			//fmt.Println(math.Pow(float64(i/uint(math.Pow(10, float64(tempLen-1-c)))%10), float64(tempLen)))
			temp += math.Pow(float64(i/uint(math.Pow(10, float64(tempLen-1-c)))%10), float64(tempLen))
		}
		if temp == float64(i) {
			res = append(res, i)
		}
		temp = 0
	}
	return
}
