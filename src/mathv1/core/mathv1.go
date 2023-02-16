package core

import (
	"context"
	"errors"

	"github.com/worryFree56/grpc_study/src/mathv1/types"
)

type MathV1 struct{}

var _ types.MathV1Server = new(MathV1)

func (MathV1) Sub(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
	if req.GetA() < req.GetB() {
		return nil, errors.New("a 小于 b")
	}
	return &types.ResMathv1{
		Res: req.GetA() - req.GetB(),
	}, nil
}
func (MathV1) Add(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
	return &types.ResMathv1{
		Res: req.GetA() + req.GetB(),
	}, nil
}
func (MathV1) Mul(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
	return &types.ResMathv1{
		Res: req.GetA() * req.GetB(),
	}, nil
}
func (MathV1) Div(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
	return &types.ResMathv1{
		Res: req.GetA() / req.GetB(),
	}, nil
}
