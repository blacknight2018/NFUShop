package main

import (
	"NFUShop/Config"
	"NFUShop/Result"
	"NFUShop/Service/Address"
	"NFUShop/Service/Banner"
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

func login(context *gin.Context) {
	phone := context.PostForm("phone")
	passWord := context.PostForm("pass_word")
	loginResult := User.CheckUserAuth(phone, passWord)
	if loginResult.Code == Result.Ok {
		userId := User.GetUserIdByPhone(phone)
		context.SetCookie("token", Utils.GenerateJWT(userId), 120, "/", "localhost", false, true)
	}
	context.Writer.Write([]byte(loginResult.Get()))
}

func register(context *gin.Context) {
	phone := context.PostForm("phone")
	passWord := context.PostForm("pass_word")
	context.Writer.Write([]byte(User.Register(phone, passWord).Get()))
}

func deleteAddress(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	addressId := Utils.StrToInt(context.Query("address_id"))
	context.Writer.Write([]byte(Address.RemoveAddress(userId, addressId).Get()))
}
func getAddress(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	limit := Utils.ContextQueryInt(context, "limit")
	offset := Utils.ContextQueryInt(context, "offset")
	addressId := Utils.ContextQueryInt(context, "address_id")
	if addressId != 0 {
		context.Writer.Write([]byte(Address.GetUserAddress(userId, addressId).Get()))
		return
	}
	context.Writer.Write([]byte(Address.GetUserAddressList(userId, limit, offset).Get()))
}

func addAddress(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	phone := context.PostForm("phone")
	name := context.PostForm("name")
	sex := context.PostForm("sex")
	detail := context.PostForm("detail")
	context.Writer.Write([]byte(Address.AddUserAddress(userId, name, phone, sex, detail).Get()))
}

func getGoods(context *gin.Context) {
	subGoodsId := Utils.StrToInt(context.Query("sub_goods_id"))
	context.Writer.Write([]byte(SubGoods.GetSubGoodsDetail(subGoodsId).Get()))
}

func queryGoods(context *gin.Context) {
	templateValue := context.Query("template_index")
	goodsId := Utils.ContextQueryInt(context, "goods_id")
	context.Writer.Write(([]byte)(SubGoods.GetSubGoodsByTemplateIndex(goodsId, templateValue).Get()))
}

func searchGoods(context *gin.Context) {
	keyWord := context.Query("key")
	limit := Utils.ContextQueryInt(context, "limit")
	offset := Utils.ContextQueryInt(context, "offset")
	context.Writer.Write([]byte(Goods.SearchGoodsByTitle(keyWord, limit, offset).Get()))
}

func getUser(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	context.Writer.Write([]byte(User.GetUser(userId).Get()))
}

func validToken(context *gin.Context) {
	var ret Result.Result
	ret.Code = Result.Ok
	userId := Utils.ContextGetInt(context, "user_id")
	ret.Data = userId
	context.Writer.Write([]byte(ret.Get()))
}

func getHot(context *gin.Context) {
	context.Writer.Write([]byte(SubGoods.GetHotSubGoods().Get()))
}

func getNewest(context *gin.Context) {
	context.Writer.Write([]byte(SubGoods.GetNewestSubGoods().Get()))
}

func getBanner(context *gin.Context) {
	context.Writer.Write([]byte(Banner.GetBannerList().Get()))
}

func getCart(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	limit := Utils.ContextQueryInt(context, "limit")
	offset := Utils.ContextQueryInt(context, "offset")
	fmt.Println(limit, offset)
	context.Writer.Write([]byte(Cart.GetUserCartList(userId, limit, offset).Get()))
}

func deleteCart(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	cartId := Utils.StrToInt(context.Query("cart_id"))
	context.Writer.Write([]byte(Cart.RemoveCart(userId, cartId).Get()))
}

func addCart(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	subGoodsId := Utils.StrToInt(context.PostForm("sub_goods_id"))
	fmt.Println(userId, subGoodsId)
	context.Writer.Write([]byte(Cart.AddSubGoodsToCart(userId, subGoodsId).Get()))
}

func getOrder(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	status := Utils.StrToInt(context.Query("status"))
	limit := Utils.ContextQueryInt(context, "limit")
	offset := Utils.ContextQueryInt(context, "offset")
	context.Writer.Write([]byte(Order.GetOrder(userId, status, limit, offset).Get()))
	fmt.Println(userId, limit, offset, status)
}

func queryOrder(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	addressId := Utils.StrToInt(context.Query("address_id"))
	cartArray := context.Query("cart_id")
	fmt.Println(userId, addressId, cartArray)
	context.Writer.Write([]byte(Order.QueryOrder(userId, addressId, Utils.GetJSONIntArray(cartArray)).Get()))
}

func submitOrder(context *gin.Context) {
	userId := Utils.ContextGetInt(context, "user_id")
	addressId := Utils.StrToInt(context.PostForm("address_id"))
	cartArray := context.PostForm("cart_id")
	fmt.Println(userId, addressId, cartArray)
	context.Writer.Write([]byte(Order.CreateOrder(userId, addressId, Utils.GetJSONIntArray(cartArray)).Get()))
}
func main() {
	Config.GetConf()
	Config.GetRandomSlaveDB()
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.Use(func(context *gin.Context) {
		if token, err := context.Cookie("token"); err == nil {
			if userId, ok := Utils.ParseJWT(token).(float64); ok {
				context.Set("user_id", userId)
			}
		}
		context.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		context.Next()
	})
	home := v1.Group("/home")
	cart := v1.Group("/cart")
	user := v1.Group("/user")
	goods := v1.Group("/goods")
	order := v1.Group("/order")
	address := v1.Group("/address")

	v1.POST("/login", login)

	v1.POST("/register", register)

	address.DELETE("", deleteAddress)

	address.GET("", getAddress)

	address.POST("", addAddress)

	goods.GET("", getGoods)

	goods.GET("/query", queryGoods)

	goods.GET("/search", searchGoods)

	user.GET("", getUser)

	user.GET("/valid", validToken)

	home.GET("/hot", getHot)

	home.GET("/newest", getNewest)

	home.GET("/banner", getBanner)

	cart.GET("", getCart)

	cart.DELETE("", deleteCart)

	cart.POST("", addCart)

	order.GET("", getOrder)

	order.GET("/query", queryOrder)

	order.POST("/submit", submitOrder)

	r.Run(":" + strconv.Itoa(Config.GetBindPort()))
}
