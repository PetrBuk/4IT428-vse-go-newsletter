package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

const (
	// AuthorizationHeader is the key for the Authorization header in the request.
	authHeader = "Authorization"
	authScheme = "Bearer "
)

type Authenticator interface {
	VerifyToken(token string) (map[string]interface{}, error)
}

func NewAutheticate(authenticator Authenticator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// ToDo
			authHeaderValue := r.Header.Get(authHeader)
			if authHeaderValue == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			token, err := parseBearerToken(authHeaderValue)

			if err != nil {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			claims, err := authenticator.VerifyToken(token)

			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			if err := verifyUserRole(claims); err != nil {
				http.Error(w, "invalid role", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func parseBearerToken(header string) (string, error) {
	if !strings.HasPrefix(header, authScheme) {
		return "", errors.New("invalid auth scheme")
	}

	token := strings.TrimPrefix(header, authScheme)

	if token == "" {
		return "", errors.New("empty token")
	}

	return token, nil
}

func verifyUserRole (claims map[string]interface{}) error {
	role, ok := claims["role"].(string)

	if !ok {
		return errors.New("invalid role claim")
	}

	if role == "authenticated" {
		return nil
	}

	return errors.New("invalid role")
}