package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Order string

const (
	ASC  Order = "asc"
	DESC Order = "desc"
)

// The paging parameters
type PagingParams struct {
	Page  int64  `json:"page"`  // The page number
	Limit int64  `json:"limit"` // The limit documents per page
	Field string `json:"field"` // The field to sort
	Order int    `json:"order"` // The order sort (ASC: ascending, DESC: descending)
}

type PagingParamsOption func(*PagingParams)

func NewPagingParams(opts ...PagingParamsOption) *PagingParams {
	pp := PagingParams{
		Page:  1,
		Limit: 10,
		Field: "_id",
		Order: 1,
	}

	for _, opt := range opts {
		opt(&pp)
	}

	return &pp
}

func WithPage(page int64) PagingParamsOption {
	return func(pp *PagingParams) {
		if page < 1 {
			page = 1
		}
		pp.Page = page
	}
}

func WithLimit(limit int64) PagingParamsOption {
	return func(pp *PagingParams) {
		if limit < 1 {
			limit = 10
		}
		pp.Limit = limit
	}
}

func WithField(field string) PagingParamsOption {
	return func(pp *PagingParams) {
		if field == "" {
			field = "_id"
		}
		pp.Field = field
	}
}

func WithOrder(order Order) PagingParamsOption {
	return func(pp *PagingParams) {
		if order == DESC {
			pp.Order = -1
		} else {
			pp.Order = 1
		}
	}
}

func GetPagingParams(ctx *gin.Context) PagingParams {
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	field := ctx.Query("field")
	order := ctx.Query("order")

	return *NewPagingParams(
		WithPage(page),
		WithLimit(limit),
		WithField(field),
		WithOrder(Order(order)),
	)
}
