package service

import (
	"context"

	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/errors"
	"vse-go-newsletter-api/service/model"
)

// CreateNewsletter saves newsletter in map under email as a key.
func (s Service) CreateNewsletter(ctx context.Context, newsletter model.Newsletter) (*model.Newsletter, error) {
	created, err := s.repository.CreateNewsletter(ctx, newsletter)
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

// GetNewsletter returns a newsletter with specified newsletterID.
func (s Service) GetNewsletter(ctx context.Context, newsletterID id.Newsletter) (*model.Newsletter, error) {
	newsletter, err := s.repository.ReadNewsletter(ctx, newsletterID)

	if err != nil {
		return nil, err
	}
	return newsletter, nil
}

// UpdateNewsletter updates attributes of a specified newsletter.
func (s Service) UpdateNewsletter(ctx context.Context, data model.Newsletter, userId string) (*model.Newsletter, error) {
	newsletter, err := s.repository.ReadNewsletter(ctx, data.ID)

	if err != nil {
		return nil, errors.ErrNotFound
	}

	if newsletter.OwnerId != userId {
		return nil, errors.ErrForbidden
	}

	updatedNewsletter, err := s.repository.UpdateNewsletter(ctx, data)

	if err != nil {
		return nil, err
	}

	return updatedNewsletter, nil
}

// DeleteNewsletter deletes newsletter from memory.
func (s Service) DeleteNewsletter(ctx context.Context, newsletterId id.Newsletter, userId string) (string, error) {
	newsletter, err := s.repository.ReadNewsletter(ctx, newsletterId)

	if err != nil {
		return "", errors.ErrNotFound
	}

	if newsletter.OwnerId != userId {
		return "", errors.ErrForbidden
	}

	deleted, err := s.repository.DeleteNewsletter(ctx, newsletterId, userId)

	if err != nil {
		return deleted, errors.ErrBadRequest
	}

	return deleted, nil
}
