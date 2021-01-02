package main

import (
	"NFUShop/Config"
	"NFUShop/DbModel"
	"NFUShop/Utils"
	"fmt"
	"testing"
)

func TestSelectGoodsLikeTitle(t *testing.T) {
	Config.GetConf()
	fmt.Println(DbModel.SelectGoodsSetLikeTitle("裤子", Utils.Int2IntPtr(5), Utils.Int2IntPtr(0)))
}

func TestSelectGoodsSetDescSell(t *testing.T) {
	Config.GetConf()
	//DbModel.SelectGoodsSetDescSell(Utils.Int2IntPtr(5), Utils.Int2IntPtr(0))
}
