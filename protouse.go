package main

import (
	"fmt"
	"os"
	"test-demo/example"

	"github.com/golang/protobuf/proto"
)

func main() {
	msg_test := &example.Person{
		Name: proto.String("Davie"),
		Age:  proto.Int(18),
		From: proto.String("China"),
	}

	//序列化
	msgDataEncoding, err := proto.Marshal(msg_test)
	if err != nil {
		panic(err.Error())
		return
	}
	//fmt.Println(msgDataEncoding)
	//return
	//var msgDataEncoding = []byte{10, 5, 68, 97, 118, 105, 101, 16, 18, 26, 5, 67, 104, 105, 110, 97}

	msgEntity := example.Person{}
	//反序列化
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Printf("姓名：%s\n", *msgEntity.Name)
	fmt.Printf("年龄：%d\n", msgEntity.GetAge())
	fmt.Printf("国籍：%s\n", msgEntity.GetFrom())
}
