package v1

import (
	"fmt"
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/rresult"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCategory(c *gin.Context) {
	var data model.Category
	var msg string
	c.ShouldBindJSON(&data)
	code = model.CheckCategoryExits(data.Name)
	//分类已存在
	if code == errmsg.ERROR_CATEGORYNAME_USED {
		fmt.Println("category used")
		msg = errmsg.GetErrmsg(code)
	}
	if code == errmsg.SUCCESS {
		code,msg = model.AddCategory(&data)
	}
	c.JSON(http.StatusOK,rresult.Result{
		Code: code,
		Message: msg,
		Data: data,
	})
}

func GetCategory(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	data,code := model.GetCategory(id)
	c.JSON(http.StatusOK,rresult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: data,
	})
}

func GetCategories(c *gin.Context) {
	name := c.Query("name")
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data,totoal := model.GetCategories(name,pageSize,pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK,rresult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: data,
		Totoal: totoal,
	})
}
func EditCategory(c *gin.Context) {
	var data model.Category
	id,_ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.CheckCategoryExits(data.Name)
	if  code == errmsg.ERROR_CATEGORYNAME_USED {
		c.Abort()
	}
	if code == errmsg.SUCCESS {
		model.EditCategory(id,data.Name)
	}
	c.JSON(http.StatusOK,rresult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	})
}
func DeleteCategory(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)
	c.JSON(http.StatusOK,rresult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	})
}