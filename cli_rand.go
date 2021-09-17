package main

import (
	"fmt"
	"os"
)

var allStudent map[int64]*student

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}
func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号:")
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}
	fmt.Print("请输入姓名:")
	_, nameErr := fmt.Scan(&name)
	if nameErr != nil {
		return
	}
	stu := newStudent(id, name)
	allStudent[id] = stu
}
func deleteStudent() {
	fmt.Print("请输入删除的学号:")
	var id int64
	_, err := fmt.Scanln(&id)
	if err != nil {
		return
	}
	delete(allStudent, id)
}

func main() {
	allStudent = make(map[int64]*student, 20)
	for true {
		fmt.Println(`
  1. 查看所有学生
  2. 添加学生
  3. 删除学生
  4. 退出
 `)
		fmt.Print("请输入:")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("无效的选项")
			return
		}
		fmt.Printf("你选择了:%d\n", choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(9999)
		default:
			fmt.Println("无效的选项")
		}
	}
}
