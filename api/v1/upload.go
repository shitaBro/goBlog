package v1

import (
	"fmt"
	"goblog/model"
	"goblog/utils/errmsg"
	rResult "goblog/utils/rresult"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file,fileHeader,_ := c.Request.FormFile("file")
	extName := path.Ext(fileHeader.Filename)
	allowExtMap := map[string]bool{
		".jpg":true,
		".jpeg":true,
		".gif":true,
		".png":true,
		".zip":true,
		".ipa":true,
	}
	if _,ok := allowExtMap[extName]; !ok {
		fmt.Println("不支持的类型:",extName)
		return
	}
	path,code := model.UploadFile(file,fileHeader,fileHeader.Size,c)
	c.JSON(http.StatusOK,rResult.Result{
		Code: code,
		Message: errmsg.GetErrmsg(code),
		Data: path,
	})
}
