package Service

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

func GetUserCartSet(UserId int) Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	if ok, data := DbModel.SelectCartSetByUserId(UserId, 5, 0); ok {
		r.Code = Result.Ok
		r.Data = data
	}
	return r
}

func RemoveCart(UserId int, CartId int) Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	if DbModel.DeleteCartByUserId(CartId, UserId) {
		r.Code = Result.Ok
	}
	return r
}
