package exchange

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"data":    data,
		"message": message,
		"status":  statusCode,
	})
}

func ResponseSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	Response(ctx, statusCode, "Success", data)
}

func ResponseError(ctx *gin.Context, statusCode int, err error) {
	Response(ctx, statusCode, err.Error(), nil)
}
