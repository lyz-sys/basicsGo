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

//token认证
type TokenAuthentication struct {
	AppKey    string
	AppSecret string
}

//组织token信息
func (ta *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  ta.AppKey,
		"appkey": ta.AppSecret,
	}, nil
}

//是否基于TLS认证进行安全传输
func (ta *TokenAuthentication) RequireTransportSecurity() bool {
	return false
}

func main() {
	auth := TokenAuthentication{
		AppKey:    "admin1",
		AppSecret: "123456",
	}
	//creds, err := credentials.NewClientTLSFromFile("/Users/lyz/go/src/demo/microservices/stream_demo/cert/server.pem", "")
	//if err != nil {
	//	panic(err)
	//}
	//, grpc.WithTransportCredentials(creds)
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(&auth))

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
