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
	
	//跨域中间件
	r.Use(middleware.Cors())
	//日志
	r.Use(middleware.Loggering())
	r.Use(gin.Recovery())
	//需要验证token
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.POST("user/add",v1.AddUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		auth.PUT("user/:id/reset",v1.ResetPsw)
		auth.PUT("admin/changepsw/:id",v1.ChangePsw)
		auth.PUT("profile",v1.UpdateProfile)

		auth.POST("category/add",v1.AddCategory)
		auth.PUT("category/:id",v1.EditCategory)
		auth.DELETE("category/:id",v1.DeleteCategory)
		
		auth.POST("article/add",v1.AddArticle)
	}
	// 不需要token的api
	routerNoAuth := r.Group("api/v1")
	{
		routerNoAuth.POST("login",v1.Login)
		routerNoAuth.GET("user/:id",v1.GetUserInfo)

		routerNoAuth.GET("profile/:id",v1.GetProfile)
		routerNoAuth.GET("category/:id",v1.GetCategory)
		routerNoAuth.GET("categories",v1.GetCategories)
		routerNoAuth.GET("article/:id",v1.GetSingleArticle)
		routerNoAuth.GET("article/list",v1.GetArticles)
	}
	_ = r.Run((utils.HttpAddress + utils.HttpPort))
}