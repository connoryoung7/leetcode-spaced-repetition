package models

type Pagaination[T any] struct {
	Data []T `json:"data" binding:"required"`
}
