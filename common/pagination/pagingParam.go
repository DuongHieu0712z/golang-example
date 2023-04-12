package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PagingParam struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type PagingParamOption func(*PagingParam)

func NewPagingParam(opts ...PagingParamOption) *PagingParam {
	pp := PagingParam{
		Page:  1,
		Limit: 10,
	}

	for _, opt := range opts {
		opt(&pp)
	}

	return &pp
}

func WithPage(page int64) PagingParamOption {
	return func(pp *PagingParam) {
		if page < 1 {
			page = 1
		}
		pp.Page = page
	}
}

func WithLimit(limit int64) PagingParamOption {
	return func(pp *PagingParam) {
		if limit < 1 || limit > 100 {
			limit = 10
		}
		pp.Limit = limit
	}
}

func GetPagingParam(ctx *gin.Context) *PagingParam {
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Query("limit"), 10, 64)

	return NewPagingParam(
		WithPage(page),
		WithLimit(limit),
	)
}
