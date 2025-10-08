package models

type Pagaination[T any] struct {
	Data       []T     `json:"data"`
	NextCursor *string `json:"nextCursor"`
	PrevCursor *string `json:"prevCursor"`
}
