package coredomain

type PaginatedResult[T any] struct {
	Data        T   `json:"data"`
	CountTotal  int `json:"total_count"`
	CurrentPage int `json:"current_page"`
	GroupedBy   int `json:"grouped_by"`
	TotalPages  int `json:"total_pages"`
}
