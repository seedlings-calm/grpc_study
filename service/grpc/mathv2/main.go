package main

import (
	"log"
	"net"

	"github.com/worryFree56/grpc_study/core"
	"github.com/worryFree56/grpc_study/types"
	"google.golang.org/grpc"
)

func main() {
	endpoint := "127.0.0.1:3333"

	grpcsrv := grpc.NewServer()

	types.RegisterMathV2Server(grpcsrv, core.MathV2{})

	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	err = grpcsrv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
