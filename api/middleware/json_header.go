package middleware

import "github.com/gin-gonic/gin"

func JsonHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx.Next()
	}
}
