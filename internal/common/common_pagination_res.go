package common

type PaginationMetadata struct {
	Page       int `json:"current_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	PageSize   int `json:"page_size"`
}
