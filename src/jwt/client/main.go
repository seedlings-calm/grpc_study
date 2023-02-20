package main

import (
	"context"
	"fmt"
	"log"

	"github.com/worryFree56/grpc_study/src/jwt/auth"
	"github.com/worryFree56/grpc_study/src/jwt/types"
	"google.golang.org/grpc"
)

func main() {
	creds := auth.JwtAuthToken{
		Token: auth.CreateJwtToken("jwt"),
	}
	conn, err := grpc.Dial(":3333",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(creds),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cli := types.NewHelloClient(conn)
	res, err := cli.Say(context.Background(), &types.MsgHello{
		Words: "jwt token!",
	})
	fmt.Println(res, err)
}
