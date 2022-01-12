package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(req string, resp *string) error {
	*resp = fmt.Sprintf("hello, %s", req)
	return nil
}

func main() {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", "0.0.0.0:8090")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(conn)
	}
}
