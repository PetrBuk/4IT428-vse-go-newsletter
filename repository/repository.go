package repository

import (
	"context"

	"vse-go-newsletter-api/pkg/id"
	dbmodel "vse-go-newsletter-api/repository/sql/model"
	"vse-go-newsletter-api/repository/sql/query"
	"vse-go-newsletter-api/service/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	*NewsletterRepository
}

func New(pool *pgxpool.Pool) (*Repository, error) {
	return &Repository{
		NewsletterRepository: NewNewsletterRepository(pool),
	}, nil
}

type NewsletterRepository struct {
	pool *pgxpool.Pool
}

func NewNewsletterRepository(pool *pgxpool.Pool) *NewsletterRepository {
	return &NewsletterRepository{
		pool: pool,
	}
}

func (r *NewsletterRepository) ReadNewsletter(ctx context.Context, newsletterID id.Newsletter) (*model.Newsletter, error) {
	var newsletter dbmodel.Newsletter
	if err := pgxscan.Get(
		ctx,
		r.pool,
		&newsletter,
		query.ReadNewsletter,
		pgx.NamedArgs{
			"id": newsletterID,
		},
	); err != nil {
		return nil, err
	}
	return &model.Newsletter{}, nil
}

func (r *NewsletterRepository) ListNewsletter(ctx context.Context) ([]model.Newsletter, error) {
	var newsletters []dbmodel.Newsletter
	if err := pgxscan.Select(
		ctx,
		r.pool,
		&newsletters,
		query.ListNewsletter,
	); err != nil {
		return nil, err
	}
	response := make([]model.Newsletter, len(newsletters))
	for i, newsletter := range newsletters {
		response[i] = model.Newsletter{
			ID: newsletter.ID,
		}
	}
	return response, nil
}

func (r *NewsletterRepository) UpdateNewsletter(ctx context.Context, newsletterID id.Newsletter, newsletter model.Newsletter) (*model.Newsletter, error) {
	var dbNewsletter dbmodel.Newsletter

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&dbNewsletter,
		query.UpdateNewsletter,
		pgx.NamedArgs{"id": newsletterID,
			"name":        newsletter.Name,
			"description": newsletter.Description,
			"owner_id":    newsletter.OwnerId,
		},
	); err != nil {
		return nil, err
	}
	return &model.Newsletter{}, nil
}
