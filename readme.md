## grpc 学习


- [mathv1 入门学习](./src/mathv1/readme.md)
    - proto基本编写，基础类型
    - rpc的实现
    - service，client的编写和调用
- [mathv2 入门学习](./src/mathv2/readme.md)
    - proto的更多类型学习
    - proto定义restful
    - grpc，restful的服务启动，调用
- [mathv3 进阶学习](./src/mathv3/readme.md)
    - 拦截器学习
    - metadata的使用
- [auth验证 进阶学习](./src/auth/readme.md)
    - 自定义token验证，是否可以执行逻辑
- [jwt验证 进阶学习](./src/jwt/readme.md)

- [tls验证 进阶学习](./src/tls/readme.md) 


### 测试功能 
1. git clone https://github.com/worryFree56/grpc_study.git
2. cd grpc_study
3. go mod tidy && go mod verify
4. go run src/mathv1/service/grpc/main.go  
5. go run src/mathv1/client/grpc/main.go


[关于proto-gen-go的使用](https://blog.csdn.net/Little_Ji/article/details/129039912)