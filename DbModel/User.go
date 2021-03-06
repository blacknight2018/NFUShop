package DbModel

import (
	"NFUShop/Utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id         int         `json:"id" gorm:"column:id;primary_key"`
	PassWord   string      `json:"-" gorm:"column:pass_word"`
	Phone      string      `json:"phone" gorm:"column:phone"`
	Avatar     string      `json:"avatar" gorm:"column:avatar"`
	NickName   string      `json:"nick_name" gorm:"column:nick_name"`
	Money      float32     `json:"money" gorm:"column:money"`
	CreateTime *Utils.Time `json:"create_time" gorm:"column:create_time" sql:"-"`
	UpdateTime *Utils.Time `json:"update_time" gorm:"column:update_time" sql:"-"`
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
	limit := 1
	offset := 0
	ok := SelectTableRecordSet((&User{}).TableName(), &userSet, map[string]interface{}{"phone": phone}, nil, &limit, &offset, Utils.EmptyString)
	if ok && len(userSet) > 0 {
		return ok, &userSet[0]
	}
	return false, nil
}

func SelectUserByUserId(userId int) (bool, *User) {
	var user User
	return SelectTableRecordById((&User{}).TableName(), userId, nil, &user), &user
}

func SelectUserSet(condition map[string]interface{}, limit *int, offset *int) (bool, []User) {
	var userSet []User
	return SelectTableRecordSet((&User{}).TableName(), &userSet, condition, nil, limit, offset, Utils.EmptyString), userSet
}
