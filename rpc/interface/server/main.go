package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/HAOlowkey/grpc-demo/rpc/interface/service"
)

var _ (service.Service) = (*HelloService)(nil)

type HelloService struct {
}

func (e *HelloService) Hello(req string, resp *string) error {
	*resp = fmt.Sprintf("Hello,%s", req)
	return nil
}

func main() {
	err := rpc.RegisterName(service.Name, new(HelloService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go rpc.ServeConn(conn)
	}
}
