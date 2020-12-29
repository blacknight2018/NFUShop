package DbModel

type Banner struct {
	Id         int    `json:"id" gorm:"column:id"`
	Img        string `json:"img" gorm:"column:img"`
	SubGoodsId int    `json:"sub_goods_id" gorm:"column:sub_goods_id"`
}

func (c *Banner) TableName() string {
	return "banner"
}

func (c *Banner) Update() bool {
	return UpdateDBObj(c)
}

func (c *Banner) Insert() bool {
	return InsertDBObj(c)
}

func (c *Banner) Delete() bool {
	return DeleteDBObj(c)
}