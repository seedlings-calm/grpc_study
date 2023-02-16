package handle

import (
	"context"
	"fmt"

	v2types "github.com/worryFree56/grpc_study/src/mathv2/types"
	v3types "github.com/worryFree56/grpc_study/src/mathv3/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var ReqCount = make(map[string]int)

func ServerUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	// 前置逻辑
	reqs := req.(*v3types.ReqMathv3)

	for _, v := range reqs.GetOper() {
		if _, ok := v2types.Operation_name[int32(v)]; !ok {
			grpclog.Fatal(fmt.Sprintf("req oper is not exist to operations,arg: %d", v))
		}
	}

	// 处理请求
	response, err := handler(ctx, req)

	// 后置逻辑
	if _, ok := ReqCount[info.FullMethod]; !ok {
		ReqCount[info.FullMethod] = 1
	} else {
		ReqCount[info.FullMethod]++
	}

	return response, err
}
