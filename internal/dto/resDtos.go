package dto

//todo 표준 응답 포맷

// Generic
type PageResult[T any] struct {
	Content []T `json:"result"`
	Total   int `json:"total"` //total content 개수
	Size    int `json:"size"`  //page당 개수
	Page    int `json:"size"`  //현재 page
}

func (p *PageResult[T]) GetTotalPage() int {
	totalPage := (p.Total + p.Size - 1) / p.Size
	return totalPage
}

func (p PageResult[T]) IsFirst() bool {
	return p.Page == 0 && p.GetTotalPage() > 0
}

func (p PageResult[T]) IsLast() bool {
	totalPages := p.GetTotalPage()
	return totalPages > 0 && p.Page == totalPages-1
}
