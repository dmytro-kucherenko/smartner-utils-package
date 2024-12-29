package dtos

type PageQueryRequest struct {
	Page *uint64 `form:"page" json:"page" binding:"omitempty,min=1"`
	Size *uint64 `form:"size" json:"size" binding:"omitempty,min=1"`
} // @name PaginationQueryDto

type PageMetaResponse struct {
	Total uint64  `json:"total"`
	Page  uint64  `json:"page"`
	Size  uint64  `json:"size"`
	Next  *uint64 `json:"next"`
	Prev  *uint64 `json:"prev"`
} // @name PaginationMetaDto
