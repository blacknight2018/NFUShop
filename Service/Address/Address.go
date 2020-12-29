package Address

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

/**
 * @Description:获取收获地址信息
 * @param userId
 * @param addressId
 * @return Result.Result
 */
func GetUserAddress(userId int, addressId int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectAddressByAddressId(addressId); ok {
		if data.UserId == userId {
			ret.Code = Result.Ok
			ret.Data = data
		}
	}
	return ret
}

/**
 * @Description: 获取用户的收货地址
 * @param userId
 * @param limit
 * @param offset
 * @return Result.Result
 */
func GetUserAddressList(userId int, limit int, offset int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow

	if ok, data := DbModel.SelectUserAddressSet(userId, &limit, &offset); ok {
		ret.Code = Result.Ok
		ret.Data = data
	}
	return ret
}

/**
 * @Description: 添加一个收货地址
 * @param userId
 * @param nickName
 * @param phone
 * @param sex
 * @param detail
 * @return Result.Result
 */
func AddUserAddress(userId int, nickName string, phone string, sex string, detail string) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow

	if len(sex) == 1 {
		var a DbModel.Address
		a.UserId = userId
		a.Phone = phone
		a.Detail = detail
		a.Sex = sex
		a.NickName = nickName
		if a.Insert() {
			ret.Code = Result.Ok
		}
	}

	return ret
}

/**
 * @Description:
 * @param userId
 * @param addressId
 */
func RemoveAddress(userId int, addressId int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	if ok, data := DbModel.SelectAddressByAddressId(addressId); ok {
		if data.UserId == userId && data.Delete() {
			ret.Code = Result.Ok
		}

	}
	return ret
}
