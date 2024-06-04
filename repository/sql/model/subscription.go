package model

import (
	"time"
	"vse-go-newsletter-api/pkg/id"
)

type Subscription struct {
	ID           string        `db:"id"`
	CreatedAt    time.Time     `db:"created_at"`
	UserId       string        `db:"user_id"`
	NewsletterId id.Newsletter `db:"newsletter_id"`
	IsConfirmed  bool          `db:"is_confirmed"`
}
