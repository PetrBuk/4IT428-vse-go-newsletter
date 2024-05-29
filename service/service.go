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
