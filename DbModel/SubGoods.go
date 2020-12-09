package DbModel

import "time"

type SubGoods struct {
	Id         int        `gorm:"column:id;primary_key"`
	Price      float32    `gorm:"column:price"`
	Stoke      int        `gorm:"column:stoke"`
	Sell       int        `gorm:"column:sell"`
	Img        string     `gorm:"column:img"`
	GoodsId    int        `gorm:"column:goods_id"`
	Template   string     `gorm:"column:template"`
	CreateTime *time.Time `gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `gorm:"column:update_time" sql:"-"`
}

func (s *SubGoods) TableName() string {
	return "sub_goods"
}

func (s *SubGoods) Update() bool {
	return UpdateDBObj(s)
}

func (s *SubGoods) Insert() bool {
	return InsertDBObj(s)
}

func FindSubGoodsBySubGoodsId(subGoodsId int) (bool, *SubGoods) {
	var subGoods SubGoods
	return FindTableById("sub_goods", subGoodsId, &subGoods), &subGoods
}
