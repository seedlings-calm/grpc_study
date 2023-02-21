## tls

基于CA证书的双向TLS认证

### 源码目录

```

```

### 证书生成

#### CA
1. 生成ca私钥
```shell
echo "生成 ca 密钥"
openssl genrsa -out ca.key 2048
```
2. 生成ca证书
```shell
echo "生成 ca 公钥"
openssl req -new -x509 -key ca.key -out ca.pem -days 3650 -subj  "/C=CN/ST=steve/L=ZhengZhou/O=plug/OU=plug/CN=Root"
```

#### SAN 客户端，服务端证书

[cert.sh  脚本内容](../../cert.sh)
```shell
#!/usr/bin/env sh

Country="CN"
State="steve"
Location="ZhengZhou"
Organization="plug"
Organizational="plug"
CommonName="ifcfx.com"

echo "生成服务端 SAN 证书"
openssl genpkey -algorithm RSA -out server.key
openssl req -new -nodes -key server.key -out server.csr -days 3650 -subj "/C=$Country/O=$Organization/OU=$Organizational/CN=$CommonName" -config ./openssl.cnf -extensions v3_req
openssl x509 -req -days 3650 -in server.csr -out server.pem -CA ca.pem -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req

echo "生成客户端 SAN 证书"
openssl genpkey -algorithm RSA -out client.key
openssl req -new -nodes -key client.key -out client.csr -days 3650 -subj "/C=$Country/O=$Organization/OU=$Organizational/CN=$CommonName" -config ./openssl.cnf -extensions v3_req
openssl x509 -req -days 3650 -in client.csr -out client.pem -CA ca.pem -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req

```


### 客户端证书使用



### 服务端证书使用 