package core

import (
	"context"

	"github.com/worryFree56/grpc_study/src/tls/types"
)

type Hello struct{}

var _ types.HelloServer = new(Hello)

func (h Hello) Say(ctx context.Context, req *types.MsgHello) (*types.MsgHello, error) {

	return &types.MsgHello{
		Words: "say :" + req.GetWords(),
	}, nil
}
