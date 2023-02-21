package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/worryFree56/grpc_study/src/tls/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../../../cert/client.pem", "../../../cert/client.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../../ca/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	// http/1.1
	// cli := http1(cert, certPool)
	// resp, err := cli.Get("https://localhost:3333/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// res, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(res))
	//http/2.0
	creds := http2(cert, certPool, "grpc.ifcfx.com")
	conn, err := grpc.Dial(":3333",
		// grpc.WithInsecure(),
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	grpccli := types.NewHelloClient(conn)

	grpcres, err := grpccli.Say(context.Background(), &types.MsgHello{
		Words: "tls auth!",
	})
	fmt.Println(grpcres, err)
}

// rest
func http1(cert tls.Certificate, certPool *x509.CertPool) *http.Client {
	c := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:            certPool,
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{
		Transport: c,
	}
}

// grpc
func http2(cert tls.Certificate, certPool *x509.CertPool, servername string) credentials.TransportCredentials {
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   servername,
		RootCAs:      certPool,
	})
}
