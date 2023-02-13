package main

import (
	"context"
	"fmt"
	"log"

	"github.com/worryFree56/grpc_study/client"
	"github.com/worryFree56/grpc_study/types"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":3333", grpc.WithInsecure(), grpc.WithUnaryInterceptor(client.ClientUnaryInterceptor))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := types.NewMathV3Client(conn)
	r, err := c.Operation(context.Background(), &types.ReqMathv3{
		A: "1",
		B: "2",
		Oper: []types.Operation{
			1, 2, 3,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("code:%d !\n", r.Code)
	fmt.Printf("result:%+v !\n", r.Result)
}
