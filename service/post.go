package service

import (
	"context"
	svcmodel "vse-go-newsletter-api/service/model"
)

func (s Service) CreatePost(ctx context.Context, title string, content string, newsletterId string) (*svcmodel.Post, error) {
	created, err := s.repository.CreatePost(ctx, title, content, newsletterId)
	if err != nil {
		return nil, err
	}
	return created, err
}

func (s Service) ListPosts(ctx context.Context) ([]svcmodel.Post, error) {
	panic("Not Implemented")
}

func (s Service) GetPost(ctx context.Context, newsletterId string) (*svcmodel.Post, error) {
	panic("Not Implemented")
}
func (s Service) UpdatePost(ctx context.Context, newsletter svcmodel.Post) (*svcmodel.Post, error) {
	panic("Not Implemented")
}
func (s Service) DeletePost(ctx context.Context, newsletter svcmodel.Post) error {
	panic("Not Implemented")
}
