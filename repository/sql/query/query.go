package query

import (
	_ "embed"
)

var (
	//go:embed scripts/newsletter/Read.sql
	ReadNewsletter string
	//go:embed scripts/newsletter/List.sql
	ListNewsletter string
	//go:embed scripts/newsletter/Update.sql
	UpdateNewsletter string
)
