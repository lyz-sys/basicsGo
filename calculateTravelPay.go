package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// travelFee Calculate the commute cost for a period of time
func travelFee(subOncePay float64, busOncePay float64, day int) (total float64) {
	if day == 0 || (subOncePay == 0 && busOncePay == 0) {
		return
	}
	times := 2 * day
	decimal.DivisionPrecision = 2
	twOff := decimal.NewFromFloat(subOncePay * 0.8)
	FiOff := decimal.NewFromFloat(subOncePay * 0.5)
	for i := 0; i < times; i++ {
		if total < 100 {
			total += subOncePay
		} else if total < 150 {
			total, _ = decimal.NewFromFloat(total).Add(twOff).Float64()
		} else {
			total, _ = decimal.NewFromFloat(total).Add(FiOff).Float64()
		}
	}
	total, _ = decimal.NewFromFloat(total).Add(decimal.NewFromFloat(float64(times) * busOncePay)).Float64()
	return
}

func main() {
	payMoney := travelFee(6.00, 1.00, 22)
	fmt.Println(payMoney)
}
