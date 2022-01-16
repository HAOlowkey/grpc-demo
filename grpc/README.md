protobuf语法

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
}

enum Option {
    option allow_alias = true;  //允许别名，允许编号冲突
    OptionA = 0;    // 枚举必须从0开始
    OptionB = 1;
    OptionC = 2;
    OptionD = 0;
}
```