package Service

import (
	"NFUShop/Config"
	"NFUShop/DbModel"
	"NFUShop/Result"
	"encoding/json"
)

const (
	UnPay      = iota
	UnDelivery = iota
	Delivery   = iota
)

/**
 * @Description: 创建订单
 * @param userId
 * @param addressId
 * @param cartIdSet
 * @return Result.Result
 */
func CreateOrder(userId int, addressId int, cartIdSet []int) Result.Result {
	tx := Config.GetOneDB()
	trans := tx.Begin()
	defer tx.Close()
	ret := Result.Result{Code: Result.UnKnow}
	var subGoodsIdSet []int
	var amountSet []int
	var totalPrice float32
	for _, v := range cartIdSet {
		if ok, data := DbModel.SelectCartByCartId(v); ok {
			subGoodsIdSet = append(subGoodsIdSet, data.SubGoodsId)
			amountSet = append(amountSet, data.Amount)
			if !data.DeleteWithDB(trans) {
				tx.Rollback()
			}
		} else {
			return Result.Result{Code: Result.UnKnow}
		}
	}
	if len(subGoodsIdSet) != len(amountSet) || len(subGoodsIdSet) == 0 {
		return Result.Result{Code: Result.Ok}
	}

	if ok, address := DbModel.SelectAddressByAddressId(addressId); ok {
		if tx != nil {
			for idx, v := range subGoodsIdSet {
				if ok2, subGoods := DbModel.SelectSubGoodsBySubGoodsId(v); ok2 {
					subGoods.Stoke -= amountSet[idx]
					totalPrice += float32(amountSet[idx]) * subGoods.Price
					if false == subGoods.UpdateWithDB(trans) {
						trans.Rollback()
						ret.Code = Result.StokeNotEnough
						return ret
					}
				} else {
					trans.Rollback()
					return ret
				}

			}
			var subGoodsIdJson string
			if bytes, err := json.Marshal(subGoodsIdSet); err == nil {
				subGoodsIdJson = string(bytes)
			} else {
				trans.Rollback()
				return ret
			}
			var order DbModel.Order
			order.UserId = userId
			order.NickName = address.NickName
			order.Sex = address.Sex
			order.Detail = address.Detail
			order.Phone = address.Phone
			order.SubGoods = subGoodsIdJson
			order.TotalPrice = totalPrice
			order.Status = UnPay

			if order.InsertOrderWithDB(trans) {
				if nil == trans.Commit().Error {
					return Result.Result{Code: Result.Ok}
				}
			} else {
				trans.Rollback()
				return Result.Result{Code: Result.CreateFailure}
			}
		}
	}
	return ret
}

/**
 * @Description:查询订单
 * @param userId
 * @param addressId
 * @param cartIdSet
 * @return Result.Result
 */
func QueryOrder(userId int, addressId int, cartIdSet []int) Result.Result {
	ret := Result.Result{Code: Result.UnKnow}
	var totalPrice float32
	type nameCart struct {
		DbModel.Cart
		Price float32 `json:"price"`
		Img   string  `json:"img"`
		Desc  string  `json:"desc"`
		Title string  `json:"title"`
	}
	type name struct {
		Address    DbModel.Address `json:"address"`
		TotalPrice float32         `json:"total_price"`
		Cart       []nameCart      `json:"cart"`
	}
	var retData name
	for _, cartId := range cartIdSet {
		if ok, cart := DbModel.SelectCartByCartId(cartId); ok && cart != nil {
			amount := cart.Amount
			var tmpNameCart nameCart
			tmpNameCart.Cart = *cart
			if ok2, subGoods := DbModel.SelectSubGoodsBySubGoodsId(cart.SubGoodsId); ok2 {
				totalPrice += float32(amount) * subGoods.Price
				tmpNameCart.Price = subGoods.Price
				tmpNameCart.Img = subGoods.Img
				_, goods := DbModel.SelectGoodsByGoodsId(subGoods.GoodsId)
				tmpNameCart.Desc = goods.Desc
				tmpNameCart.Title = goods.Title
			} else {
				return ret
			}

			retData.Cart = append(retData.Cart, tmpNameCart)
		} else {
			return ret
		}

		if ok, data := DbModel.SelectAddressByAddressId(addressId); ok && data != nil {
			retData.Address = *data
		}
	}
	retData.TotalPrice = totalPrice
	ret.Data = retData
	ret.Code = Result.Ok

	return ret
}
