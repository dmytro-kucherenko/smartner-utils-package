package dtos

type ErrorResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Details []string `json:"details"`
} // @name ErrorDto
