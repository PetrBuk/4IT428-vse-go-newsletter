package service

import (
	"context"

	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/model"
)

type Repository interface {
	ReadNewsletter(ctx context.Context, newsletterID id.Newsletter) (*model.Newsletter, error)
	ListNewsletter(ctx context.Context) ([]model.Newsletter, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) (Service, error) {
	return Service{
		repository: repository,
	}, nil
}
