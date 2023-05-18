package middleware

import (
	"example/common/errs"
	"example/common/exchange"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorHandler(ctx *gin.Context) {
	if r := recover(); r != nil {
		statusCode := http.StatusInternalServerError
		var err error

		switch e := r.(type) {
		case errs.HttpError:
			statusCode = e.StatusCode
			err = e
		case error:
			err = e
		default:
			err = fmt.Errorf("%v", e)
		}

		log.Println(err)
		exchange.ResponseError(ctx, statusCode, err)
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer errorHandler(ctx)
		ctx.Next()
	}
}
