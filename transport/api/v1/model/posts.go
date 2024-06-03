package model

import "vse-go-newsletter-api/pkg/id"

type Post struct {
	ID           string        `json:"id"`
	Title        string        `json:"title" validate:"required"`
	Content      string        `json:"content" validate:"required"`
	NewsletterId id.Newsletter `json:"newsletterId"`
}
