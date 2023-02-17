package main

import (
	"context"
	"log"
	"net"

	"github.com/worryFree56/grpc_study/src/auth/auth"
	"github.com/worryFree56/grpc_study/src/auth/core"
	"github.com/worryFree56/grpc_study/src/auth/types"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	gsrv := grpc.NewServer(grpc.UnaryInterceptor(authTokenInterceptor))

	types.RegisterAuthTokenServer(gsrv, &core.Auth{})

	err = gsrv.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve:%s", err)
	}

}

// auth token 拦截器
func authTokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 前置逻辑
	log.Println("进入service 拦截器")

	err = auth.IsValidAuthToken(ctx)
	if err != nil {
		return
	}
	// 处理请求
	response, err := handler(ctx, req)

	// 后置逻辑
	log.Println("client 请求正在处理")

	return response, err
}
