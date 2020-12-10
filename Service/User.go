package Service

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

func GetUserIdByPhone(phone string) int {
	var userId int
	if ok, user := DbModel.SelectUserByPhone(phone); ok {
		userId = user.Id
	}
	return userId
}

func CheckUserAuth(phone string, passWord string) Result.Result {
	if ok, user := DbModel.SelectUserByPhone(phone); ok {
		if user.PassWord == passWord {
			return Result.Result{Code: Result.Ok}
		} else {
			return Result.Result{Code: Result.PassWordNotRight}
		}
	}
	return Result.Result{Code: Result.UserNotExit}
}

func Register(phone string, passWord string) Result.Result {
	var u DbModel.User
	if ok, _ := DbModel.SelectUserByPhone(phone); ok {
		return Result.Result{Code: Result.UserExit}
	}
	u.Phone = phone
	u.PassWord = passWord
	if u.Insert() {
		return Result.Result{Code: Result.Ok}
	}
	return Result.Result{Code: Result.UnKnow}
}
