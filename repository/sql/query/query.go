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
	//go:embed scripts/newsletter/Delete.sql
	DeleteNewsletter string
	//go:embed scripts/newsletter/Create.sql
	CreateNewsletter string
	//go:embed scripts/post/Create.sql
	CreatePost string
)
