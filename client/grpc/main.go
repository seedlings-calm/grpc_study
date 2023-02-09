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
	client := types.NewMathV1Client(conn)
	resp, err := client.Add(context.Background(), &types.ReqMath{
		A: 1,
		B: 2,
	})
	log.Println(resp, err)
}
