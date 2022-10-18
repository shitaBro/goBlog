package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowMethods: []string{"GET","POST","OPTIONS","PUT","PATCH","DELETE","UPDATE"},
			AllowHeaders: []string{"*"},
			ExposeHeaders: []string{"Content-Length","Content-Type","Authorization"},
			AllowCredentials: true,
			MaxAge: 12*time.Hour,
		},
	)
}