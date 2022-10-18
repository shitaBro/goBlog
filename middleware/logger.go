package middleware

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)
var (
	Log = logrus.New()
)
func Loggering() gin.HandlerFunc {
	filepath := "log/ginblog"
	linkName := "latest_log.log"
	src,err := os.OpenFile(filepath,os.O_RDWR|os.O_CREATE,0755)
	if err != nil {
		fmt.Println("open file err:",err)
	}
	
	Log.Out = src
	Log.SetLevel(logrus.DebugLevel)
	logWriter,_ := rotatelogs.New(
		filepath + "%y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),//最大保存时间，7天
		rotatelogs.WithRotationTime(24*time.Hour),// 日志切割时间间隔
		rotatelogs.WithLinkName(linkName),
	)
	//日志写出配置
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap,&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.AddHook(Hook)
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms",int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName,err := os.Hostname()//主机名
		if err != nil {
			hostName = "unknown"
		}
		statusCode := ctx.Writer.Status() //状态码
		clientIp := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		dataSize := ctx.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := ctx.Request.Method
		path := ctx.Request.RequestURI
		body,err  := ioutil.ReadAll(ctx.Request.Body)
	
		fmt.Println("read body err:",err)
		entry := Log.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize": dataSize,
			"Agent":     userAgent,
			"body": body,
			"param": ctx.Params,
			
		})
		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		}else if statusCode >= 400{
			entry.Warn()
		}else {
			entry.Info()
		}

	}
}