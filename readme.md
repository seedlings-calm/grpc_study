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
- 实现拦截器 ，用于打印日志，查看请求方法，统计接口访问次数等功能
- #TODO 未解决 引入gogoproto编译文件未生效配置（customtype，customname等条件）
- 启动grpc,restful


### mathv...
- 实现传输安全调用
...