package response

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, statusCode int, data interface{}, err error) {
	json := gin.H{
		"data":   data,
		"error":  nil,
		"status": statusCode,
	}
	if err != nil {
		json["error"] = err.Error()
	}

	ctx.JSON(statusCode, json)
}
