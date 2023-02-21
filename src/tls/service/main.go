package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/worryFree56/grpc_study/src/tls/core"
	"github.com/worryFree56/grpc_study/src/tls/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../../../cert/server.pem", "../../../cert/server.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../../ca/ca.pem")
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
	gsrv := grpc.NewServer(grpc.Creds(c))

	types.RegisterHelloServer(gsrv, &core.Hello{})
	// lis,err := net.Listen("tcp",":3333")
	// gsrv.Serve(lis)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Proto)
		fmt.Println(r.Header)
		gsrv.ServeHTTP(w, r)
	})
	httpServer := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}
	err = httpServer.ListenAndServeTLS("../../../cert/server.pem", "../../../cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

}
