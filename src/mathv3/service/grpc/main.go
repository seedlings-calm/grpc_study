package main

import (
	"log"
	"net"

	"github.com/worryFree56/grpc_study/src/mathv3/core"
	"github.com/worryFree56/grpc_study/src/mathv3/handle"
	"github.com/worryFree56/grpc_study/src/mathv3/types"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}

	//注册拦截器，token安全认证
	s := grpc.NewServer(grpc.UnaryInterceptor(handle.ServerUnaryInterceptor))
	//注册grpc服务
	types.RegisterMathV3Server(s, &core.MathV3{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
