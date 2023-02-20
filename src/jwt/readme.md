## jwt

gRPC 可用 metadata 自定义认证信息，需要实现接口`credentials.PerRPCCredentials`。客户端使用 `grpc.WithPerRPCCredentials` 方法，服务端使用 `metadata.FromIncomingContext` 方法从 RPC 消息的上下文中获取 metadata。

### 源码目录

```

```

### proto 简单定义实现一个rpc
```proto
syntax = "proto3";

package grpc_study.src.jwt;

option go_package = "github.com/worryFree56/grpc_study/src/jwt/types";
message MsgHello {
    string words = 1;
}

service Hello {
    rpc Say (MsgHello) returns (MsgHello);
}
```

### 项目根目录（grpc_study/）编译文件
```shell
protoc \
    -I=src  \
    --go_out=plugins=grpc:. \
    ./src/jwt/proto/jwt.proto

```
```shell
mv github.com/worryFree56/grpc_study/src/jwt/types/* ./src/jwt/types/
rm -rf github.com/  
```

### 定义自验证 JwtAuthToken
自定义认证 实现`credentials.PerRPCCredentials`接口（GetRequestMetadata，RequireTransportSecurity）
`google.golang.org/grpc/credentials/oauth`包也对 oauth2 和 jwt 提供了支持，但是`RequireTransportSecurity`默认实现为true，需要安全传输。所以自定义来实现下jwt的创建和 验证
```go
var (
	JwtSecret       = "jwttest1234567"
	headerAuthorize = "jwt"
)

type JwtAuthToken struct {
	Token string
}

func (j JwtAuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		headerAuthorize: j.Token,
	}, nil
}

func (c JwtAuthToken) RequireTransportSecurity() bool {
	return false
}

//创建新的jwttoken  格式  *****.******.***** 
func CreateJwtToken(name string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "grpc-demo-server",
		"aud":      "grpc-demo-server",
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"sub":      "auth",
		"username": name,
	})
	tokenStr, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenStr
}

// 验证token的有效性 Claims defines the struct containing the token claims.
type Claims struct {
	jwt.StandardClaims
	// Username defines the identity of the user.
	Username string `json:"username"`
}

func IsValidJWToken(ctx context.Context) (bool, error) {
	fmt.Println("开始验证jwt token")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, errors.New("missing metadata")
	}
	// 从metadata中获取授权信息
	tokenStr := ""
	if val, ok := md[headerAuthorize]; ok {
		tokenStr = val[0]
	}
	if len(tokenStr) == 0 {
		return false, errors.New("get token from context error")
	}

	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			panic("ErrInvalidAlgorithm")
		}
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return false, errors.New("jwt parse error")
	}

	if !token.Valid {
		return false, errors.New("ErrInvalidToken")
	}

	fmt.Println("验证jwt token ok")
	fmt.Printf("token:%+v", clientClaims)
	return true, nil
}
```

### 通过拦截器才验证 token有效性
开启service服务时，写法 `gsrv := grpc.NewServer(grpc.UnaryInterceptor(serviceUnaryServerInterceptor))`
```go

// jwt token 拦截器
func serviceUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 前置逻辑
	log.Println("进入service 拦截器")

   
	ok, err := auth.IsValidJWToken(ctx)
	if !ok {
		return nil, errors.New("valid err:" + err.Error())
	}

	res, err := handler(ctx, req)

	return res, err
}

```


### 客户端请求
需要传入jwt验证数据
```go
    creds := auth.JwtAuthToken{
		Token: auth.CreateJwtToken("jwt"),
	}
	conn, err := grpc.Dial(":3333",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(creds),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cli := types.NewHelloClient(conn)
	res, err := cli.Say(context.Background(), &types.MsgHello{
		Words: "jwt token!",
	})
	fmt.Println(res, err)
```
