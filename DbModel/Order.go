package DbModel

import (
	"NFUShop/Config"
	"NFUShop/Utils"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type Order struct {
	Id              int         `json:"id" gorm:"column:id;primary_key"`
	UserId          int         `json:"user_id" gorm:"column:user_id"`
	NickName        string      `json:"nick_name" gorm:"column:nick_name"`
	Sex             string      `json:"sex" gorm:"column:sex;type:char(1)"`
	Phone           string      `json:"phone" gorm:"column:phone"`
	Detail          string      `json:"detail" gorm:"column:detail"`
	SubGoods        string      `json:"sub_goods" gorm:"column:sub_goods"`
	Amount          string      `json:"amount" gorm:"column:amount"`
	Status          int         `json:"status" gorm:"column:status"`
	TotalPrice      float32     `json:"total_price" gorm:"column:total_price"`
	CreateTime      *Utils.Time `json:"create_time" gorm:"column:create_time" sql:"-"`
	UpdateTime      *Utils.Time `json:"update_time" gorm:"column:update_time" sql:"-"`
	DeliveryCode    string      `json:"delivery_code" gorm:"column:delivery_code"`
	DeliveryCompany string      `json:"delivery_company" gorm:"column:delivery_company"`
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

func (o *Order) Delete() bool {
	return DeleteDBObj(o)
}

func (o *Order) DeleteWithDB(db *gorm.DB) bool {
	return db.Delete(o).Error == nil
}

func (o *Order) InsertOrderWithDB(db *gorm.DB) bool {
	return db.Create(o).Error == nil
}

func SelectOrderByOrderId(orderId int) (bool, *Order) {
	var order Order
	return SelectTableRecordById((&Order{}).TableName(), orderId, nil, &order), &order
}

func SelectOrderSet(condition map[string]interface{}, limit *int, offset *int, order string) (bool, []Order) {
	var orderSet []Order
	return SelectTableRecordSet((&Order{}).TableName(), &orderSet, condition, nil, limit, offset, order), orderSet
}

func SelectCreateTimeOutOrderSet(status int, timeOutSeconds int) (bool, []Order) {
	var orderSet []Order
	db := Config.GetOneDB()
	if db == nil {
		return false, nil
	}
	defer db.Close()

	rawSql := "select * from `order` where status = XXX and  now()-create_time > YYY"
	rawSql = strings.ReplaceAll(rawSql, "XXX", strconv.Itoa(status))
	rawSql = strings.ReplaceAll(rawSql, "YYY", strconv.Itoa(timeOutSeconds))
	if db.Raw(rawSql).Scan(&orderSet).Error == nil {
		return true, orderSet
	}
	return false, orderSet
}
