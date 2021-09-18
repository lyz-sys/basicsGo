package main

import "fmt"

type Interf interface {
	M()
}

type Ts struct {
	S string
}

// M 此方法表示类型 Ts 实现了接口 Interf，但我们无需显式声明此事。
func (t Ts) M() {
	fmt.Println(t.S)
}

func main() {
	var i Interf = Ts{"hello"}
	i.M()
}
