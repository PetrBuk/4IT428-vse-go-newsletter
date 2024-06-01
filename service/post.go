package service

import (
	"context"
	svcmodel "vse-go-newsletter-api/service/model"
)

func (s Service) CreatePost(ctx context.Context, post svcmodel.Post, userId string) (*svcmodel.Post, error) {
	created, err := s.repository.CreatePost(ctx, post, userId)
	if err != nil {
		return nil, err
	}
	return created, err
}

func (s Service) ListPosts(ctx context.Context) ([]svcmodel.Post, error) {
	posts, err := s.repository.ListPosts(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s Service) GetPost(ctx context.Context, postId string) (*svcmodel.Post, error) {
	post, err := s.repository.ReadPost(ctx, postId)

	if err != nil {
		return nil, err
	}
	return post, nil
}
func (s Service) UpdatePost(ctx context.Context, post svcmodel.Post, userId string) (*svcmodel.Post, error) {
	updatedPost, err := s.repository.UpdatePost(ctx, post, userId)
	if err != nil {
		return nil, err
	}
	return updatedPost, err
}
func (s Service) DeletePost(ctx context.Context, postId string, userId string) (string, error) {
	deleted, err := s.repository.DeletePost(ctx, postId, userId)
	if err != nil {
		return deleted, err
	}
	return deleted, err
}

func (s Service) PublishPost(ctx context.Context, postId string, userId string) (*svcmodel.Post, error) {
	updatedPost, err := s.repository.PublishPost(ctx, postId, userId)
	if err != nil {
		return nil, err
	}
	return updatedPost, err
}
