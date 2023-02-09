## grpc 学习



### 目录结构 

```
|-- src
    |-- .gitignore
    |-- go.mod
    |-- go.sum
    |-- protoc.sh
    |-- readme.md
    |-- client //客户端调用
    |   |-- grpc
    |   |   |-- main.go
    |   |-- rest
    |       |-- main.go
    |-- core //逻辑实现
    |   |-- mathv1.go
    |-- proto // proto 文件
    |   |-- mathv1.proto
    |-- service // 启动服务
    |   |-- grpc
    |   |   |-- main.go
    |   |-- rest
    |       |-- main.go
    |-- third_party //第三方proto
    |   |-- gogoproto
    |   |   |-- gogo.proto
    |   |-- google
    |       |-- .DS_Store
    |       |-- api
    |       |   |-- annotations.proto
    |       |   |-- http.proto
    |       |   |-- httpbody.proto
    |       |-- protobuf
    |           |-- any.proto
    |           |-- descriptor.proto
    |-- types
        |-- mathv1.pb.go
```