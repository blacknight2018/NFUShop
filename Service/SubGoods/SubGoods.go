package SubGoods

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
	"NFUShop/Service/Goods"
	"NFUShop/Utils"
)

/**
 * @Description: 获取最热销的子商品
 * @return Result.Result
 */
func GetHotSubGoods() Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	if ok, data := DbModel.SelectSubGoodsSetDescSell(nil, Utils.Int2IntPtr(6), Utils.Int2IntPtr(0)); ok {
		r.Code = Result.Ok

		type name struct {
			Title string `json:"title"`
			DbModel.SubGoods
		}
		var retData []name
		for _, v := range data {
			if ok, goods := DbModel.SelectGoodsByGoodsId(v.GoodsId); ok {
				var tmp name
				tmp.Title = goods.Title
				tmp.SubGoods = v
				retData = append(retData, tmp)

			}
		}
		r.Data = retData
	}
	return r
}

/**
 * @Description: 获取最新的子商品
 * @return Result.Result
 */
func GetNewestSubGoods() Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	if ok, data := DbModel.SelectSubGoodsSetDescCreateTime(nil, Utils.Int2IntPtr(6), Utils.Int2IntPtr(0)); ok {
		r.Code = Result.Ok
		type name struct {
			Title string `json:"title"`
			DbModel.SubGoods
		}
		var retData []name
		for _, v := range data {
			if ok, goods := DbModel.SelectGoodsByGoodsId(v.GoodsId); ok {
				var tmp name
				tmp.Title = goods.Title
				tmp.SubGoods = v
				retData = append(retData, tmp)
			}
		}
		r.Data = retData
	}
	return r
}

/**
 * @Description: 返回子商品信息
 * @param subGoodsId
 * @return Result.Result
 */
func GetSubGoodsDetail(subGoodsId int) Result.Result {
	r := Result.Result{Code: Result.UnKnow}
	type name struct {
		DbModel.Goods
		Price    float32          `json:"price"`
		Sell     int              `json:"sell"`
		SubGoods DbModel.SubGoods `json:"sub_goods"`
	}
	var retData name
	if ok, subGoods := DbModel.SelectSubGoodsBySubGoodsId(subGoodsId); ok {
		if ok2, goods := DbModel.SelectGoodsByGoodsId(subGoods.GoodsId); ok2 && goods != nil && subGoods != nil {
			retData.Goods = *goods
			retData.SubGoods = *subGoods
			retData.Price = Goods.GetLeastPrice(goods.Id)
			retData.Sell = Goods.GetGoodsSell(goods.Id)
			r.Code = Result.Ok
			r.Data = retData
		}

	}
	return r
}

/**
 * @Description: 根据商品规格查询子商品
 * @param goodsId
 * @param templateIndex
 * @return Result.Result
 */
func GetSubGoodsByTemplateIndex(goodsId int, templateIndex string) Result.Result {
	r := Result.Result{Code: Result.GoodsNotFound}
	if ok, data := DbModel.SelectSubGoodsByTemplateIndex(goodsId, templateIndex); ok {
		r.Data = data
		r.Code = Result.Ok
	}
	return r
}
