package DbModel

import (
	"ny2/utils"
	"time"
)

type Order struct {
	Id         int        `gorm:"column:id;primary_key"`
	UserId     int        `gorm:"column:user_id"`
	NickName   string     `gorm:"column:nick_name"`
	Sex        string     `gorm:"column:sex;type:char(1)"`
	Phone      string     `gorm:"column:phone"`
	Detail     string     `gorm:"column:detail"`
	Goods      string     `gorm:"column:goods"`
	TotalPrice float32    `gorm:"column:total_price"`
	CreateTime *time.Time `gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `gorm:"column:update_time" sql:"-"`
}

func (o *Order) TableName() string {
	return "order"
}

func (o *Order) Update() bool {
	return UpdateDBObj(o)
}

func (o *Order) Insert() bool {
	return InsertDBObj(o)
}

func SelectOrderByOrderId(orderId int) (bool, *Order) {
	var order Order
	return SelectTableRecordById((&Order{}).TableName(), orderId, nil, &order), &order
}

func SelectOrderSet(condition map[string]interface{}, limit int, offset int) (bool, []Order) {
	var orderSet []Order
	return SelectTableRecordSet((&Order{}).TableName(), &orderSet, condition, limit, offset, utils.EmptyString), orderSet
}
