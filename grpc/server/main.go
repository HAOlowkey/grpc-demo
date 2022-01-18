package main

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/HAOlowkey/grpc-demo/grpc/auth"
	"github.com/HAOlowkey/grpc-demo/grpc/protocol"
	"google.golang.org/grpc"
)

var _ protocol.HelloServiceServer = (*Service)(nil)

type Service struct {
	protocol.UnimplementedHelloServiceServer
}

func (s *Service) Hello(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	return &protocol.Response{Value: "hello: " + req.Value}, nil
}

func (s *Service) Chat(stream protocol.HelloService_ChatServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("chat exit")
				return nil
			}
			return err
		}

		if err = stream.Send(&protocol.Response{Value: "hello: " + req.Value}); err != nil {
			return err
		}
	}
}

func main() {
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(auth.GrpcAuthUnaryServerInterceptor()))

	protocol.RegisterHelloServiceServer(server, new(Service))

	lsr, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server.Serve(lsr)
}
