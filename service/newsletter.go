package service

import (
	"context"

	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/model"
)

// CreateNewsletter saves newsletter in map under email as a key.
func (s Service) CreateNewsletter(ctx context.Context, name string, description string, ownerId string) (*model.Newsletter, error) {
	created, err := s.repository.CreateNewsletter(ctx, name, description, ownerId)
	if err != nil {
		return nil, err
	}
	return created, err
}

// ListNewsletters returns list of newsletters in array of newsletters.
func (s Service) ListNewsletters(ctx context.Context) ([]model.Newsletter, error) {
	newsletters, err := s.repository.ListNewsletter(ctx)
	if err != nil {
		return nil, err
	}
	return newsletters, nil
}

// GetNewsletter returns an newsletter with specified newsletterID.
func (s Service) GetNewsletter(ctx context.Context, newsletterID id.Newsletter) (*model.Newsletter, error) {
	newsletter, err := s.repository.ReadNewsletter(ctx, newsletterID)

	if err != nil {
		return nil, err
	}
	return newsletter, nil
}

// UpdateNewsletter updates attributes of a specified newsletter.
func (s Service) UpdateNewsletter(ctx context.Context, newsletterID id.Newsletter, name string, description string, ownerId string) (*model.Newsletter, error) {
	updatedNewsletter, err := s.repository.UpdateNewsletter(ctx, newsletterID, name, description, ownerId)
	if err != nil {
		return nil, err
	}
	return updatedNewsletter, err
}

// DeleteNewsletter deletes newsletter from memory.
func (s Service) DeleteNewsletter(ctx context.Context, newsletterID id.Newsletter, ownerId string) (string, error) {
	deleted, err := s.repository.DeleteNewsletter(ctx, newsletterID, ownerId)
	if err != nil {
		return deleted, err
	}
	return deleted, err
}
