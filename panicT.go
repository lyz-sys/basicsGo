package main

import (
	"demo/logger"
	"errors"
	"fmt"
)

// RecoverPanic is a func to recover panic
func RecoverPanic() {
	if p := recover(); p != nil {
		fmt.Println("panic recover! p:", p)
		str, ok := p.(string)
		if ok {
			logger.Logger.Print(str)
		}
		//debug.PrintStack()
	}
}

func funcA() (err error) {
	defer RecoverPanic()
	return funcB()
}

func funcB() error {
	panic("foo")
	return errors.New("success")
}

func main() {
	err := funcA()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("is nil")
	}
}
