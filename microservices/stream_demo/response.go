package main

import (
	"context"
	_ "context"
	_ "errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"test-demo/config"
	"test-demo/microservices/stream_demo/message"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type OrderStruct struct {
	orderMap map[string]message.OrderInfo
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

		result, ok := OrderData.orderMap[orderRequest.GetOrderId()]
		if !ok {
			fmt.Println("请求Id不存在")
			return nil
		}
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
	//creds, err := credentials.NewServerTLSFromFile("/Users/lyz/go/src/demo/microservices/stream_demo/cert/server.pem", "/Users/lyz/go/src/demo/microservices/stream_demo/cert/server.key")
	//if err != nil {
	//	grpclog.Fatal("加载在证书文件失败", err)
	//}
	//grpc.Creds(creds)
	server := grpc.NewServer(grpc.UnaryInterceptor(TokenInterceptor))

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

// 拦截器
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	//通过metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var appKey string
	var appSecret string
	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}
	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "admin" || appSecret != "123456" {
		return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
	}
	//通过token验证，继续处理请求
	return handler(ctx, req)
}
