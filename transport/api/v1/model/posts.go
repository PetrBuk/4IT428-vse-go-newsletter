package model

type Post struct {
	ID           string `json:"id"`
	Title        string `json:"title" validate:"required"`
	Content      string `json:"content" validate:"required"`
	NewsletterId string `json:"newsletter_id" validate:"required"`
}
