package token

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type (
	TokenCreator interface {
		GenerateAccessToken(ctx context.Context, ID, email, name string) (accessToken string, err error)
	}

	TokenVerificator interface {
		VerifyToken(accessToken string) (*jwt.Token, error)
		ValidateToken(signedToken *jwt.Token) error
	}

	TokenCreatorVerificator interface {
		TokenCreator
		TokenVerificator
	}
)
