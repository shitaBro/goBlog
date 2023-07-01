package middleware

import (
	"goblog/utils"
	"goblog/utils/errmsg"
    "goblog/utils/rresult"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)
var code int
type MyClamis struct {
	Username string `json:"username"`
	Id int `json:"id"`
	jwt.StandardClaims
}
func SetToken(username string,id int) (string,int) {
	expireTime := time.Now().Add(24*time.Hour)
	setClaims := MyClamis{
		Username: username,
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "jjw",
		},
	}
	requireClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,setClaims)
	token,err := requireClaim.SignedString(JwtKey)
	if  err != nil {
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCESS
}
func CheckToken(token string) (*MyClamis,int) {
	setToken,_ := jwt.ParseWithClaims(token,&MyClamis{},func(t *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
	if key,_:= setToken.Claims.(*MyClamis); setToken.Valid {
		return key,errmsg.SUCCESS
		
	}else {
		return nil,errmsg.ERROR
	}
}

func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			ctx.JSON(http.StatusOK,rResult.Result{
				Code: code,
				Message: errmsg.GetErrmsg(code),
			})
			ctx.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader," ",2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK,rResult.Result{
				Code: code,
				Message: errmsg.GetErrmsg(code),
			})
			ctx.Abort()
			return
		}
		key,tcode := CheckToken(checkToken[1])
		ctx.Set("user_id",key.Id)
		if tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			ctx.JSON(http.StatusOK,rResult.Result{
				Code: code,
				Message: errmsg.GetErrmsg(code),
			})
			ctx.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME 
			ctx.JSON(http.StatusOK,rResult.Result{
				Code: code,
				Message: errmsg.GetErrmsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Set("username",key.Username)
		ctx.Next()
	}
}