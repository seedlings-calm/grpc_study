## grpc 学习



### 目录结构 

```
|-- src
    |-- .gitignore
    |-- go.mod
    |-- go.sum
    |-- protoc.sh
    |-- readme.md
    |-- client //客户端
    |   |-- grpc
    |   |   |-- main.go
    |   |-- rest
    |       |-- main.go
    |-- core //逻辑实现
    |   |-- mathv1.go
    |-- proto // 定义消息结构和服务接口
    |   |-- mathv1.proto
    |-- service // 服务端 
    |   |-- grpc
    |   |   |-- main.go
    |   |-- rest
    |       |-- main.go
    |-- third_party //第三方proto
    |   |-- gogoproto
    |   |   |-- gogo.proto
    |   |-- google
    |       |-- api
    |       |   |-- annotations.proto
    |       |   |-- http.proto
    |       |   |-- httpbody.proto
    |       |-- protobuf
    |           |-- any.proto
    |           |-- descriptor.proto
    |-- types //编译的客户端和服务端的代码 
        |-- mathv1.pb.go
```


### mathv1
- 完成proto基础编写，转成golang代码
- 实现加减乘除
- 启动grpc服务
- 客户端调用


### mathv2
- 使用proto的修饰符，实现多个类型
- 引入`google.api.http`,参阅`third_pary/google/api/annotations.proto`
- 调用高精度包 `github.com/shopspring/decimal` ，处理加减乘除
- 启动grpc,restful
- 客户端调用



### mathv3
- grpc安全认证ssl/tls 实现

```js
//CA证书制作
//openssl 生成证书 .key私钥文件
openssl genrsa -out ca.key 2048
//生成.csr 证书签名请求文件
openssl req -new -key ca.key -out ca.csr -subj "/C=CN/L=ZHENGZHOU/O=PLUGCHAINN/OU=PC/CN=127.0.0.1"
// 自签名升成.crt证书文件
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt  -subj "/C=CN/L=ZHENGZHOU/O=PLUGCHAINN/OU=PC/CN=127.0.0.1"

//服务端证书制作
//生成.key  私钥文件
openssl genrsa -out server.key 2048

//生成.csr 证书签名请求文件
openssl req -new -key server.key -out server.csr \
	-subj "/C=CN/L=ZHENGZHOU/O=PLUGCHAINN/OU=PC/CN=127.0.0.1" \
	-reqexts SAN \
	-config <(cat /usr/local/etc/openssl@3/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:127.0.0.1"))

//签名生成.crt 证书文件
openssl x509 -req -days 3650 \
   -in server.csr -out server.crt \
   -CA ca.crt -CAkey ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat /usr/local/etc/openssl@3/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:127.0.0.1"))

   
```

### mathv...
- 实现传输安全调用
...