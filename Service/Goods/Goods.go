package Goods

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
	"NFUShop/Utils"
)

/**
 * @Description:获取商品下所有子商品出售的数量总和
 * @param goodsId
 * @return int
 */
func GetGoodsSell(goodsId int) int {
	var ret int
	if ok, data := DbModel.SelectSubGoodsSetByGoodsId(goodsId); ok {
		for _, v := range data {
			ret += v.Sell
		}
	}
	return ret
}

/**
 * @Description: 获取商品下所有子商品库存的数量总和
 * @param goodsId
 * @return int
 */
func GetGoodsStoke(goodsId int) int {
	var ret int
	if ok, data := DbModel.SelectSubGoodsSetByGoodsId(goodsId); ok {
		for _, v := range data {
			ret += *v.Stoke
		}
	}
	return ret
}

/**
 * @Description:获取改商品下的子商品最低价格
 * @param goodsId
 * @return float32
 */
func GetLeastPrice(goodsId int) float32 {
	var ret float32
	if ok, data := DbModel.SelectSubGoodsSetDescPriceByGoodsId(goodsId, Utils.Int2IntPtr(1), Utils.Int2IntPtr(0)); ok {
		if len(data) > 0 {
			ret = data[0].Price
		}
	}
	return ret
}

/**
 * @Description: 通过商品标题模糊搜索商品
 * @param title
 * @param limit
 * @param offset
 * @param descCreateTime
 * @param descPrice
 * @return Result.Result
 */
func SearchGoodsByTitle(title string, limit int, offset int, descCreateTime *bool, descPrice *bool) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	type name struct {
		DbModel.Goods
		Sell     int     `json:"sell"`
		Stoke    int     `json:"stoke"`
		Price    float32 `json:"price"`
		SubGoods []int   `json:"sub_goods"`
	}
	var retData []name
	if ok, data := DbModel.SelectGoodsSetLikeTitle(title, &limit, &offset, descCreateTime, descPrice); ok {
		var tmp name
		for _, v := range data {
			tmp.Goods = v
			tmp.Sell = GetGoodsSell(v.Id)
			tmp.Stoke = GetGoodsStoke(v.Id)
			tmp.Price = GetLeastPrice(v.Id)
			var subGoodsSet []int
			if ok2, data2 := DbModel.SelectSubGoodsSetByGoodsId(v.Id); ok2 {
				for _, v2 := range data2 {
					subGoodsSet = append(subGoodsSet, v2.Id)
				}
			}
			tmp.SubGoods = subGoodsSet
			retData = append(retData, tmp)
		}

		ret.Code = Result.Ok
		ret.Data = retData
	}
	return ret
}
