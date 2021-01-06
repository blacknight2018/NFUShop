package User

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
	"NFUShop/Utils"
	"encoding/base64"
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
 * @Description:
 * @param phone
 * @param passWord
 * @param avatar
 * @param nickName
 * @return Result.Result
 */
func Register(phone string, passWord string, avatar string, nickName string) Result.Result {
	var u DbModel.User
	if ok, _ := DbModel.SelectUserByPhone(phone); ok {
		return Result.Result{Code: Result.UserExit}
	}
	_, err := base64.StdEncoding.DecodeString(avatar)
	if err != nil {
		return Result.Result{Code: Result.UnKnow}
	}
	u.Avatar = Utils.UploadImg(avatar)
	if len(u.Avatar) == 0 {
		return Result.Result{Code: Result.UnKnow}
	}
	u.Phone = phone
	u.PassWord = passWord
	u.NickName = nickName
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
