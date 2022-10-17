package routes

import (
	v1 "goblog/api/v1"
	"goblog/middleware"
	"goblog/utils"

	"github.com/gin-gonic/gin"
)
func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	//需要验证token
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.POST("user/add",v1.AddUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		auth.PUT("user/:id/reset",v1.ResetPsw)
	}
	// 不需要token的api
	routerNoAuth := r.Group("api/v1")
	{
		routerNoAuth.POST("login",v1.Login)
		routerNoAuth.GET("user/:id",v1.GetUserInfo)
	}
	_ = r.Run((utils.HttpAddress + utils.HttpPort))
}