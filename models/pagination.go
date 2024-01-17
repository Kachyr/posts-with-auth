package models

type PaginatedContent[T any] struct {
	Data       []T `json:"data"`
	Page       int `json:"page"`
	TotalPages int `json:"totalPages"`
}
