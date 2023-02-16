package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// url := "http://127.0.0.1:3333/grpc_study/math/v3/oper/"
	// url += "2"
	// url += "?a=1&b=2222"
	url1 := "http://127.0.0.1:3333/grpc_study/math/v3/oper/10/0,1/10"
	resp, err := http.Get(url1)
	if err != nil {
		log.Fatal(err)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}
