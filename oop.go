package main

import (
	"fmt"
)

//定义父级的数据结构（Pet）
type Pet struct {
}

//父级有一个speak()方法
func (pet *Pet) speak() {
	fmt.Println("Pet的speak()方法")
}

//父级有还有一个speakTo()方法，它调用了父级的 speak() 方法
func (pet *Pet) speakTo() {
	pet.speak()
	fmt.Println("Pet的其它代码")
}

//定义子集的数据结构（Dog）
type Dog struct {
	//Go 语言提供了一种匿名嵌套类型，可以直接声明父级的数据结构
	Pet
}

//在Dog中重写了Pet的speak()方法
func (dog *Dog) speak() {
	fmt.Println("Dog的speak()方法")
}

//测试 Go 是否支持继承
//根据 LSP 原则，调用父类中的方法，如果子类中重写了父类的方法，那么会自动调用子类中的同名函数。
//如果 Go 支持继承，那么此处的输出应该为Dog的speak()方法输出的结果。
//根据代码运行结果，我们得知 Go 不支持继承
func TestClient() {
	dog := new(Dog)
	dog.speakTo()
}
func main() {
	TestClient()
}
