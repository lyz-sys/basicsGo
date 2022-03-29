package main

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"test-demo/config"
	"test-demo/microservices/stream_demo/message"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:"+strconv.Itoa(config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background())
	if err != nil {
		panic(err)
	}
	orderRequest := message.OrderRequest{OrderId: "201907300001", TimeStamp: time.Now().Unix()}
	err = orderInfo.Send(&orderRequest)
	if err != nil {
		panic(err)
	}
	err = orderInfo.CloseSend()
	if err != nil {
		panic(err)
	}
	info, err := orderInfo.Recv()
	if err == io.EOF {
		fmt.Println("读取结束")
	}
	if err != nil {
		return
	}
	fmt.Println("读取到的信息：", info)
}
