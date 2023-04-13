package pagination

import "math"

// Pagination
type PagedList struct {
	Data       interface{} `json:"data"`
	Page       int64       `json:"page"`
	Limit      int64       `json:"limit"`
	TotalDocs  int64       `json:"totalDocs"`
	TotalPages int64       `json:"totalPages"`
	HasPrev    bool        `json:"hasPrev"`
	HasNext    bool        `json:"hasNext"`
}

func NewPagedList(data interface{}, page, limit, totalDocs int64) *PagedList {
	pagedList := PagedList{
		Data:      data,
		Page:      page,
		Limit:     limit,
		TotalDocs: totalDocs,
	}

	totalPaged := float64(pagedList.TotalDocs) / float64(pagedList.Limit)
	pagedList.TotalPages = int64(math.Ceil(totalPaged))
	pagedList.HasPrev = pagedList.Page > 1
	pagedList.HasNext = pagedList.Page < pagedList.TotalPages

	return &pagedList
}
