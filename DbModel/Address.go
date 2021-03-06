package DbModel

import (
	"NFUShop/Utils"
)

type Address struct {
	Id         int         `json:"id" gorm:"column:id;primary_key"`
	UserId     int         `json:"user_id" gorm:"column:user_id"`
	NickName   string      `json:"nick_name" gorm:"column:nick_name"`
	Sex        string      `json:"sex" gorm:"column:sex;type:char(1)"`
	Phone      string      `json:"phone" gorm:"column:phone"`
	Detail     string      `json:"detail" gorm:"column:detail"`
	CreateTime *Utils.Time `json:"create_time" gorm:"column:create_time" sql:"-"`
	UpdateTime *Utils.Time `json:"update_time" gorm:"column:update_time" sql:"-"`
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

func (a *Address) Delete() bool {
	return DeleteDBObj(a)
}

func SelectUserAddressSet(userId int, limit *int, offset *int) (bool, []Address) {
	condition := map[string]interface{}{"user_id": userId}
	return SelectAddressSet(condition, limit, offset, "create_time desc")
}

func SelectAddressByAddressId(addressId int) (bool, *Address) {
	var address Address
	return SelectTableRecordById("address", addressId, nil, &address), &address
}

func SelectAddressSet(condition map[string]interface{}, limit *int, offset *int, order string) (bool, []Address) {
	var addressSet []Address
	return SelectTableRecordSet("address", &addressSet, condition, nil, limit, offset, order), addressSet
}
