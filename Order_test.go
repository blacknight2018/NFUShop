package main

import (
	"NFUShop/Config"
	"NFUShop/DbModel"
	"testing"
)

func TestSelectTimeOutOrderSet(t *testing.T) {
	Config.GetConf()
	DbModel.SelectCreateTimeOutOrderSet(0, 60)
}
