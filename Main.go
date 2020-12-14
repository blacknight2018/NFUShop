package main

import (
	"NFUShop/Config"
	"NFUShop/Result"
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
	v1.Use(func(context *gin.Context) {
		if token, err := context.Cookie("token"); err == nil {
			if userId, ok := Utils.ParseJWT(token).(float64); ok {
				context.Set("user_id", userId)
				fmt.Println(userId)
			}
		}
		context.Next()
	})
	home := v1.Group("/home")
	cart := v1.Group("/cart")
	user := v1.Group("/user")
	goods := v1.Group("/goods")

	v1.POST("/login", func(context *gin.Context) {
		phone := context.PostForm("phone")
		passWord := context.PostForm("pass_word")
		loginResult := Service.CheckUserAuth(phone, passWord)
		if loginResult.Code == Result.Ok {
			userId := Service.GetUserIdByPhone(phone)
			context.SetCookie("token", Utils.GenerateJWT(userId), 120, "/", "localhost", false, true)
		}
		context.Writer.Write([]byte(loginResult.Get()))
	})

	v1.POST("/register", func(context *gin.Context) {
		phone := context.PostForm("phone")
		passWord := context.PostForm("pass_word")
		context.Writer.Write([]byte(Service.Register(phone, passWord).Get()))
	})
	goods.GET("", func(context *gin.Context) {
		subGoodsId := Utils.StrToInt(context.Query("sub_goods_id"))
		context.Writer.Write([]byte(Service.GetSubGoodsDetail(subGoodsId).Get()))
	})
	goods.GET("/query", func(context *gin.Context) {
		templateValue := context.Query("template_index")
		goodsId := Utils.ContextQueryInt(context, "goods_id")
		context.Writer.Write(([]byte)(Service.GetSubGoodsByTemplateIndex(goodsId, templateValue).Get()))
	})
	user.GET("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		context.Writer.Write([]byte(Service.GetUser(userId).Get()))
	})

	home.GET("/hot", func(context *gin.Context) {
		context.Writer.Write([]byte(Service.GetHotSubGoods().Get()))
	})
	home.GET("/newest", func(context *gin.Context) {
		context.Writer.Write([]byte(Service.GetNewestSubGoods().Get()))
	})

	cart.GET("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		context.Writer.Write([]byte(Service.GetUserCartSet(userId).Get()))
	})
	cart.DELETE("", func(context *gin.Context) {
		cartId := Utils.StrToInt(context.Query("cart_id"))
		fmt.Println(cartId)
	})

	r.Run(":" + strconv.Itoa(Config.GetBindPort()))
}
