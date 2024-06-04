package service

import (
	"context"
	"errors"
	"fmt"
	"vse-go-newsletter-api/service/mail"
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

	post, err := s.repository.ReadPost(ctx, postId)

	if post.IsPublished {
		return post, errors.New(fmt.Sprintf("Post with id %s has already been published!", post.ID))
	}
	if err == nil {
		subscribers, err := s.repository.GetSubscribers(ctx, post.NewsletterId)
		if len(subscribers) == 0 {
			return nil, errors.New(fmt.Sprintf("There are no subscribers for newsletter with id: %s", post.NewsletterId.String()))
		}
		if err == nil {
			err := mail.SendMail(subscribers, post.Content)
			if err == nil {
				updatedPost, err := s.repository.PublishPost(ctx, postId, userId)
				if err != nil {
					return nil, err
				}
				return updatedPost, err
			}
		}
	}
	return nil, err
}
