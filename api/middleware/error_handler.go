package middleware

import (
	"example/common/errs"
	"example/common/response"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorHandler(ctx *gin.Context) {
	if r := recover(); r != nil {
		var statusCode int
		var err error

		switch e := r.(type) {
		case errs.HttpError:
			statusCode = e.StatusCode
			err = e
		case error:
			statusCode = http.StatusInternalServerError
			err = e
		default:
			statusCode = http.StatusInternalServerError
			err = fmt.Errorf("%v", e)
		}

		log.Println(err)
		response.Response(ctx, statusCode, nil, err)
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer errorHandler(ctx)
		ctx.Next()
	}
}
