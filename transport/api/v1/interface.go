package v1

import (
	"context"

	"vse-go-newsletter-api/pkg/id"
	svcmodel "vse-go-newsletter-api/service/model"
)

type Service interface {
	CreateNewsletter(ctx context.Context, newsletter svcmodel.Newsletter) error
	ListNewsletters(ctx context.Context) ([]svcmodel.Newsletter, error)
	GetNewsletter(ctx context.Context, newsletterId id.Newsletter) (*svcmodel.Newsletter, error)
	UpdateNewsletter(ctx context.Context, newsletterId id.Newsletter, newsletter svcmodel.Newsletter) (*svcmodel.Newsletter, error)
	DeleteNewsletter(ctx context.Context, newsletterId id.Newsletter) error
}
