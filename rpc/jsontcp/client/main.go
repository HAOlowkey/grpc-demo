package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/HAOlowkey/grpc-demo/rpc/jsontcp/service"
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
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (h *HelloServiceClient) Hello(req string, resp *string) error {
	return h.client.Call(service.Name+".Hello", req, resp)
}

func main() {
	client, err := NewHelloServiceClient("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	var resp string
	err = client.Hello("HAHAHAHA", &resp)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
