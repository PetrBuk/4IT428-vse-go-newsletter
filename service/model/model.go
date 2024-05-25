package model

import (
	"time"

	"vse-go-newsletter-api/pkg/id"
)

type Newsletter struct {
	ID        id.Newsletter
	createAt  time.Time
	updatedAt time.Time
	Title     string
	Content   string
	// name     string
}
