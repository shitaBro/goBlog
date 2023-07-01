package v1

import (
	"fmt"
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/rresult"
	// "goblog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetProfile(id)
	c.JSON(http.StatusOK, rResult.Result{
		Code:    code,
		Message: errmsg.GetErrmsg(code),
		Data:    data,
	})
}
func UpdateProfile(c *gin.Context) {
	var data model.Profile
	midid, _ := c.Get("user_id")
	fmt.Println("get id:", midid)
	// msg, code := validator.Validate(&data)
	// if code == errmsg.ERROR {
	// 	c.JSON(http.StatusOK, rresult.Result{
	// 		Message: msg,
	// 		Code:    code,
	// 	})
	// 	return
	// }
	// id,_ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.UpdateProfile(midid.(int), &data)
	c.JSON(http.StatusOK, rResult.Result{
		Code:    code,
		Message: errmsg.GetErrmsg(code),
	})
}
