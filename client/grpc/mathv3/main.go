package main

import (
	"context"
	"fmt"
	"log"

	"github.com/worryFree56/grpc_study/client"
	"github.com/worryFree56/grpc_study/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial(":3333", grpc.WithInsecure(), grpc.WithUnaryInterceptor(client.ClientUnaryInterceptor))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//自定义header
	md := metadata.New(map[string]string{
		"name":     "client-name",
		"nickname": "client-nickname",
	})
	//初始化一个 新的context
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//追加内容
	ctx = metadata.AppendToOutgoingContext(ctx, "description", "This is desc")

	c := types.NewMathV3Client(conn)
	var resmd metadata.MD
	r, err := c.Operation(ctx, &types.ReqMathv3{
		A: "1",
		B: "2",
		Oper: []types.Operation{
			1, 2, 3,
		},
	}, grpc.Header(&resmd))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("code:%d !\n", r.Code)
	fmt.Printf("result:%+v !\n", r.Result)
}
