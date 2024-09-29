package dto

type DataDTO[T any] struct {
	Data T `json:"data"`
}
