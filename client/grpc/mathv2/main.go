package main

import (
	"context"
	"log"

	"github.com/worryFree56/grpc_study/types"
	"google.golang.org/grpc"
)

func main() {
	endpoint := "127.0.0.1:3333"

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := types.NewMathV2Client(conn)
	resp, err := client.Operation(context.Background(), &types.ReqMathv2{
		A: 1,
		B: 2,
		Oper: []types.Operation{
			3, 1, 2, 0,
		},
	})
	log.Println(resp.GetResult(), err)
}
