package client

import (
	"context"

	"google.golang.org/grpc"
)

// 客户端拦截器 - 记录请求和响应日志
func ClientUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 前置逻辑
	// log.Printf("[Client Interceptor] send request: %s", method)

	// 发起请求
	err := invoker(ctx, method, req, reply, cc, opts...)

	// 后置逻辑
	// log.Printf("[Client Interceptor] response: %s", reply)

	return err
}
