package service

import (
	"context"

	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/model"
)

// CreateNewsletter saves newsletter in map under email as a key.
func (Service) CreateNewsletter(_ context.Context, newsletter model.Newsletter) error {
	panic("not implemented")
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
func (Service) UpdateNewsletter(_ context.Context, newsletterId id.Newsletter, newsletter model.Newsletter) (*model.Newsletter, error) {
	panic("not implemented")
}

// DeleteNewsletter deletes newsletter from memory.
func (Service) DeleteNewsletter(_ context.Context, newsletterId id.Newsletter) error {
	panic("not implemented")
}
