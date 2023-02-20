package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
)

var (
	JwtSecret       = "jwttest1234567"
	headerAuthorize = "jwt"
)

type JwtAuthToken struct {
	Token string
}

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

func (j JwtAuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		headerAuthorize: j.Token,
	}, nil
}

func (c JwtAuthToken) RequireTransportSecurity() bool {
	return false
}

// Claims defines the struct containing the token claims.
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
