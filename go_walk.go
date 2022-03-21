package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	visit(t, ch)
	close(ch)
}

func visit(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		visit(t.Left, ch)
	}
	if t.Right != nil {
		visit(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {

	i := 0
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for n := range ch1 {
		fmt.Println(i)
		i ^= n
	}
	fmt.Println()
	for n := range ch2 {
		fmt.Println(i)
		i ^= n
	}
	return i == 0
}

func main() {
	b1 := Same(tree.New(1), tree.New(1))
	b2 := Same(tree.New(2), tree.New(1))

	fmt.Println(b1, b2)
}
