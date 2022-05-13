package main

import (
	"log"
	"net"
	"net/rpc"
	tr "Practice/rpc"
)




func RegisterHelloService(svc tr.HelloServiceInterface) error {
	return rpc.RegisterName(tr.HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	if err := RegisterHelloService(new(HelloService)); err != nil {
		return
	}
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("Accept error:", err)
		}
		rpc.ServeConn(conn)
	}
}
