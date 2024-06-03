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
	UpdateNewsletter(ctx context.Context, newsletter model.Newsletter) (*model.Newsletter, error)
	DeleteNewsletter(ctx context.Context, newsletterId id.Newsletter, userId string) (string, error)
	CreateNewsletter(ctx context.Context, newsletter model.Newsletter) (*model.Newsletter, error)
	CreatePost(ctx context.Context, post model.Post, userId string) (*model.Post, error)
	ListPosts(ctx context.Context) ([]model.Post, error)
	ReadPost(ctx context.Context, postId string) (*model.Post, error)
	UpdatePost(ctx context.Context, post model.Post, userId string) (*model.Post, error)
	DeletePost(ctx context.Context, postId string, userId string) (string, error)
	PublishPost(ctx context.Context, postId string, userId string) (*model.Post, error)
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
