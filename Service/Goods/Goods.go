package Goods

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

func GetGoodsSell(goodsId int) int {
	var ret int
	if ok, data := DbModel.SelectSubGoodsSetByGoodsId(goodsId); ok {
		for _, v := range data {
			ret += v.Sell
		}
	}
	return ret
}

func GetGoodsStoke(goodsId int) int {
	var ret int
	if ok, data := DbModel.SelectSubGoodsSetByGoodsId(goodsId); ok {
		for _, v := range data {
			ret += *v.Stoke
		}
	}
	return ret
}

func GetLeastPrice(goodsId int) float32 {
	var ret float32
	if ok, data := DbModel.SelectSubGoodsSetDescPriceByGoodsId(goodsId, 1, 0); ok {
		ret = data[0].Price
	}
	return ret
}

/**
 * @Description: 搜索
 * @param title
 * @param limit
 * @param offset
 * @return Result.Result
 */
func SearchGoodsByTitle(title string, limit int, offset int) Result.Result {
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
	if ok, data := DbModel.SelectGoodsSetLikeTitle(title, limit, offset); ok {
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
