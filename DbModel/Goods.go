package DbModel

import (
	"NFUShop/Config"
	"NFUShop/Utils"
)

type Goods struct {
	Id         int         `json:"id" gorm:"column:id;primary_key"`
	Title      string      `json:"title" gorm:"column:title"`
	Banner     string      `json:"banner" gorm:"column:banner"`
	Template   string      `json:"template" gorm:"column:template"`
	CreateTime *Utils.Time `json:"-" gorm:"column:create_time" sql:"-"`
	UpdateTime *Utils.Time `json:"-" gorm:"column:update_time" sql:"-"`
	Desc       string      `json:"desc" gorm:"column:desc"`
	DetailImg  string      `json:"detail_img" gorm:"column:detail_img"`
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

func SelectGoodsSetLikeTitle(title string, limit *int, offset *int) (bool, []Goods) {
	var goodsSet []Goods

	db := Config.GetOneDB()
	if db == nil {
		return false, nil
	}
	defer db.Close()
	db = db.Where("title like ?", "%"+title+"%")
	if limit != nil {
		db = db.Limit(*limit)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	err := db.Find(&goodsSet).Error
	return err == nil, goodsSet
}

func SelectGoodsByGoodsId(goodsId int) (bool, *Goods) {
	var goods Goods
	return SelectTableRecordById((&Goods{}).TableName(), goodsId, nil, &goods), &goods
}

func SelectGoodsSet(condition map[string]interface{}, limit *int, offset *int) (bool, []Goods) {
	var goodsSet []Goods
	return SelectTableRecordSet((&Goods{}).TableName(), &goodsSet, condition, limit, offset, Utils.EmptyString), goodsSet
}
