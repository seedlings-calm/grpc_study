package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/worryFree56/grpc_study/src/jwt/auth"
	"github.com/worryFree56/grpc_study/src/jwt/core"
	"github.com/worryFree56/grpc_study/src/jwt/types"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}

	gsrv := grpc.NewServer(grpc.UnaryInterceptor(serviceUnaryServerInterceptor))
	types.RegisterHelloServer(gsrv, &core.Hello{})

	err = gsrv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func serviceUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	ok, err := auth.IsValidJWToken(ctx)
	if !ok {
		return nil, errors.New("valid err:" + err.Error())
	}

	res, err := handler(ctx, req)

	return res, err
}
