package routes

import (
	v1 "goblog/api/v1"
	"goblog/utils"

	"github.com/gin-gonic/gin"
)
func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	// 不需要token的api
	routerNoAuth := r.Group("api/v1")
	{
		routerNoAuth.POST("login",v1.Login)
	}
	_ = r.Run((utils.HttpAddress + utils.HttpPort))
}