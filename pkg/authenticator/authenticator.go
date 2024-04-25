package authenticator

import "github.com/golang-jwt/jwt/v5"

type JWTAuthenticator struct {
	secret string
}

func NewJWTAuthenticator(secret string) JWTAuthenticator {
	return JWTAuthenticator{secret: secret}
}

func (a JWTAuthenticator) VerifyToken(token string) (map[string]interface{}, error) {
	var claims JwtClaims
	if _, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secret), nil
	}); err!= nil {
		return nil, err
	}

	result := map[string]interface{}{
		"role": claims.Role,
		"userID": claims.UserID,
		"email": claims.Email,
		"token": token,
	}

	return result, nil
}

type JwtClaims struct {
	jwt.RegisteredClaims

	Role string 	`json:"role"`
	UserID string `json:"sub"`
	Email string 	`json:"email"`
}