package core

import (
	"context"

	"github.com/worryFree56/grpc_study/types"
)

type MathV1 struct{}

var _ types.MathV1Server = new(MathV1)

func (MathV1) Sub(ctx context.Context, req *types.ReqMathv1) (res *types.ResMathv1, err error) {
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
