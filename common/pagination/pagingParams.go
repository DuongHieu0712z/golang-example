package pagination

type PagingParams struct {
	Page  int64  `form:"page,omitempty,default=1"    binding:"omitempty,min=1"`
	Limit int64  `form:"limit,omitempty,default=10"  binding:"omitempty,min=1"`
	Field string `form:"field,omitempty,default=_id"`
	Order string `form:"order,omitempty,default=asc" binding:"omitempty,oneof=asc desc"`
}
