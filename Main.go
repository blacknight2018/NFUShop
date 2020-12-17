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
			}
		}
		context.Next()
	})
	home := v1.Group("/home")
	cart := v1.Group("/cart")
	user := v1.Group("/user")
	goods := v1.Group("/goods")
	address := v1.Group("/address")

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
	address.DELETE("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		addressId := Utils.StrToInt(context.Query("address_id"))
		context.Writer.Write([]byte(Service.RemoveAddress(userId, addressId).Get()))
	})
	address.GET("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		limit := Utils.ContextQueryInt(context, "limit")
		offset := Utils.ContextQueryInt(context, "offset")
		addressId := Utils.ContextQueryInt(context, "address_id")
		if addressId != 0 {
			context.Writer.Write([]byte(Service.GetUserAddress(userId, addressId).Get()))
			return
		}
		context.Writer.Write([]byte(Service.GetUserAddressSet(userId, limit, offset).Get()))
	})
	address.POST("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		phone := context.PostForm("phone")
		name := context.PostForm("name")
		sex := context.PostForm("sex")
		detail := context.PostForm("detail")
		context.Writer.Write([]byte(Service.AddUserAddress(userId, name, phone, sex, detail).Get()))
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
		limit := Utils.ContextQueryInt(context, "limit")
		offset := Utils.ContextQueryInt(context, "offset")
		fmt.Println(limit, offset)
		context.Writer.Write([]byte(Service.GetUserCartSet(userId, limit, offset).Get()))
	})
	cart.DELETE("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		cartId := Utils.StrToInt(context.Query("cart_id"))
		context.Writer.Write([]byte(Service.RemoveCart(userId, cartId).Get()))

	})
	cart.POST("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		subGoodsId := Utils.StrToInt(context.PostForm("sub_goods_id"))
		fmt.Println(userId, subGoodsId)
		context.Writer.Write([]byte(Service.AddSubGoodsToCart(userId, subGoodsId).Get()))
	})

	r.Run(":" + strconv.Itoa(Config.GetBindPort()))
}
