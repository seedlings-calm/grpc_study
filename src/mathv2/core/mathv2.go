package core

import (
	"context"
	"reflect"

	"github.com/shopspring/decimal"
	"github.com/worryFree56/grpc_study/src/mathv2/types"
)

type MathV2 struct{}

var _ types.MathV2Server = new(MathV2)

func (MathV2) Operation(ctx context.Context, req *types.ReqMathv2) (res *types.ResMathv2, err error) {
	a := req.GetA()
	b := req.GetB()
	var operRes = make(map[string]string)
	for _, v := range req.GetOper() {
		adeci := decimal.NewFromInt(a)
		bdeci := decimal.NewFromInt(b)

		oper := v.String()

		value := reflect.ValueOf(adeci)
		f := value.MethodByName(oper) //通过反射获取它对应的函数，然后通过call来调用

		params := make([]reflect.Value, 1) //参数

		params[0] = reflect.ValueOf(bdeci) //参数设置为b

		rs := f.Call(params)
		operRes[oper] = rs[0].Interface().(decimal.Decimal).String()
	}
	return &types.ResMathv2{
		Result: operRes,
	}, nil
}
