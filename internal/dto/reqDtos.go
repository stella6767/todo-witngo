package dto

//Generic

type PageResult[T any] struct {
	Content []T `json:"result"`
	Total   int `json:"total"`
	Size    int `json:"size"`
}

type Pageable struct {
	Page int `form:"page"` // ?page=0
	Size int `form:"size"` // ?size=10
}

func (p *PageResult[T]) GetTotalPage() int {
	totalPage := (p.Total + p.Size - 1) / p.Size
	return totalPage
}
