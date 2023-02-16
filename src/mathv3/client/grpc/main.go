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
	)
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

	c := v3types.NewMathV3Client(conn)

	// var resmd metadata.MD
	r, err := c.Operation(ctx, &v3types.ReqMathv3{
		A: "1",
		B: "2",
		Oper: []v2types.Operation{
			1, 2, 3,
		},
	}, grpc.Header(&md))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("code:%d !\n", r.Code)
	fmt.Printf("result:%+v !\n", r.Result)
}
