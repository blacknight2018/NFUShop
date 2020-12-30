package Banner

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

func GetBannerList() Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectBannerSet(nil, nil, nil); ok {
		ret.Data = data
		ret.Code = Result.Ok
	}
	return ret
}

func UpdateBanner(bannerId int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow

	return ret
}
