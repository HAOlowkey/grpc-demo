package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/HAOlowkey/grpc-demo/pbrpc/codec/client"
	"github.com/HAOlowkey/grpc-demo/pbrpc/service"
)

var _ service.Service = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (service.Service, error) {
	// client, err := rpc.Dial(network, address)
	// if err != nil {
	// 	return nil, err
	// }
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	// client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	client := rpc.NewClientWithCodec(client.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (h *HelloServiceClient) Hello(req *service.Request, resp *service.Response) error {
	return h.client.Call(service.Name+".Hello", req, resp)
}

func main() {
	client, err := NewHelloServiceClient("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	req := service.Request{Value: "HAHAHAHA"}
	var resp service.Response
	err = client.Hello(&req, &resp)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
