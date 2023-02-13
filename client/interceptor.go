package client

import (
	"context"
	"fmt"
	"log"

	"github.com/worryFree56/grpc_study/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// 客户端拦截器 - 记录请求和响应日志
func ClientUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 前置逻辑
	log.Printf("[Client Interceptor] send request: %s", method)
	reqs := req.(*types.ReqMathv3)

	for _, v := range reqs.GetOper() {
		if _, ok := types.Operation_name[int32(v)]; !ok {
			grpclog.Fatal(fmt.Sprintf("req oper is not exist to operations,arg: %d", v))
		}
	}
	// 发起请求
	err := invoker(ctx, method, req, reply, cc, opts...)

	// 后置逻辑
	log.Printf("[Client Interceptor] response: %s", reply)

	return err
}
