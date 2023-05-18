package exchange

import (
	"example/common/errs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Bind(ctx *gin.Context, obj interface{}) {
	if err := ctx.ShouldBind(obj); err != nil {
		panic(errs.BadRequestError(err))
	}
}

func BindJson(ctx *gin.Context, obj interface{}) {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		panic(errs.BadRequestError(err))
	}
}

func BindForm(ctx *gin.Context, obj interface{}) {
	if err := ctx.ShouldBindWith(obj, binding.Form); err != nil {
		panic(errs.BadRequestError(err))
	}
}

func BindFormMultipart(ctx *gin.Context, obj interface{}) {
	if err := ctx.ShouldBindWith(obj, binding.FormMultipart); err != nil {
		panic(errs.BadRequestError(err))
	}
}
