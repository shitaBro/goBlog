package v1

import (
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)
func Login(c *gin.Context) {
	var data model.User
	// var token string
	var code int
	c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.Username,data.Password)
	if code == errmsg.SUCCESS {
		
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrmsg(code),
	})
}