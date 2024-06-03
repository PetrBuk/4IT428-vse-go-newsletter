package model

import (
	"time"

	"vse-go-newsletter-api/pkg/id"
)

type Newsletter struct {
	ID          id.Newsletter
	CreateAt    time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	OwnerId     string
}

type Post struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Title        string
	Content      string
	NewsletterId id.Newsletter
	IsPublished  bool
}

type Subscription struct {
	ID           string
	CreatedAt    time.Time
	UserId       string
	NewsletterId id.Newsletter
}
