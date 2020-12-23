package User

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

/**
 * @Description:通过phone查询用户ID
 * @param phone
 * @return int
 */
func GetUserIdByPhone(phone string) int {
	var userId int
	if ok, user := DbModel.SelectUserByPhone(phone); ok {
		userId = user.Id
	}
	return userId
}

/**
 * @Description: 验证用户密码是否匹配
 * @param phone
 * @param passWord
 * @return Result.Result
 */
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

/**
 * @Description: 注册
 * @param phone
 * @param passWord
 * @return Result.Result
 */
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

/**
 * @Description: 用户ID查询用户
 * @param userId
 * @return Result.Result
 */
func GetUser(userId int) Result.Result {
	if ok, user := DbModel.SelectUserByUserId(userId); ok {
		return Result.Result{Code: Result.Ok, Data: user}
	}
	return Result.Result{Code: Result.UnKnow}
}
