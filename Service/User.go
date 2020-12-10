package Service

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

func LoginUser(phone string, passWord string) Result.Result {
	if ok, user := DbModel.SelectUserByPhone(phone); ok {
		if user.PassWord == passWord {
			return Result.Result{Code: Result.Ok}
		} else {
			return Result.Result{Code: Result.PassWordNotRight}
		}
	}
	return Result.Result{Code: Result.UserNotExit}
}
