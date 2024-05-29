package model

import (
	"time"

	"vse-go-newsletter-api/pkg/id"
)

type Newsletter struct {
	ID          id.Newsletter
	createAt    time.Time
	updatedAt   time.Time
	Name        string
	Description string
	OwnerId     string
}

type Post struct {
	ID       string
	createAt time.Time
	Title    string
	Content  string
}
