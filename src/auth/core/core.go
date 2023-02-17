package core

import (
	"context"

	"github.com/worryFree56/grpc_study/src/auth/types"
)

type Auth struct {
}

// 判断 Auth是否实现AuthServer接口
var _ types.AuthTokenServer = new(Auth)

func (*Auth) SayMsg(ctx context.Context, req *types.SayMessage) (*types.SayMessage, error) {
	msg := "client send msg to:"
	return &types.SayMessage{
		Export: msg + req.Export,
	}, nil
}
