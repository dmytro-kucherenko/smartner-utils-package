package dtos

type ErrorItem struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Details []string `json:"details"`
} // @name ErrorDto
