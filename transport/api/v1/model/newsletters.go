package model

type NewsLetter struct {
	Name	string `json:"name" validate:"required"`
}
