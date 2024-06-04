package service

import (
	"context"
	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/mail"
	svcmodel "vse-go-newsletter-api/service/model"
)

var subscriptionConfirmed string = "Thank you for subscription!"

func (s Service) SubscribeNewsletter(ctx context.Context, newsletterId id.Newsletter, email string) (*svcmodel.Subscription, error) {
	subscription, err := s.repository.SubscribeNewsletter(ctx, newsletterId, email)
	if err != nil {
		return nil, err
	}

	return subscription, err
}

func (s Service) UnsubscribeNewsletter(ctx context.Context, newsletterId id.Newsletter, email string) (string, error) {
	subscription, err := s.repository.UnsubscribeNewsletter(ctx, newsletterId, email)
	if err != nil {
		return subscription, err
	}
	return subscription, err
}

func (s Service) ConfirmSubscription(ctx context.Context, newsletterId id.Newsletter, email string) (*svcmodel.Subscription, error) {
	subscription, err := s.repository.ConfirmSubscription(ctx, newsletterId, email)
	if err != nil {
		return nil, err
	}

	//TODO
	var subscribers []string
	if subscription != nil {
		subscribers = append(subscribers, subscription.Email)
	}

	errEmail := mail.SendMail(subscribers, subscriptionConfirmed)
	if errEmail != nil {
		return nil, errEmail
	}
	return subscription, err
}
