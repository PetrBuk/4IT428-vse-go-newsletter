package repository

import (
	"context"
	"fmt"

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
	return &model.Newsletter{
		ID:          newsletter.ID,
		Name:        newsletter.Name,
		Description: newsletter.Description,
		OwnerId:     newsletter.OwnerId,
		CreateAt:    newsletter.CreatedAt,
		UpdatedAt:   newsletter.UpdatedAt,
	}, nil
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
			ID:          newsletter.ID,
			Name:        newsletter.Name,
			Description: newsletter.Description,
			OwnerId:     newsletter.OwnerId,
			CreateAt:    newsletter.CreatedAt,
			UpdatedAt:   newsletter.UpdatedAt,
		}
	}
	return response, nil
}

func (r *NewsletterRepository) UpdateNewsletter(ctx context.Context, newsletterID id.Newsletter, name string, description string, ownerId string) (*model.Newsletter, error) {
	var dbNewsletter dbmodel.Newsletter

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&dbNewsletter,
		query.UpdateNewsletter,
		pgx.NamedArgs{"id": newsletterID,
			"name":        name,
			"description": description,
			"owner_id":    ownerId,
		},
	); err != nil {
		return nil, err
	}

	updatedNewsletter := &model.Newsletter{
		ID:          dbNewsletter.ID,
		Name:        dbNewsletter.Name,
		Description: dbNewsletter.Description,
		OwnerId:     dbNewsletter.OwnerId,
		UpdatedAt:   dbNewsletter.UpdatedAt,
	}

	return updatedNewsletter, nil
}

func (r *NewsletterRepository) DeleteNewsletter(ctx context.Context, newsletterID id.Newsletter, ownerId string) (string, error) {
	if _, err := r.pool.Exec(
		ctx,
		query.DeleteNewsletter,
		pgx.NamedArgs{"id": newsletterID,
			"owner_id": ownerId,
		},
	); err != nil {
		message := fmt.Sprintf("newsletter not deleted! ID: %s", newsletterID)
		return message, err
	}
	message := fmt.Sprintf("newsletter deleted successfully! ID: %s", newsletterID)
	return message, nil

}

func (r *NewsletterRepository) CreateNewsletter(ctx context.Context, name string, description string, ownerId string) (*model.Newsletter, error) {
	var createdNewsletter dbmodel.Newsletter

	// Execute the SQL insert query with RETURNING clause
	err := pgxscan.Get(
		ctx,
		r.pool,
		&createdNewsletter,
		query.CreateNewsletter,
		pgx.NamedArgs{"name": name,
			"description": description,
			"owner_id":    ownerId,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create newsletter: %w", err)
	}

	// Convert the db model to the application model
	newNewsletter := &model.Newsletter{
		ID:          createdNewsletter.ID,
		Name:        createdNewsletter.Name,
		Description: createdNewsletter.Description,
		OwnerId:     createdNewsletter.OwnerId,
		CreateAt:    createdNewsletter.CreatedAt,
		UpdatedAt:   createdNewsletter.UpdatedAt,
	}

	// Return the newly created newsletter object
	return newNewsletter, nil
}
