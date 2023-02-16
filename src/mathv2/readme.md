## mathv2
定义一个公用操作方法，根据enum参数来决定执行什么方法,实现`grpc`的服务和客户端调用，实现`restful`的服务和调用

### 实例目录
|-- src
    |-- readme.md
    |-- client
    |   |-- grpc
    |   |   |-- main.go
    |   |-- rest
    |       |-- main.go
    |-- core
    |   |-- mathv2.go
    |-- proto
    |   |-- mathv2.proto
    |-- service
    |   |-- grpc
    |   |   |-- main.go
    |   |-- rest
    |       |-- main.go
    |-- types
        |-- mathv2.pb.go
        |-- mathv2.pb.gw.go
### proto文件定义
```proto
//github.com/worryFree56/grpc_study/src/mathv2/proto
syntax = "proto3";

package grpc_study.math;

//引入third_party文件夹文件，
//protoc时 需要带上 -I "third_party/"
import  "google/api/annotations.proto";

option go_package =  "github.com/worryFree56/grpc_study/src/mathv2/types";

message ReqMathv2 {
    int64 a = 1;
    int64 b = 2;

    repeated Operation oper = 11;
    /*
    保留以后要用的编号和变量名，如果 定义了reserved，然后
    int64  c = 11; 或者 int64 e=6;  编译都会报错
    */
    reserved 3,4,5 to 10;
    reserved "c","d";
}
//定义包含的操作项
enum Operation {
    Add  = 0;
    Sub = 1;
    Mul = 2;
    Div = 3;
}

message ResMathv2 {
    map<string,string> result = 1;   
}


service MathV2 {
    rpc Operation (ReqMathv2) returns (ResMathv2) {
        // http:localhost:port/grpc_study/math/v2/oper/0,2,3?a=1&b=2
        option (google.api.http).get = "/grpc_study/math/v2/oper/{oper}";
    }
}
```

### 项目根目录（grpc_study/）编译文件，支持restful，编译命令需要加上`--grpc-gateway_out`参数

```shell
protoc \
    -I=src  \
    -I=third_party \
    --go_out=plugins=grpc:. \
    --grpc-gateway_out=logtostderr=true:. \
    ./src/mathv2/proto/mathv2.proto
```

### protoc命令的 `.` 表示当前目录，执行protoc命令生成的`*.pb.go`,`*.pb.gw.go`文件 会在 以当前目录生成`go_package`指定目录里面，以下命令把编译文件移到go_package目录

```
mv github.com/worryFree56/grpc_study/src/mathv2/types/* src/mathv2/types/
rm -rf github.com/  
```

### 实现 `Operation` 方法

实现中，使用`github.com/shopsring/decimal`仓库作为方法的技术，根据enum的参数和`reflect`来锚定调用的具体方法

```go
package core

import (
	"context"
	"reflect"

	"github.com/shopspring/decimal"
	"github.com/worryFree56/grpc_study/src/mathv2/types"
)

type MathV2 struct{}

var _ types.MathV2Server = new(MathV2)

func (MathV2) Operation(ctx context.Context, req *types.ReqMathv2) (res *types.ResMathv2, err error) {
	a := req.GetA()
	b := req.GetB()
	var operRes = make(map[string]string)

	for _, v := range req.GetOper() {
		adeci := decimal.NewFromInt(a)
		bdeci := decimal.NewFromInt(b)

		oper := v.String()

		value := reflect.ValueOf(adeci)
		f := value.MethodByName(oper) //通过反射获取它对应的函数，然后通过call来调用

		params := make([]reflect.Value, 1) //参数

		params[0] = reflect.ValueOf(bdeci) //参数设置为b

		rs := f.Call(params)
		operRes[oper] = rs[0].Interface().(decimal.Decimal).String()
	}
	return &types.ResMathv2{
		Result: operRes,
	}, nil
}
```

### 启动服务和调用
- grpc
    * service启动
    ```go
        endpoint := "127.0.0.1:3333"

        grpcsrv := grpc.NewServer()

        types.RegisterMathV2Server(grpcsrv, core.MathV2{})

        l, err := net.Listen("tcp", endpoint)
        if err != nil {
            log.Fatal(err)
        }
        err = grpcsrv.Serve(l)
        if err != nil {
            log.Fatal(err)
        }
    ```
    * client调用
    ```go
        endpoint := "127.0.0.1:3333"

        conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
        if err != nil {
            log.Fatal(err)
        }
        client := types.NewMathV2Client(conn)
        resp, err := client.Operation(context.Background(), &types.ReqMathv2{
            A: 1,
            B: 2,
            Oper: []types.Operation{
                3, 1, 2, 0,
            },
        })
        log.Println(resp.GetResult(), err)
    ```
- restful
    * service 启动
    ```go
        ctx := context.Background()
        mux := runtime.NewServeMux()
        types.RegisterMathV2HandlerServer(ctx, mux, core.MathV2{})
        err := types.RegisterMathV2HandlerFromEndpoint(ctx, mux, "localhost:3002", []grpc.DialOption{grpc.WithInsecure()})
        if err != nil {
            log.Fatal(err)
        }
        srv := &http.Server{
            Handler: mux,
        }
        lis, err := net.Listen("tcp", ":3333")
        if err != nil {
            log.Fatal(err)
        }

        err = srv.Serve(lis)
        if err != nil {
            log.Fatal(err)
        }
    ```
    * client 调用
    ```go
        url := "http://127.0.0.1:3333/grpc_study/math/v2/oper/"
        url += "0,1,2,3"
        url += "?a=1&b=2222"
        resp, err := http.Get(url)
        if err != nil {
            log.Fatal(err)
        }
        res, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(res))
    ```