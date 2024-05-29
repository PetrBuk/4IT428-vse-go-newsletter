package model

import (
	"time"

	"vse-go-newsletter-api/pkg/id"
)

type Newsletter struct {
	ID          id.Newsletter `db:"id"`
	CreatedAt   time.Time     `db:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at"`
	Name        string        `db:"name"`
	Description string        `db:"description"`
	Owner_id    string        `db:"owner_id"`
}
