package main

import (
	"NFUShop/Config"
	"NFUShop/DbModel"
	"fmt"
	"testing"
)

func TestFindSubGoodsSet(t *testing.T) {
	Config.GetConf()
	fmt.Println(DbModel.SelectSubGoodsSet(nil, 10, 0))
}

func TestFindSubGoodsSetDescCreateTime(t *testing.T) {
	Config.GetConf()
	ok, cc := DbModel.SelectSubGoodsSetDescCreateTime(nil, 10, 0)
	fmt.Println(ok, cc)
}

func TestFindSubGoodsSetDescSell(t *testing.T) {
	Config.GetConf()
	ok, cc := DbModel.SelectSubGoodsSetDescSell(nil, 10, 0)
	fmt.Println(ok, cc)
}
