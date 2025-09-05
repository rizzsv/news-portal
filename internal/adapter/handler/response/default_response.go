package response

type ErrorResponseDefault struct {
	Meta
}

type Meta struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type DefaultSuccessResponse struct {
	Meta       Meta               `json:"meta"`
	Data       interface{}        `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	TotalRecords int64 `json:"total_records"`
	Page         int64 `json:"page"`
	PerPage      int64 `json:"per_page"`
	TotalPages   int64 `json:"total_pages"`
}
