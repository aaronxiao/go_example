package rpc

const HelloServiceName = "path/to/pkg.HelloService"


type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}