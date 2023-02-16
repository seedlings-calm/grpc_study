package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/worryFree56/grpc_study/client"
	"github.com/worryFree56/grpc_study/core"
	"github.com/worryFree56/grpc_study/types"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	mux := runtime.NewServeMux()
	types.RegisterMathV3HandlerServer(ctx, mux, core.MathV3{})
	options := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(client.ClientUnaryInterceptor),
	}
	err = types.RegisterMathV3HandlerFromEndpoint(ctx, mux, "127.0.0.1:4444", options)
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Handler: mux,
	}

	err = srv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
