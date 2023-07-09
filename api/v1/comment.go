package v1

import (
	"goblog/model"
	"goblog/utils/errmsg"
	rResult "goblog/utils/rresult"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	code := model.AddComment(&data)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Data: data,
		Message: errmsg.GetErrmsg(code),
	})
}

func GetComment(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	data,code := model.GetComment(id)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: data,
	})
}

func DeleteComment(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	code := model.DeleteComment(uint(id))
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	})

}

func GetCommentCount(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	totoal := model.GetCommentCount(id)
	c.JSON(http.StatusOK,rResult.Result{
		Code: errmsg.SUCCESS,
		Totoal: totoal,
	})
}

func GetCommentList(c *gin.Context) {
	page,size := HandleSize(c)
	id,_ := strconv.Atoi(c.Param("id"))
	data,totoal,code := model.GetCommentList(id,size,page)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Data: data,
		Totoal: totoal,
		Message: errmsg.GetErrmsg(code),
	})
}
func CheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id,_ := strconv.Atoi(c.Param("id"))
	code := model.CheckComment(id,&data)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	})
}
func UnCheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id,_ := strconv.Atoi(c.Param("id"))
	code := model.UnCheckComment(id,&data)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	})
}
