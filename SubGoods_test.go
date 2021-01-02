package main

import (
	"NFUShop/Config"
	"NFUShop/DbModel"
	"NFUShop/Utils"
	"fmt"
	"testing"
)

func TestFindSubGoodsSet(t *testing.T) {
	Config.GetConf()
	fmt.Println(DbModel.SelectSubGoodsSet(nil, Utils.Int2IntPtr(10), Utils.Int2IntPtr(0)))
}

func TestFindSubGoodsSetDescCreateTime(t *testing.T) {
	Config.GetConf()
	ok, cc := DbModel.SelectSubGoodsSetDescCreateTime(nil, Utils.Int2IntPtr(10), Utils.Int2IntPtr(0))
	fmt.Println(ok, cc)
}

func TestFindSubGoodsSetDescSell(t *testing.T) {
	Config.GetConf()
	ok, cc := DbModel.SelectSubGoodsSetDescSell(nil, Utils.Int2IntPtr(10), Utils.Int2IntPtr(0))
	fmt.Println(ok, cc)
}
