package main

import (
	"NFUShop/Config"
	"NFUShop/DbModel"
	"fmt"
	"testing"
)

func TestSelectGoodsLikeTitle(t *testing.T) {
	Config.GetConf()
	fmt.Println(DbModel.SelectGoodsSetLikeTitle("裤子", 5, 0))
}

func TestSelectGoodsSetDescSell(t *testing.T) {
	Config.GetConf()
	DbModel.SelectGoodsSetDescSell(5, 0)
}
