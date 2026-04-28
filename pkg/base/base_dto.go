package base

type BasePaginationRes struct {
	Total       int64 `json:"total"`
	HasNext     bool  `json:"has_next"`
	HasPrevious bool  `json:"has_previous"`
}

func (p *BasePaginationRes) SetPagination(total int64, page, limit int) {
	p.Total = total
	p.HasNext = total > int64((page)*limit)
	p.HasPrevious = page > 0
}
