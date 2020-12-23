package main

import (
	"NFUShop/Config"
	"NFUShop/Result"
	"NFUShop/Service/Address"
	"NFUShop/Service/Cart"
	"NFUShop/Service/Goods"
	"NFUShop/Service/Order"
	"NFUShop/Service/SubGoods"
	"NFUShop/Service/User"
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
	order := v1.Group("/order")
	address := v1.Group("/address")

	v1.POST("/login", func(context *gin.Context) {
		phone := context.PostForm("phone")
		passWord := context.PostForm("pass_word")
		loginResult := User.CheckUserAuth(phone, passWord)
		if loginResult.Code == Result.Ok {
			userId := User.GetUserIdByPhone(phone)
			context.SetCookie("token", Utils.GenerateJWT(userId), 120, "/", "localhost", false, true)
		}
		context.Writer.Write([]byte(loginResult.Get()))
	})

	v1.POST("/register", func(context *gin.Context) {
		phone := context.PostForm("phone")
		passWord := context.PostForm("pass_word")
		context.Writer.Write([]byte(User.Register(phone, passWord).Get()))
	})
	address.DELETE("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		addressId := Utils.StrToInt(context.Query("address_id"))
		context.Writer.Write([]byte(Address.RemoveAddress(userId, addressId).Get()))
	})
	address.GET("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		limit := Utils.ContextQueryInt(context, "limit")
		offset := Utils.ContextQueryInt(context, "offset")
		addressId := Utils.ContextQueryInt(context, "address_id")
		if addressId != 0 {
			context.Writer.Write([]byte(Address.GetUserAddress(userId, addressId).Get()))
			return
		}
		context.Writer.Write([]byte(Address.GetUserAddressSet(userId, limit, offset).Get()))
	})
	address.POST("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		phone := context.PostForm("phone")
		name := context.PostForm("name")
		sex := context.PostForm("sex")
		detail := context.PostForm("detail")
		context.Writer.Write([]byte(Address.AddUserAddress(userId, name, phone, sex, detail).Get()))
	})
	goods.GET("", func(context *gin.Context) {
		subGoodsId := Utils.StrToInt(context.Query("sub_goods_id"))
		context.Writer.Write([]byte(SubGoods.GetSubGoodsDetail(subGoodsId).Get()))
	})
	goods.GET("/query", func(context *gin.Context) {
		templateValue := context.Query("template_index")
		goodsId := Utils.ContextQueryInt(context, "goods_id")
		context.Writer.Write(([]byte)(SubGoods.GetSubGoodsByTemplateIndex(goodsId, templateValue).Get()))
	})

	goods.GET("/search", func(context *gin.Context) {
		keyWord := context.Query("key")
		limit := Utils.ContextQueryInt(context, "limit")
		offset := Utils.ContextQueryInt(context, "offset")
		context.Writer.Write([]byte(Goods.SearchGoodsByTitle(keyWord, limit, offset).Get()))
	})
	user.GET("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		context.Writer.Write([]byte(User.GetUser(userId).Get()))
	})

	home.GET("/hot", func(context *gin.Context) {
		context.Writer.Write([]byte(SubGoods.GetHotSubGoods().Get()))
	})
	home.GET("/newest", func(context *gin.Context) {
		context.Writer.Write([]byte(SubGoods.GetNewestSubGoods().Get()))
	})

	cart.GET("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		limit := Utils.ContextQueryInt(context, "limit")
		offset := Utils.ContextQueryInt(context, "offset")
		fmt.Println(limit, offset)
		context.Writer.Write([]byte(Cart.GetUserCartSet(userId, limit, offset).Get()))
	})
	cart.DELETE("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		cartId := Utils.StrToInt(context.Query("cart_id"))
		context.Writer.Write([]byte(Cart.RemoveCart(userId, cartId).Get()))

	})
	cart.POST("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		subGoodsId := Utils.StrToInt(context.PostForm("sub_goods_id"))
		fmt.Println(userId, subGoodsId)
		context.Writer.Write([]byte(Cart.AddSubGoodsToCart(userId, subGoodsId).Get()))
	})
	order.GET("", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		status := Utils.StrToInt(context.Query("status"))
		limit := Utils.ContextQueryInt(context, "limit")
		offset := Utils.ContextQueryInt(context, "offset")
		context.Writer.Write([]byte(Order.GetOrder(userId, status, limit, offset).Get()))
		fmt.Println(userId, limit, offset, status)
	})
	order.GET("/query", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		addressId := Utils.StrToInt(context.Query("address_id"))
		cartArray := context.Query("cart_id")
		fmt.Println(userId, addressId, cartArray)
		context.Writer.Write([]byte(Order.QueryOrder(userId, addressId, Utils.GetJSONIntArray(cartArray)).Get()))
	})
	order.POST("/submit", func(context *gin.Context) {
		userId := Utils.ContextGetInt(context, "user_id")
		addressId := Utils.StrToInt(context.PostForm("address_id"))
		cartArray := context.PostForm("cart_id")
		fmt.Println(userId, addressId, cartArray)
		context.Writer.Write([]byte(Order.CreateOrder(userId, addressId, Utils.GetJSONIntArray(cartArray)).Get()))
	})

	r.Run(":" + strconv.Itoa(Config.GetBindPort()))
}
