package service

import (
	"context"
	"vse-go-newsletter-api/pkg/id"
	svcmodel "vse-go-newsletter-api/service/model"
)

func (s Service) SubscribeNewsletter(ctx context.Context, newsletterId id.Newsletter, userId string) (*svcmodel.Subscription, error) {
	subscription, err := s.repository.SubscribeNewsletter(ctx, newsletterId, userId)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

func (s Service) UnsubscribeNewsletter(ctx context.Context, newsletterId id.Newsletter, userId string) (string, error) {
	subscription, err := s.repository.UnsubscribeNewsletter(ctx, newsletterId, userId)
	if err != nil {
		return subscription, err
	}
	return subscription, err
}

func (s Service) ConfirmSubscription(ctx context.Context, newsletterId id.Newsletter, userId string) (*svcmodel.Subscription, error) {
	subscription, err := s.repository.ConfirmSubscription(ctx, newsletterId, userId)
	if err != nil {
		return nil, err
	}
	return subscription, err
}
