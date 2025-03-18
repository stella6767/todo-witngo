package dto

type Pageable struct {
	Page int `form:"page"` // ?page=0
	Size int `form:"size"` // ?size=10
}
