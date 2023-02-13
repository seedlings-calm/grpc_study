package service

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func ServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 前置逻辑
	log.Printf("[Server Interceptor] accept request: %s", info.FullMethod)

	// 处理请求
	response, err := handler(ctx, req)

	// 后置逻辑
	log.Printf("[Server Interceptor] response: %s", response)

	return response, err
}
