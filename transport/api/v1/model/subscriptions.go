package model

import "vse-go-newsletter-api/pkg/id"

type Subscription struct {
	ID           string        `json:"id"`
	Email        string        `json:"email" validate:"required"`
	NewsletterId id.Newsletter `json:"newsletterId" validate:"required"`
}
