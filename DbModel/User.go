package DbModel

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	Id         int        `gorm:"column:id;primary_key"`
	PassWord   string     `gorm:"column:pass_word"`
	Phone      string     `gorm:"column:phone"`
	CreateTime *time.Time `gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `gorm:"column:update_time" sql:"-"`
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

func FindUserByUserId(userId int) (bool, *User) {
	var user User
	return FindTableById("user", userId, &user), &user
}
