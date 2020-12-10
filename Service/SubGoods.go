package Service

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

func GetHotSubGoods() Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	if ok, data := DbModel.SelectSubGoodsSetDescSell(nil, 6, 0); ok {
		r.Code = Result.Ok
		r.Data = data
	}
	return r
}

func GetNewestSubGoods() Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	if ok, data := DbModel.SelectSubGoodsSetDescCreateTime(nil, 6, 0); ok {
		r.Code = Result.Ok
		r.Data = data
	}
	return r
}
