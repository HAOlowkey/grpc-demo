protobuf语法

```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install 
```

```
syntax = "proto3";      //使用protobuf3的语法解析

import "google//any"    //导入外部包 需要-I时指定

package demo.grpc.hello;

option go_package="github.com/HAOlowkey/grpc-demo";

// field_rule 有4种 optional repeated oneof any
message Response {      //定义结构体 字段要用蛇形缩写 转换时会转成驼峰
    reserved 10 , 13 to 15;     //保留编号
    reserved "foo","bar";       //保留字段
    string value = 1;
    optional string key = 2;   //会用字符串指针表示 允许为nil  key="", 加了optional 可以用enabled != nil { 添加过滤条件 }
    bool enabled = 3;
    int32 page_number = 4;
    float
    double metric_value = 6;  // 转换成float64
    bytes                     // 转换成[]byte
    repeated Result results = 8;   // 元素类型为result的数组
    map <string, Project> project = 9; // 定义map
    Result result = 10;     // 结构体嵌套 如果夸包应用 需要先import 并且指定类型时要加上包的名称
    oneof one_sample {      // 类似于泛形 等同于枚举但是值的类型是个复杂结构 
        Stub1 stub1 = 11;
        Stub2 stub2 = 12;
    }
    /*
    使用时
    oneOfSample := grpc.Response{OneofSample: &grpc.Response_Stub2{Stub2: &grpc.Stub2{Value: "stub2"}}}
	fmt.Println(oneOfSample.GetStub1())
    fmt.Println(oneOfSample.GetStub2())
    */
    google.protobuf.Any details = 13;         // 需要先import "google/protobuf/any.proto"; 然后在生成代码时通过-I /usr/local/include 导入
    /*
    使用时如果知道类型 客户端行为
    any, _ := anypb.New(&grpc.Stub1{Value: "test any stub1"})
	var temp grpc.Stub1
	any.UnmarshalTo(&temp)
	fmt.Println(temp)

    不知道类型时 服务端行为
    des, _ := anypb.UnmarshalNew(any, proto.UnmarshalOptions{})
	switch v := des.(type) {
	case *grpc.Stub1:
		fmt.Println(v)
	}
    */
}

message Stub1 {
    string name = 1;
}

message Stub2 {
    string name = 1;
}

message Result {
    string value = 1;
}

enum Option {
    option allow_alias = true;  //允许别名，允许编号冲突
    OptionA = 0;    // 枚举必须从0开始
    OptionB = 1;
    OptionC = 2;
    OptionD = 0;
}
```