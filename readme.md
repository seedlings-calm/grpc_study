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


### 初步学习 
- 完成proto基础编写，转成golang代码
- 实现加减乘除
- 启动grpc服务
- 客户端调用


### 进步学习1
- 改变proto基础类型，设置一些参数
- 实现加减乘除
- 启动grpc，restful api
- 客户端调用


### 进步学习2
- 实现传输安全调用
...