package DbModel

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ny2/utils"
	"time"
)

type User struct {
	Id         int        `json:"id" gorm:"column:id;primary_key"`
	PassWord   string     `json:"pass_word" gorm:"column:pass_word"`
	Phone      string     `json:"phone" gorm:"column:phone"`
	Avatar     string     `json:"avatar" gorm:"column:avatar"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time" sql:"-"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Update() bool {
	return UpdateDBObj(u)
}

func (u *User) Insert() bool {
	return InsertDBObj(u)
}

func SelectUserByPhone(phone string) (bool, *User) {
	var userSet []User
	ok := SelectTableRecordSet((&User{}).TableName(), &userSet, map[string]interface{}{"phone": phone}, 1, 0, utils.EmptyString)
	if ok && len(userSet) > 0 {
		return ok, &userSet[0]
	}
	return false, nil
}

func SelectUserByUserId(userId int) (bool, *User) {
	var user User
	return SelectTableRecordById((&User{}).TableName(), userId, nil, &user), &user
}

func SelectUserSet(condition map[string]interface{}, limit int, offset int) (bool, []User) {
	var userSet []User
	return SelectTableRecordSet((&User{}).TableName(), &userSet, condition, limit, offset, utils.EmptyString), userSet
}