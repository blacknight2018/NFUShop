package DbModel

import "time"

type Goods struct {
	Id         int        `gorm:"column:id;primary_key"`
	Title      string     `gorm:"column:title"`
	Banner     string     `gorm:"column:banner"`
	Template   string     `gorm:"column:template"`
	CreateTime *time.Time `gorm:"column:create_time" sql:"-"`
	UpdateTime *time.Time `gorm:"column:update_time" sql:"-"`
	Desc       string     `gorm:"column:desc"`
	DetailImg  string     `gorm:"column:detail_img"`
}

func (g *Goods) TableName() string {
	return "goods"
}

func (g *Goods) Update() bool {
	return UpdateDBObj(g)
}

func (g *Goods) Insert() bool {
	return InsertDBObj(g)
}

func FindGoodsByGoodsId(goodsId int) (bool, *Goods) {
	var goods Goods
	return SelectTableRecordById("goods", goodsId, &goods), &goods
}

func FindGoodsSet(condition map[string]interface{}, limit int, offset int) (bool, []Goods) {
	var goodsSet []Goods
	return SelectTableRecordSet("goods", &goodsSet, condition, limit, offset), goodsSet
}
