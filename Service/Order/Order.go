package Order

import (
	"NFUShop/Config"
	"NFUShop/DbModel"
	"NFUShop/Result"
	"NFUShop/Utils"
	"encoding/json"
	"fmt"
	"time"
)

const (
	UnPay    = iota
	Pay      = iota
	Delivery = iota
	All      = iota
)

func init() {
	fmt.Println("init")
	Config.GetConf()
	go func() {
		for {
			if ok, data := DbModel.SelectCreateTimeOutOrderSet(UnPay, Config.GetOrderTime()); ok {
				for i := range data {
					fmt.Println("delete ", data[i].Id, CancelOrder(data[i].Id))
				}
			}
			time.Sleep(time.Second * 10)
		}
	}()
}
func CancelOrder(orderId int) bool {
	if ok, data := DbModel.SelectOrderByOrderId(orderId); ok {
		tx := Config.GetOneDB()
		trans := tx.Begin()
		defer tx.Close()
		subGoodsId := Utils.JSON2IntArray(data.SubGoods)
		amount := Utils.JSON2IntArray(data.Amount)
		if len(subGoodsId) != len(amount) {
			trans.Rollback()
			return false
		}
		for i := range amount {
			if ok, subGoods := DbModel.SelectSubGoodsBySubGoodsId(subGoodsId[i]); ok {
				subGoods.Stoke = Utils.Int2IntPtr(amount[i] + *subGoods.Stoke)
				subGoods.Sell = Utils.Int2IntPtr(*subGoods.Sell - amount[i])
				if false == subGoods.UpdateWithDB(trans) {
					trans.Rollback()
					return false
				}
			}
		}
		if false == data.DeleteWithDB(trans) {
			trans.Rollback()
			return false
		}

		if trans.Commit().Error == nil {
			return true
		}
	}
	return false
}

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
			var amount []int
			for idx, v := range subGoodsIdSet {
				if ok2, subGoods := DbModel.SelectSubGoodsBySubGoodsId(v); ok2 {
					tmp := *subGoods.Stoke - amountSet[idx]
					subGoods.Stoke = &tmp
					subGoods.Sell = Utils.Int2IntPtr(amountSet[idx] + *subGoods.Sell)
					totalPrice += float32(amountSet[idx]) * subGoods.Price
					amount = append(amount, amountSet[idx])
					if tmp < 0 || false == subGoods.UpdateWithDB(trans) {
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
			order.Amount = Utils.Any2JSON(amount)
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
 * @Description:查询订单需要支付的价格
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

/**
 * @Description: 查询订单
 * @param userId
 * @param status
 * @param limit
 * @param offset
 * @return Result.Result
 */
func GetOrder(userId int, status int, limit int, offset int) Result.Result {
	ret := Result.Result{Code: Result.UnKnow}
	condition := make(map[string]interface{})
	if status != All {
		condition["status"] = status
	}

	type name struct {
		DbModel.Order
		Img []string `json:"img"`
	}
	var retData []name
	if ok, data := DbModel.SelectOrderSet(condition, &limit, &offset, "id desc"); ok {
		for _, order := range data {
			var tmp name
			var img []string
			var subGoodsId []int
			json.Unmarshal([]byte(order.SubGoods), &subGoodsId)
			for _, subGoodsId := range subGoodsId {
				if ok2, subGoods := DbModel.SelectSubGoodsBySubGoodsId(subGoodsId); ok2 {
					img = append(img, subGoods.Img)
				}
			}
			tmp.Img = img
			tmp.Order = order

			retData = append(retData, tmp)
		}
		ret.Data = retData
		ret.Code = Result.Ok
	}

	return ret
}

/**
 * @Description:
 * @param orderId
 * @return Result.Result
 */
func PayOrder(orderId int) Result.Result {
	ret := Result.Result{Code: Result.UnKnow}
	if ok, data := DbModel.SelectOrderByOrderId(orderId); ok {
		data.Status = Pay
		if data.Update() {
			ret.Code = Result.Ok
		}
	}
	return ret
}
