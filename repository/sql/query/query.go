package query

import (
	_ "embed"
)

var (

	// Newsletter commands
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

	// Post commands
	//go:embed scripts/post/Read.sql
	ReadPost string
	//go:embed scripts/post/List.sql
	ListPost string
	//go:embed scripts/post/Update.sql
	UpdatePost string
	//go:embed scripts/post/Delete.sql
	DeletePost string
	//go:embed scripts/post/Create.sql
	CreatePost string
	//go:embed scripts/post/Publish.sql
	PublishPost string
)
