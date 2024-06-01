package model

import (
	"time"

	"vse-go-newsletter-api/pkg/id"
)

type Post struct {
	ID           string        `db:"id"`
	CreatedAt    time.Time     `db:"created_at"`
	Title        string        `db:"title"`
	Content      string        `db:"content"`
	NewsletterId id.Newsletter `db:"newsletter_id"`
	UpdatedAt    time.Time     `db:"updated_at"`
}
