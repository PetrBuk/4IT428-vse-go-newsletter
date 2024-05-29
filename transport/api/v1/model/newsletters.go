package model

type NewsLetter struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
