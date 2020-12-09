package DbModel

import "time"

type Cart struct {
	Id         int        `gorm:"column:id;primary_key"`
	UserId     int        `gorm:"column:user_id"`
	SubGoodsId int        `gorm:"column:sub_goods_id"`
	Amount     int        `gorm:"column:amount"`
	CreateTime *time.Time `gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `gorm:"column:update_time" sql:"-"`
}

func (c *Cart) TableName() string {
	return "cart"
}

func (c *Cart) Update() bool {
	return UpdateDBObj(c)
}

func (c *Cart) Insert() bool {
	return InsertDBObj(c)
}

func FindCartByCartId(cartId int) (bool, *Cart) {
	var cart Cart
	return FindTableById("cart", cartId, &cart), &cart
}
