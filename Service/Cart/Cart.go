package Cart

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

/**
 * @Description: 获取用户的购物车列表
 * @param UserId
 * @param limit
 * @param offset
 * @return Result.Result
 */
func GetUserCartList(userId int, limit int, offset int) Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	if ok, data := DbModel.SelectCartSetByUserId(userId, limit, offset); ok {
		type name struct {
			DbModel.Cart
			Title    string           `json:"title"`
			Desc     string           `json:"desc"`
			SubGoods DbModel.SubGoods `json:"sub_goods"`
		}
		var retData []name
		for _, v := range data {
			var t name
			t.Cart = v
			if ok2, data2 := DbModel.SelectSubGoodsBySubGoodsId(v.SubGoodsId); ok2 && data2 != nil {
				t.SubGoods = *data2
				if ok3, data3 := DbModel.SelectGoodsByGoodsId(data2.GoodsId); ok3 {
					t.Title = data3.Title
					t.Desc = data3.Desc
				}
			}

			retData = append(retData, t)
		}

		r.Code = Result.Ok
		r.Data = retData
	}
	return r
}

/**
 * @Description: 添加商品进用户的购物车
 * @param userId
 * @param subGoodsId
 * @return Result.Result
 */
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

/**
 * @Description: 移除商品出用户的购物车
 * @param UserId
 * @param CartId
 * @return Result.Result
 */
func RemoveCart(userId int, cartId int) Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	if DbModel.DeleteCartByUserId(cartId, userId) {
		r.Code = Result.Ok
	}
	return r
}
