package model

import (
	"fmt"
	"goblog/middleware"
	
	"goblog/utils/errmsg"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"github.com/gin-gonic/gin"
)

func UploadFile(file multipart.File,fileHeader *multipart.FileHeader, fileSize int64,gtx *gin.Context) (string,int) {
	middleware.Log.Info("upload file Name:",fileHeader.Filename)
	day := GetDate()
	dir := "static/upload/" + day
	if _,err := os.Stat(dir); os.IsNotExist(err) {
		error := os.MkdirAll(dir,0777)
		fmt.Printf("dir err:%v\n",error)
		
	}
	extName := path.Ext(fileHeader.Filename)
	fileUnixTime := strconv.FormatInt(GetUnix(),10)
	pathDir := path.Join(dir,fileUnixTime + extName)
	err := gtx.SaveUploadedFile(fileHeader,pathDir)
	if err != nil {
		return "",errmsg.ERROR
	}
	return pathDir,errmsg.SUCCESS
}