package main

import (
	"context"
	"fmt"
	"time"

	"github.com/HAOlowkey/grpc-demo/grpc/protocol"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := protocol.NewHelloServiceClient(conn)

	resp, err := client.Hello(context.Background(), &protocol.Request{Value: "Bob"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	stream, err := client.Chat(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			if err = stream.Send(&protocol.Request{Value: "Alice"}); err != nil {
				fmt.Println(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		resp, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp.Value)
	}
}
