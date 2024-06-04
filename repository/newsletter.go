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

type NewsletterRepository struct {
	pool *pgxpool.Pool
}

func NewNewsletterRepository(pool *pgxpool.Pool) *NewsletterRepository {
	return &NewsletterRepository{
		pool: pool,
	}
}

func (r *NewsletterRepository) CreateNewsletter(ctx context.Context, newsletter model.Newsletter) (*model.Newsletter, error) {
	var createdNewsletter dbmodel.Newsletter

	err := pgxscan.Get(
		ctx,
		r.pool,
		&createdNewsletter,
		query.CreateNewsletter,
		pgx.NamedArgs{
			"name":        newsletter.Name,
			"description": newsletter.Description,
			"owner_id":    newsletter.OwnerId,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create newsletter: %w", err)
	}

	newNewsletter := &model.Newsletter{
		ID:          createdNewsletter.ID,
		Name:        createdNewsletter.Name,
		Description: createdNewsletter.Description,
		OwnerId:     createdNewsletter.OwnerId,
		CreateAt:    createdNewsletter.CreatedAt,
		UpdatedAt:   createdNewsletter.UpdatedAt,
	}

	return newNewsletter, nil
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

func (r *NewsletterRepository) UpdateNewsletter(ctx context.Context, newsletter model.Newsletter) (*model.Newsletter, error) {
	var dbNewsletter dbmodel.Newsletter

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&dbNewsletter,
		query.UpdateNewsletter,
		pgx.NamedArgs{"id": newsletter.ID,
			"name":        newsletter.Name,
			"description": newsletter.Description,
			"owner_id":    newsletter.OwnerId,
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

func (r *NewsletterRepository) DeleteNewsletter(ctx context.Context, newsletterId id.Newsletter, userId string) (string, error) {
	result, err := r.pool.Exec(
		ctx,
		query.DeleteNewsletter,
		pgx.NamedArgs{
			"id":       newsletterId,
			"owner_id": userId,
		},
	)

	if err != nil {
		message := fmt.Sprintf("newsletter not deleted! ID: %s", newsletterId)
		return message, err
	}

	if result.RowsAffected() == 0 {
		message := fmt.Sprintf("You are not allowed to delete it. ID: %s", newsletterId)
		return message, fmt.Errorf("no rows affected")
	}
	message := fmt.Sprintf("newsletter deleted successfully! ID: %s", newsletterId)
	return message, nil
}

func (r *NewsletterRepository) SubscribeNewsletter(ctx context.Context, newsletterId id.Newsletter, email string) (*model.Subscription, error) {
	var subscription dbmodel.Subscription

	err := pgxscan.Get(
		ctx,
		r.pool,
		&subscription,
		query.SubscribeNewsletter,
		pgx.NamedArgs{
			"newsletter_id": newsletterId,
			"email":         email,
		},
	)
	if err != nil {
		return nil, err
	}

	serviceSubscription := &model.Subscription{
		ID:           subscription.ID,
		Email:        subscription.Email,
		NewsletterId: subscription.NewsletterId,
		CreatedAt:    subscription.CreatedAt,
		IsConfirmed:  subscription.IsConfirmed,
	}

	return serviceSubscription, nil
}

func (r *NewsletterRepository) UnsubscribeNewsletter(ctx context.Context, newsletterId id.Newsletter, email string) (string, error) {
	result, err := r.pool.Exec(ctx, query.UnsubscribeNewsletter, pgx.NamedArgs{
		"newsletter_id": newsletterId,
		"email":         email,
	})

	if err != nil {
		message := fmt.Sprintf("failed to cancel subscription to newsletter ID: %s", newsletterId)
		return message, err
	}

	if result.RowsAffected() == 0 {
		message := fmt.Sprintf("You are not subscribed to the newsletter ID: %s. There was no need to unsubscribe", newsletterId)
		return message, fmt.Errorf("no rows affected")
	}
	message := fmt.Sprintf("Successful newsletter unsubscription (ID: %s)", newsletterId)
	return message, nil
}

func (r *NewsletterRepository) GetSubscribers(ctx context.Context, newsletter id.Newsletter) ([]string, error) {
	var subscribers []string
	newsletterId := newsletter.String()

	err := pgxscan.Select(
		ctx,
		r.pool,
		&subscribers,
		query.GetSubscribers,
		pgx.NamedArgs{
			"newsletterId": newsletterId,
		})
	response := make([]string, len(subscribers))
	for i, subscriber := range subscribers {
		response[i] = subscriber
	}

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *NewsletterRepository) ConfirmSubscription(ctx context.Context, newsletterId id.Newsletter, email string) (*model.Subscription, error) {
	var subscription dbmodel.Subscription

	err := pgxscan.Get(
		ctx,
		r.pool,
		&subscription,
		query.ConfirmSubscription,
		pgx.NamedArgs{
			"newsletter_id": newsletterId,
			"email":         email,
		},
	)
	if err != nil {
		return nil, err
	}

	serviceSubscription := &model.Subscription{
		ID:           subscription.ID,
		Email:        subscription.Email,
		NewsletterId: subscription.NewsletterId,
		CreatedAt:    subscription.CreatedAt,
		IsConfirmed:  subscription.IsConfirmed,
	}

	return serviceSubscription, nil
}
