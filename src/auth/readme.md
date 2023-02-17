## auth

gRPC 可用 metadata 自定义认证信息，需要实现接口`credentials.PerRPCCredentials`。客户端使用 `grpc.WithPerRPCCredentials` 方法，服务端使用 `metadata.FromIncomingContext` 方法从 RPC 消息的上下文中获取 metadata。

### 源码目录

```
|-- src
    |-- readme.md
    |-- auth //自验证实现
    |   |-- auth.go
    |-- client //客户端发起
    |   |-- main.go
    |-- core //rpc实现
    |   |-- core.go
    |-- proto
    |   |-- auth.proto
    |-- service //服务端
    |   |-- main.go
    |-- types
        |-- auth.pb.go

```

### proto 简单定义实现一个rpc
```proto
syntax = "proto3";


package  grpc_study.src.auth;

option go_package = "github.com/worryFree56/grpc_study/src/auth/types";

message SayMessage {
    string export = 1;
}

service AuthToken  {
    rpc SayMsg (SayMessage) returns (SayMessage);
}

```

### 项目根目录（grpc_study/）编译文件
```shell
protoc \
    -I=src  \
    --go_out=plugins=grpc:. \
    ./src/auth/proto/auth.proto


```
```shell
mv github.com/worryFree56/grpc_study/src/auth/types/* src/auth/types/
rm -rf github.com/  
```

### 定义自验证 AuthToken
自定义认证 实现`credentials.PerRPCCredentials`接口（GetRequestMetadata，RequireTransportSecurity）

```go
type AuthToken struct {
	Name, Password string
}

// GetRequestMetadata 定义授权信息的具体存放形式，最终会按这个格式存放到 metadata map 中。
func (a *AuthToken) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	log.Println("getRequestMetadata:", a)
	return map[string]string{
		"name":     a.Name,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 加密连接进行安全传输
func (a *AuthToken) RequireTransportSecurity() bool {
	return false
}

// 验证token的有效性
func IsValidAuthToken(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	var (
		name, pwd string
	)
	if val, ok := md["name"]; ok {
		name = val[0]
	}

	if val, ok := md["password"]; ok {
		pwd = val[0]
	}
    //可以接入数据库有效数据
	if name != "admin" || pwd != "123456" {
		return status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	return nil
}
```

### 通过拦截器才验证 token有效性
开启service服务时，写法 `gsrv := grpc.NewServer(grpc.UnaryInterceptor(authTokenInterceptor))`
```go

// auth token 拦截器
func authTokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 前置逻辑
	log.Println("进入service 拦截器")

    //进行验证token
	err = auth.IsValidAuthToken(ctx)
	if err != nil {
		return
	}
	// 处理请求
	response, err := handler(ctx, req)

	// 后置逻辑
	log.Println("client 请求正在处理")

	return response, err
}

```


### 客户端请求
需要传入自定义验证数据
```go
//自验证数据
	at := &auth.AuthToken{
		Name:     "admin",
		Password: "1234567",
	}
	conn, err := grpc.Dial(":3333",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		//使用自定义验证
		grpc.WithPerRPCCredentials(at),
	)
```
