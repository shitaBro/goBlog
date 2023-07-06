package model

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//timestemp 转换成string
func UnitToDate(timestemp int)string {
	t := time.Unix(int64(timestemp),0)
	return t.Format("2006-01-02 15:04:05")
}
//时间转换成timestemp
func DateToUnix(str string) int64 {
	dateFormat := "2006-01-02 15:04:05"
	t,err := time.ParseInLocation(dateFormat,str,time.Local)
	if err != nil {
		fmt.Printf("err:%v\n",err)
		return 0
	}
	return t.Unix()
}
func GetUnix() int64 {
	return time.Now().Unix()
}
func GetDate() string {
	dateFormat := "2006-01-02 15:04:05"
	return time.Now().Format(dateFormat)
}
func MD5(str string)string {
	data := []byte(str)
return fmt.Sprintf("%x",md5.Sum(data))
}
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
 func GetParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}