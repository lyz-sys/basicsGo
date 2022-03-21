package main

import (
	"errors"
	"fmt"
)

func funcA() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			//debug.PrintStack()
		}
	}()
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
