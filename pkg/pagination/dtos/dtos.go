package dtos

type PageParams struct {
	Page *uint64 `form:"page" json:"page" validate:"omitempty,min=1"`
	Size *uint64 `form:"size" json:"size" validate:"omitempty,min=1"`
} // @name PageParamsDTO

type PageMeta struct {
	Total uint64  `json:"total"`
	Page  uint64  `json:"page"`
	Size  uint64  `json:"size"`
	Next  *uint64 `json:"next"`
	Prev  *uint64 `json:"prev"`
} // @name PageMetaDTO
