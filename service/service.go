package service

import (
	"context"

	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/model"

	"github.com/supabase-community/gotrue-go"
)

type Repository interface {
	ReadNewsletter(ctx context.Context, newsletterID id.Newsletter) (*model.Newsletter, error)
	ListNewsletter(ctx context.Context) ([]model.Newsletter, error)
	UpdateNewsletter(ctx context.Context, newsletterID id.Newsletter, name string, description string, ownerId string) (*model.Newsletter, error)
	DeleteNewsletter(ctx context.Context, newsletterID id.Newsletter, ownerId string) (string, error)
	CreateNewsletter(ctx context.Context, name string, description string, ownerId string) (*model.Newsletter, error)
	CreatePost(ctx context.Context, title string, content string, newsletterId string) (*model.Post, error)
	ListPosts(ctx context.Context) ([]model.Post, error)
	GetPost(ctx context.Context, newsletterId string) (*model.Post, error)
	UpdatePost(ctx context.Context, newsletter model.Post) (*model.Post, error)
	DeletePost(ctx context.Context, newsletter model.Post) error
}

type Service struct {
	repository Repository
	authClient gotrue.Client
}

func NewService(repository Repository, authClient gotrue.Client) (Service, error) {
	return Service{
		repository: repository,
		authClient: authClient,
	}, nil
}
