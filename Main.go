package main

import (
	"NFUShop/Config"
	"NFUShop/Service"
	"NFUShop/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	Config.GetConf()
	r := gin.Default()
	v1 := r.Group("/v1")

	home := v1.Group("/home")
	cart := v1.Group("/cart")

	v1.POST("/login", func(context *gin.Context) {
		phone := context.Query("phone")
		passWord := context.Query("pass_word")
		context.Writer.Write([]byte(Service.LoginUser(phone, passWord).Get()))
	})

	home.GET("/hot", func(context *gin.Context) {
		context.Writer.Write([]byte(Service.GetHotSubGoods().Get()))
	})
	home.GET("/newest", func(context *gin.Context) {
		context.Writer.Write([]byte(Service.GetNewestSubGoods().Get()))
	})

	cart.GET("", func(context *gin.Context) {
		userId := Utils.StrToInt(context.Query("user_id"))
		fmt.Println(userId)
	})
	cart.DELETE("", func(context *gin.Context) {
		cartId := Utils.StrToInt(context.Query("cart_id"))
		fmt.Println(cartId)
	})

	r.Run(":" + strconv.Itoa(Config.GetBindPort()))
}
