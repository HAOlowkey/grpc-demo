package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

	var rep string
	err = client.Call("HelloService.Hello", "Alice", &rep)
	if err != nil {
		panic(err)
	}

	fmt.Println(rep)
}
