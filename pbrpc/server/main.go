package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/HAOlowkey/grpc-demo/pbrpc/codec/server"
	"github.com/HAOlowkey/grpc-demo/pbrpc/service"
)

var _ (service.Service) = (*HelloService)(nil)

type HelloService struct {
}

func (e *HelloService) Hello(req *service.Request, resp *service.Response) error {
	resp.Value = fmt.Sprintf("Hello,%s", req.Value)
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

		// go rpc.ServeConn(conn)
		// go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		go rpc.ServeCodec(server.NewServerCodec(conn))
	}
}
