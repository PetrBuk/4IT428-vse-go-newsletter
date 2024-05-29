package service

import (
	"context"

	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/model"
)

// CreateNewsletter saves newsletter in map under email as a key.
func (s Service) CreateNewsletter(ctx context.Context, name string, description string, ownerId string) (bool, error) {
	created, err := s.repository.CreateNewsletter(ctx, name, description, ownerId)
	if err != nil {
		return false, err
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
func (s Service) UpdateNewsletter(ctx context.Context, newsletter model.Newsletter) (*model.Newsletter, error) {
	//panic("not implemented")
	dbNewsletter, err := s.repository.UpdateNewsletter(ctx, newsletter.ID, newsletter)
	if err != nil {
		return nil, err
	}
	return dbNewsletter, nil
}

// DeleteNewsletter deletes newsletter from memory.
func (s Service) DeleteNewsletter(ctx context.Context, newsletter model.Newsletter) error {
	err := s.repository.DeleteNewsletter(ctx, newsletter.ID, newsletter)
	if err != nil {
		return err
	}
	return nil
}
