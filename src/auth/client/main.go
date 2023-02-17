package main

import (
	"context"
	"fmt"
	"log"

	"github.com/worryFree56/grpc_study/src/auth/auth"
	"github.com/worryFree56/grpc_study/src/auth/types"
	"google.golang.org/grpc"
)

func main() {
	//自验证数据
	at := &auth.AuthToken{
		Name:     "admin",
		Password: "1234567",
	}
	conn, err := grpc.Dial(":3333",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		//使用自定义验证
		grpc.WithPerRPCCredentials(at),
	)
	if err != nil {
		log.Fatalf("failed to dial:%s", err)
	}
	defer conn.Close()
	cli := types.NewAuthTokenClient(conn)

	sres, err := cli.SayMsg(context.Background(), &types.SayMessage{Export: "say hello world"})
	fmt.Println(sres, err)

}
