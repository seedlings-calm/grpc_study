package auth

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 自定义认证 实现credentials.PerRPCCredentials接口（GetRequestMetadata，RequireTransportSecurity）
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
	if name != "admin" || pwd != "123456" {
		return status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	return nil
}
