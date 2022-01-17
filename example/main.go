package main

import (
	"fmt"

	"github.com/HAOlowkey/grpc-demo/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	// var s1 = protobuf.String{
	// 	Value: "abc",
	// }
	// temp, err := proto.Marshal(&s1)
	// if err != nil {
	// 	panic(err)
	// }

	// var s2 protobuf.String
	// err = proto.Unmarshal(temp, &s2)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(s2)

	// enum
	fmt.Printf("%d\n", grpc.Option_option_c)

	// oneof
	oneOfSample := &grpc.Response{OneofSample: &grpc.Response_Stub2{Stub2: &grpc.Stub2{Value: "stub2"}}}
	fmt.Println(oneOfSample.GetStub1())
	fmt.Println(oneOfSample.GetStub2())

	// any
	any, _ := anypb.New(&grpc.Stub1{Value: "test any stub1"})
	b := grpc.Response{Details: any}
	fmt.Println(b)
	var temp grpc.Stub1
	any.UnmarshalTo(&temp)
	fmt.Println(temp)

	des, _ := anypb.UnmarshalNew(any, proto.UnmarshalOptions{})
	switch v := des.(type) {
	case *grpc.Stub1:
		fmt.Println(v)
	default:
		fmt.Println("unkown")
	}
}
