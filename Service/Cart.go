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

func AddSubGoodsToCart(userId int, subGoodsId int) Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	var cart DbModel.Cart
	cart.UserId = userId
	cart.SubGoodsId = subGoodsId
	if ok, data := DbModel.SelectUserCartBySubGoodsId(userId, subGoodsId); ok {
		data.Amount++
		if data.Update() {
			r.Code = Result.Ok
		}
	} else {
		cart.Amount = 1
		if cart.Insert() {
			r.Code = Result.Ok
		}
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
