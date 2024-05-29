package model

type Post struct {
	ID           string `json:"id"`
	Title        string `json:"name" validate:"required"`
	Content      string `json:"description" validate:"required"`
	NewsletterId string `json:"NewsletterId" validate:"required"`
}
