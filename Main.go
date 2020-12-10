package main

import (
	"NFUShop/Config"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {

	Config.GetConf()

	//

	r := gin.Default()
	r.Run(":" + strconv.Itoa(Config.GetBindPort()))
	//var or DbModel.Order
	//or.UserId = 1
	//or.NickName = "33"
	//or.Sex = "M"
	//or.Phone = "123213213"
	//or.Detail = "666"
	//or.Goods = "888"
	//or.TotalPrice = 3366.33
	//or.Insert()
	//fmt.Println(or)

	//var sg DbModel.SubGoods
	//sg.Template = "1"
	//sg.GoodsId = 33
	//sg.Img = "66"
	//sg.Price = 123.2
	//sg.Sell = 1
	//sg.Stoke = 77
	//fmt.Println(sg.Insert())

	//var g DbModel.Goods
	//g.Banner = "11"
	//g.Title = "33"
	//g.Template = "44"
	//g.Desc = "55"
	//g.DetailImg = "66"
	//g.Insert()
	//fmt.Println(g)

	//fmt.Println(DbModel.FindCartByCartId(1))
	//var a DbModel.Address
	//a.Phone = "1"
	//a.Detail = "dd"
	//a.Sex = "F"
	//a.NickName = "aa"
	//fmt.Println(a.Insert())
	//if ok, data := DbModel.FindUserByUid(2); ok {
	//	data.PassWord = "0123"
	//	data.Update()
	//}
	//var u DbModel.User
	//u.Phone = "22"
	//u.PassWord = "Pass"
	//u.Insert()
}
