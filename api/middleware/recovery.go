package middleware

import (
	"example/common/errs"
	"example/common/exchange"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

func errorHandler(ctx *gin.Context, logger *log.Logger) {
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

		goErr := errors.Wrap(err, 2)
		reset := "\033[0m"
		logger.Printf("%v\n%s%s", err, goErr.Stack(), reset)

		exchange.ResponseError(ctx, statusCode, err)
		ctx.Abort()
	}
}

func Recovery() gin.HandlerFunc {
	logger := log.New(gin.DefaultErrorWriter, "\033[31m[ERROR] ", log.LstdFlags)

	return func(ctx *gin.Context) {
		defer errorHandler(ctx, logger)
		ctx.Next()
	}
}
