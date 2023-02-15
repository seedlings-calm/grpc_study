package core

import (
	"context"
	"log"
	"reflect"

	"github.com/shopspring/decimal"
	"github.com/worryFree56/grpc_study/types"
	"google.golang.org/grpc/metadata"
)

type MathV3 struct{}

var _ types.MathV3Server = new(MathV3)

func (MathV3) Operation(ctx context.Context, req *types.ReqMathv3) (res *types.ResMathv3, err error) {
	//获取client设置的header信息
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Got md: %v", md)
	}

	//处理逻辑
	a := req.GetA()
	b := req.GetB()
	var resv3 []*types.Res
	for _, v := range req.GetOper() {
		adeci, err := decimal.NewFromString(a)
		if err != nil {
			return nil, err
		}
		bdeci, err := decimal.NewFromString(b)
		if err != nil {
			return nil, err
		}
		oper := v.String()

		value := reflect.ValueOf(adeci)
		f := value.MethodByName(oper) //通过反射获取它对应的函数，然后通过call来调用

		params := make([]reflect.Value, 1) //参数

		params[0] = reflect.ValueOf(bdeci) //参数设置为b

		rs := f.Call(params)
		operRes := types.Res{
			Type:  oper,
			Value: rs[0].Interface().(decimal.Decimal).String(),
		}
		resv3 = append(resv3, &operRes)
	}

	return &types.ResMathv3{
		Code:   200,
		Result: resv3,
	}, nil
}
