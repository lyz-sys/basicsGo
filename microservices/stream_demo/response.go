package main

import (
	_ "context"
	_ "errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"test-demo/config"
	"test-demo/microservices/stream_demo/message"

	"google.golang.org/grpc"
)

type OrderStruct struct {
	orderMap map[string]message.OrderInfo
	mux      sync.Mutex
}

//订单服务实现
type OrderServiceImpl1 struct {
}

//获取订单信息s
func (os *OrderServiceImpl1) GetOrderInfo(stream message.OrderService_GetOrderInfoServer) error {
	OrderData := new(OrderStruct)
	OrderData.orderMap = map[string]message.OrderInfo{
		"201907300001": {OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": {OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": {OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}

	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(" 数据读取结束 ")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}

		fmt.Println("请求：", orderRequest.GetOrderId())
		fmt.Println("请求：", orderRequest)

		OrderData.mux.Lock()
		result := OrderData.orderMap[orderRequest.GetOrderId()]
		OrderData.mux.Unlock()
		//发送数据
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func main() {
	server := grpc.NewServer()

	message.RegisterOrderServiceServer(server, new(OrderServiceImpl1))

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		return
	}
}
