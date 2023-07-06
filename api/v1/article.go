package v1

import (
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/rresult"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func AddArticle(c *gin.Context) {
	var data model.Article
	c.ShouldBindJSON(&data)
	 code := model.CreateArticle(&data)
	 c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	 })
}

func GetSingleArticle(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	article,code := model.GetOneArticle(id)
	mddata := []byte(article.Content)
	article.Content = string(template.HTML(blackfriday.Run(mddata)))
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: article,
	})
}
func GetArticles(c *gin.Context) {
	keywords := c.Query("keywords")
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNum,_ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data,code,totoal := model.GetArticles(keywords,pageSize,pageNum)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Totoal: totoal,
		Data: data,
	})
}

func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Query("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditArticle(id,&data)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	})
}
func DeleteArticle(c *gin.Context) {
	id,_ := strconv.Atoi(c.Query("id"))
	code := model.DeleteArticle(id)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	})
}