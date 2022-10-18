package v1

import (
	"goblog/middleware"
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/rresult"
	"goblog/utils/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	
)
func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int
	c.ShouldBindJSON(&data)
	msg,code := validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK,rresult.Result{
			Code: code,
			Message: msg,
		})
		return
	}
	code,id := model.CheckLogin(data.Username,data.Password)
	if code == errmsg.SUCCESS {
		token,code = middleware.SetToken(data.Username,id) 
	}else{

	}
	c.JSON(http.StatusOK,rresult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: token,
	})
}