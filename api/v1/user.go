package v1

import (
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/rresult"
	"goblog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
var code int
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_  = c.ShouldBindJSON(&data)
	msg,code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK,rResult.Result{
			Code: code,
			Message: msg,
		})
		return
	}
	code = model.CheckUserExits(data.Username)
	// if code == errmsg.ERROR_USERNAME_USED {

	// }
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)

	}
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: data,
	})
}
func GetUserInfo(c *gin.Context) {
	id ,_ := strconv.Atoi(c.Param("id"))
	data,code := model.GetUserInfo(id)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: data,
	})
}
func GetUsers(c *gin.Context) {
	
	username := c.Query("name")
	page,size := HandleSize(c)
	users,count := model.GetUserList(username,page,size)
	c.JSON(http.StatusOK,rResult.Result{
		Code: errmsg.SUCCESS,
		Data: users,
		Totoal: count,
	})
}
func DeleteUser(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: id,
	})
}
func ResetPsw(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	code = model.ResetPsw(id)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		
	})
}
func ChangePsw(c *gin.Context) {
	var data model.User
	id,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code  := model.ChangePsw(id,&data)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	})

}
func HandleSize(c * gin.Context) (int,int) {
	page,_ := strconv.Atoi(c.Query("page"))
	size,_ := strconv.Atoi(c.Query("size"))
	
	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
		
	}
	if page == 0 {
		page = 1
	}
	return page,size
}
