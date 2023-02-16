package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/worryFree56/grpc_study/src/mathv2/core"
	"github.com/worryFree56/grpc_study/src/mathv2/types"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	types.RegisterMathV2HandlerServer(ctx, mux, core.MathV2{})
	err := types.RegisterMathV2HandlerFromEndpoint(ctx, mux, "localhost:3002", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Handler: mux,
	}
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}

	err = srv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
