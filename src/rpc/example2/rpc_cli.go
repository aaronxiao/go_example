package main

import (
	"net/rpc"
	"log"
	"fmt"
	tr "Practice/rpc"
)



func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call(tr.HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
