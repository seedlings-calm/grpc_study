## mathv3

使用拦截器，可以实现记录日志，身份验证，统计数据等操作


### 源码目录

```
|-- src
    |-- readme.md
    |-- client
    |   |-- grpc
    |       |-- main.go
    |-- core
    |   |-- mathv3.go
    |-- proto
    |   |-- mathv3.proto
    |-- service
    |   |-- grpc
    |       |-- main.go
    |-- types
        |-- mathv3.pb.go

```

### proto 定义，嵌套message，引入自定义mathv2.proto的 `enum Operation`
```proto
syntax = "proto3";

package grpc_study.math;

import "mathv2/proto/mathv2.proto";

option go_package = "github.com/worryFree56/grpc_study/src/mathv3/types";

message ReqMathv3 {

    string a = 1 ;
    string b = 2 ;
    repeated Operation oper = 3 ;
}

message ResMathv3 {
    int64 code = 1;
    repeated Res result = 2;
}

message Res {
    string type  = 1;
    string value = 2;
}

service MathV3 {
    rpc Operation (ReqMathv3) returns (ResMathv3);
}
```

### 项目根目录（grpc_study/）编译文件
```shell
 protoc \
    -I=third_party \
    -I=src  \
    --go_out=plugins=grpc:. \
    ./src/mathv3/proto/mathv3.proto
```
```shell
 mv github.com/worryFree56/grpc_study/src/mathv3/types/* src/mathv3/types/
 rm -rf github.com 
```

### 定义拦截器
- 服务端拦截器
初始化一个grpc服务时需要以option `grpc.UnaryInterceptor()`的方式添加拦截器,参数格式是 `type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)`

```go
s := grpc.NewServer(grpc.UnaryInterceptor(serverUnaryInterceptor))
```

示例代码
```go
// 服务端拦截器 -
func serverUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 前置逻辑
	log.Printf("进入service 拦截器")

	// 处理请求
	response, err := handler(ctx, req)

	// 后置逻辑
	log.Printf("client 请求正在处理")

	return response, err
}
```
- 客户端拦截器

创建一个grpc的conn时，需要以option `grpc.WithUnaryInterceptor` 来加载拦截器，参数格式是 `type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error`

```go
	conn, err := grpc.Dial(":3333",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(clientUnaryInterceptor)
	)
```
示例代码
```go
// 客户端拦截器 
func clientUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 前置逻辑
	log.Printf("进入client 拦截器 \n")

	// 发起请求
	err := invoker(ctx, method, req, reply, cc, opts...)

	// 后置逻辑
	log.Printf("client 请求已经发起 \n")

	return err
}
```

### 定义metadata参数
gRPC 通过 `google.golang.org/grpc/metadata` 包内的 `MD` 类型提供相关的功能接口
```go
// MD is a mapping from metadata keys to values. Users should use the following
// two convenience functions New and Pairs to generate MD.
type MD map[string][]string
```

#### 客户端定义发送
客户端请求的 metadata 是通过设置 context 使用的，metadata 包提供了两个 context 相关的方法，设置好 context 后直接在调用 rpc 方法时传入即可
```go
	md := metadata.New(map[string]string{
		"name": "client-name",
	})
	//初始化一个 新的context
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//追加内容
	ctx = metadata.AppendToOutgoingContext(ctx, "description", "This is client")

```


#### 服务端定义接受
对应客户端请求的 metadata 是使用 context 设置的，那么服务端在接收时自然也是从 context 中读取，metadata 包中的 FromIncommingContext 方法就是用来读取 context 中的 metadata数据的,`以下代码植入拦截器或者逻辑方法中都可获取数据`
```go
	//通过上下文获取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Printf("metadata to: %v \n", md)
	}
```

### 启动服务和调用

服务端启动
```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/worryFree56/grpc_study/src/mathv3/core"
	"github.com/worryFree56/grpc_study/src/mathv3/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}

	//注册拦截器，获取metadata数据
	s := grpc.NewServer(grpc.UnaryInterceptor(serverUnaryInterceptor))
	//注册grpc服务
	types.RegisterMathV3Server(s, &core.MathV3{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

// 服务端拦截器 -
func serverUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 前置逻辑
	log.Println("进入service 拦截器")

	//通过上下文获取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Printf("metadata to: %v \n", md)
	}

	// 处理请求
	response, err := handler(ctx, req)

	// 后置逻辑
	log.Println("client 请求正在处理")

	return response, err
}

```
客户端请求
```go
package main

import (
	"context"
	"fmt"
	"log"

	v2types "github.com/worryFree56/grpc_study/src/mathv2/types"
	v3types "github.com/worryFree56/grpc_study/src/mathv3/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial(":3333",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(clientUnaryInterceptor),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	md := metadata.New(map[string]string{
		"description": "This is client",
		"name":        "client-name",
	})
	//初始化一个 新的context
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//追加内容
	ctx = metadata.AppendToOutgoingContext(ctx, "name", "client-name-alias")

	c := v3types.NewMathV3Client(conn)

	r, err := c.Operation(ctx, &v3types.ReqMathv3{
		A: "1",
		B: "2",
		Oper: []v2types.Operation{
			1, 2, 3,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("code:%d \n", r.Code)
	fmt.Printf("result:%v \n", r.Result)
}

// 客户端拦截器 -
func clientUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 前置逻辑
	log.Printf("进入client 拦截器 \n")

	// 发起请求
	err := invoker(ctx, method, req, reply, cc, opts...)

	// 后置逻辑
	log.Printf("client 请求已经发起 \n")

	return err
}

```