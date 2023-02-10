package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:3333/grpc_study/math/v2/oper/"
	url += "0,1,2,3"
	url += "?a=1&b=2222"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}
