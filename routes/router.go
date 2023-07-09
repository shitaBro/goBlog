package routes

import (
	v1 "goblog/api/v1"
	"goblog/middleware"
	"goblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
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
		//用户管理
		auth.POST("user/add", v1.AddUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.PUT("user/:id/reset", v1.ResetPsw)
		auth.PUT("admin/changepsw/:id", v1.ChangePsw)
		auth.PUT("profile", v1.UpdateProfile)
		//分类管理
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//文章管理
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/edit", v1.EditArticle)
		auth.DELETE("article/delete", v1.DeleteArticle)

		//上传文件
		auth.POST("upload", v1.UpdateProfile)

		//获取个人信息
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.POST("profile", v1.UpdateProfile)
		//评论模块
		auth.GET("comment/list/:id", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UnCheckComment)

	}
	// 不需要token的api
	routerNoAuth := r.Group("api/v1")
	{
		//用户
		routerNoAuth.POST("login", v1.Login)
		routerNoAuth.POST("loginFront", v1.LoginFront)
		routerNoAuth.GET("user/:id", v1.GetUserInfo)
		routerNoAuth.GET("users", v1.GetUsers)
		routerNoAuth.GET("profile/:id", v1.GetProfile)
		//分类
		routerNoAuth.GET("category/:id", v1.GetCategory)
		routerNoAuth.GET("categories", v1.GetCategories)
		//获取category下的文章 ?cid=%s&page=%s&size=%s
		routerNoAuth.GET("category/article", v1.GetCategoryArticles)
		//文章
		routerNoAuth.GET("article/:id", v1.GetSingleArticle)
		routerNoAuth.GET("article/list", v1.GetArticles)
		routerNoAuth.GET("profile/:id", v1.GetProfile)
		//评论
		routerNoAuth.POST("addcomment", v1.AddComment)
		routerNoAuth.GET("comment/info/:id", v1.GetComment)
		routerNoAuth.GET("commentfront/:id", v1.GetCommentList)
		routerNoAuth.GET("commentcount/:id", v1.GetCommentCount)
	}
	_ = r.Run((utils.HttpAddress + utils.HttpPort))
}
