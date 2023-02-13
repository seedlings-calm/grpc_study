package main

import (
	"log"
	"net"

	"github.com/worryFree56/grpc_study/core"
	"github.com/worryFree56/grpc_study/service"
	"github.com/worryFree56/grpc_study/types"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}

	//注册拦截器
	s := grpc.NewServer(grpc.UnaryInterceptor(service.ServerUnaryInterceptor))
	//注册grpc服务
	types.RegisterMathV3Server(s, &core.MathV3{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
