## mathv1

### 实现数学加减乘除

### 源码目录
```
|-- src
    |-- readme.md
    |-- client
    |   |-- grpc
    |       |-- main.go
    |-- core
    |   |-- mathv1.go
    |-- proto
    |   |-- mathv1.proto
    |-- service
    |   |-- grpc
    |       |-- main.go
    |-- types
        |-- mathv1.pb.go
```
### 完成protobuf基础编写，转成golang代码

```proto
//github.com/worryFree56/grpc_study/src/mathv1/proto
syntax = "proto3";

package grpc_study.math;

option go_package = "github.com/worryFree56/grpc_study/src/mathv1/types";


message ReqMathv1 {
    int64 a = 1;
    int64 b = 2;
}

message ResMathv1 {
    int64 res = 1;   
}

service MathV1 {
    rpc Add (ReqMathv1) returns (ResMathv1);
    rpc Sub (ReqMathv1) returns (ResMathv1);
    rpc Mul (ReqMathv1) returns (ResMathv1);
    rpc Div (ReqMathv1) returns (ResMathv1);
}
  
```
### 项目根目录（grpc_study/）编译proto文件为golang文件
```shell
protoc \
 -I=src  \
 --go_out=plugins=grpc:. \
 ./src/proto/mathv1.proto
```

### protoc命令的 `.` 表示当前目录，执行protoc命令生成的`*.pb.go`文件 会在 以当前目录生成`go_package`指定目录里面，以下命令把编译文件移到go_package目录

```
mv github.com/worryFree56/grpc_study/src/mathv1/types/* src/mathv1/types/
rm -rf github.com/  
```
### 实现加减乘除接口逻辑
以下代码，定义一个结构体，实现types.MathV1Server的接口
```golang
package core

import (
	"context"

	"github.com/worryFree56/grpc_study/src/mathv1/types"
)

type MathV1 struct{}

var _ types.MathV1Server = new(MathV1)

func (MathV1) Sub(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
	return &types.ResMathv1{
		Res: req.GetA() - req.GetB(),
	}, nil
}
func (MathV1) Add(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
	return &types.ResMathv1{
		Res: req.GetA() + req.GetB(),
	}, nil
}
func (MathV1) Mul(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
	return &types.ResMathv1{
		Res: req.GetA() * req.GetB(),
	}, nil
}
func (MathV1) Div(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
	return &types.ResMathv1{
		Res: req.GetA() / req.GetB(),
	}, nil
}
```

### 具体功能实现了之后，就可以开启service，client调用了
- grpc 服务启动
```golang
package main

import (
	"log"
	"net"

	"github.com/worryFree56/grpc_study/src/mathv1/core"
	"github.com/worryFree56/grpc_study/src/mathv1/types"
	"google.golang.org/grpc"
)

func main() {
	endpoint := "127.0.0.1:3333"

	grpcsrv := grpc.NewServer()

	types.RegisterMathV1Server(grpcsrv, core.MathV1{})

	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	err = grpcsrv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
```
- 客户端调用
```go
package main

import (
	"context"
	"log"

	"github.com/worryFree56/grpc_study/src/mathv1/types"
	"google.golang.org/grpc"
)

func main() {
	endpoint := "127.0.0.1:3333"

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := types.NewMathV1Client(conn)
    //功能调用
    resp, err := client.Add(context.Background(), &types.ReqMathv1{
		A: 1,
		B: 2,
	})
	log.Println(resp, err)
}

```
