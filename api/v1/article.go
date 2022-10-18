package v1

import (
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/rresult"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddArticle(c *gin.Context) {
	var data model.Article
	c.ShouldBindJSON(&data)
	 code := model.CreateArticle(&data)
	 c.JSON(http.StatusOK,rresult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
	 })
}