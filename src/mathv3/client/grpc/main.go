package main

import (
	"context"
	"fmt"
	"log"

	v2types "github.com/worryFree56/grpc_study/src/mathv2/types"
	v3types "github.com/worryFree56/grpc_study/src/mathv3/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial(":3333",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(clientUnaryInterceptor),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	md := metadata.New(map[string]string{
		"description": "This is client",
		"name":        "client-name",
	})
	//初始化一个 新的context
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//追加内容
	ctx = metadata.AppendToOutgoingContext(ctx, "name", "client-name-alias")

	c := v3types.NewMathV3Client(conn)

	r, err := c.Operation(ctx, &v3types.ReqMathv3{
		A: "1",
		B: "2",
		Oper: []v2types.Operation{
			1, 2, 3,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("code:%d \n", r.Code)
	fmt.Printf("result:%v \n", r.Result)
}

// 客户端拦截器 -
func clientUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 前置逻辑
	log.Printf("进入client 拦截器 \n")

	// 发起请求
	err := invoker(ctx, method, req, reply, cc, opts...)

	// 后置逻辑
	log.Printf("client 请求已经发起 \n")

	return err
}
