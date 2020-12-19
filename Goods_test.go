package main

import (
	"NFUShop/Config"
	"NFUShop/DbModel"
	"fmt"
	"testing"
)

func TestSelectGoodsLikeTitle(t *testing.T) {
	Config.GetConf()
	fmt.Println(DbModel.SelectGoodsLikeTitle("裤子", 5, 0))
}
