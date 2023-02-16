package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/worryFree56/grpc_study/src/mathv3/core"
	"github.com/worryFree56/grpc_study/src/mathv3/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}

	//注册拦截器，获取metadata数据
	s := grpc.NewServer(grpc.UnaryInterceptor(serverUnaryInterceptor))
	//注册grpc服务
	types.RegisterMathV3Server(s, &core.MathV3{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

// 服务端拦截器 -
func serverUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 前置逻辑
	log.Println("进入service 拦截器")

	//通过上下文获取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Printf("metadata to: %v \n", md)
	}

	// 处理请求
	response, err := handler(ctx, req)

	// 后置逻辑
	log.Println("client 请求正在处理")

	return response, err
}
