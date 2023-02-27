package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/worryFree56/grpc_study/src/tls/core"
	"github.com/worryFree56/grpc_study/src/tls/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, err := tls.LoadX509KeyPair("./cert/server.pem", "./cert/server.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./ca/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certpool.Appendcertsfrompem  err")
	}
	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	// gsrv := grpc.NewServer(grpc.Creds(c))

	// types.RegisterHelloServer(gsrv, &core.Hello{})
	// lis,err := net.Listen("tcp",":3333")
	// gsrv.Serve(lis)

	// mux := http.NewServeMux()
	mux := runtime.NewServeMux()
	types.RegisterHelloHandlerServer(context.Background(), mux, &core.Hello{})
	err = types.RegisterHelloHandlerFromEndpoint(context.Background(), mux, ":4443", []grpc.DialOption{
		grpc.WithTransportCredentials(c),
	})
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":4444",
		Handler: mux,
	}
	err = httpServer.ListenAndServeTLS("./cert/server.pem", "./cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

}
