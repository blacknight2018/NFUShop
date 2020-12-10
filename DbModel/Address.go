package DbModel

import "time"

type Address struct {
	Id         int        `gorm:"column:id;primary_key"`
	UserId     int        `gorm:"column:user_id"`
	NickName   string     `gorm:"column:nick_name"`
	Sex        string     `gorm:"column:sex;type:char(1)"`
	Phone      string     `gorm:"column:phone"`
	Detail     string     `gorm:"column:detail"`
	CreateTime *time.Time `gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `gorm:"column:update_time" sql:"-"`
}

func (a *Address) TableName() string {
	return "address"
}

func (a *Address) Update() bool {
	return UpdateDBObj(a)
}

func (a *Address) Insert() bool {
	return InsertDBObj(a)
}

func FindAddressByAddressId(addressId int) (bool, *Address) {
	var address Address
	return SelectTableRecordById("address", addressId, &address), &address
}

func FindAddressSet(condition map[string]interface{}, limit int, offset int) (bool, []Address) {
	var addressSet []Address
	return SelectTableRecordSet("address", &addressSet, condition, limit, offset), addressSet
}
