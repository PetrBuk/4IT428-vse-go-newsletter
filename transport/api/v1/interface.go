package v1

import (
	"context"
	"vse-go-newsletter-api/pkg/id"
	svcmodel "vse-go-newsletter-api/service/model"

	types "github.com/supabase-community/gotrue-go/types"
)

type RouteService interface {
	// Auth handlers
	Login(ctx context.Context, email string, password string) (*types.TokenResponse, error)
	Register(ctx context.Context, email string, password string) (*types.SignupResponse, error)
	ChangePassword(ctx context.Context, jwtToken string, email string, oldPassword string, newPassword string) (string, error)
	Verify(ctx context.Context, verificationType types.VerificationType, email string, otpToken string) (*types.VerifyForUserResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*types.TokenResponse, error)
	// Newsletter handlers
	CreateNewsletter(ctx context.Context, name string, description string, ownerId string) (*svcmodel.Newsletter, error)
	ListNewsletters(ctx context.Context) ([]svcmodel.Newsletter, error)
	GetNewsletter(ctx context.Context, newsletterId id.Newsletter) (*svcmodel.Newsletter, error)
	UpdateNewsletter(ctx context.Context, newsletterID id.Newsletter, name string, description string, ownerId string) (*svcmodel.Newsletter, error)
	DeleteNewsletter(ctx context.Context, newsletterID id.Newsletter, ownerId string) (string, error)
	// Post handlers
	CreatePost(ctx context.Context, title string, content string, newsletterId string) (bool, error)
	ListPosts(ctx context.Context) ([]svcmodel.Post, error)
	GetPost(ctx context.Context, newsletterId string) (*svcmodel.Post, error)
	UpdatePost(ctx context.Context, newsletter svcmodel.Post) (*svcmodel.Post, error)
	DeletePost(ctx context.Context, newsletter svcmodel.Post) error
}
