package main

import (
	"fmt"

	"github.com/HAOlowkey/grpc-demo/protobuf"
	"google.golang.org/protobuf/proto"
)

func main() {
	var s1 = protobuf.String{
		Value: "abc",
	}
	temp, err := proto.Marshal(&s1)
	if err != nil {
		panic(err)
	}

	var s2 protobuf.String
	err = proto.Unmarshal(temp, &s2)
	if err != nil {
		panic(err)
	}

	fmt.Println(s2)
}
