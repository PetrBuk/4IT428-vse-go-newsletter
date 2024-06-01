package service

import (
	"context"

	"github.com/supabase-community/gotrue-go/types"
)

func (s Service) Login(ctx context.Context, email string, password string) (*types.TokenResponse, error) {
	result, err := s.authClient.SignInWithEmailPassword(email, password)

	return result, err
}

func (s Service) Register(ctx context.Context, email string, password string) (*types.SignupResponse, error) {
	signupData := types.SignupRequest{
		Email:    email,
		Password: password,
	}

	settings, err := s.authClient.Signup(signupData)

	return settings, err
}

func (s Service) ChangePassword(ctx context.Context, jwtToken string, email string, oldPassword string, newPassword string) (string, error) {
	// Check old password
	// ToDo: would be fine to have a separate VerifyPassword function for this
	_, loginErr := s.authClient.SignInWithEmailPassword(email, oldPassword)

	// ToDo: Better error handling
	if loginErr != nil {
		return "", loginErr
	}

	requestData := types.UpdateUserRequest{
		Password: &newPassword,
	}

	loggedClient := s.authClient.WithToken(jwtToken)

	_, err := loggedClient.UpdateUser(requestData)

	if err != nil {
		return "error", err
	}

	return "success", err
}

func (s Service) Verify(ctx context.Context, verificationType types.VerificationType, email string, otpToken string) (*types.VerifyForUserResponse, error) {
	requestData := types.VerifyForUserRequest{
		Type:       verificationType,
		Token:      otpToken,
		Email:      email,
		RedirectTo: "http://localhost:3000",
	}

	resp, err := s.authClient.VerifyForUser(requestData)

	return resp, err
}

func (s Service) RefreshToken(ctx context.Context, refreshToken string) (*types.TokenResponse, error) {
	resp, err := s.authClient.RefreshToken(refreshToken)

	return resp, err
}
