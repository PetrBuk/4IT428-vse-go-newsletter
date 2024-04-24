package model

import (
	"time"

	"vse-go-newsletter-api/pkg/id"
)

type Newsletter struct {
	ID        id.Newsletter   `db:"id"`
	CreatedAt time.Time 			`db:"created_at"`
	UpdatedAt time.Time 			`db:"updated_at"`
}