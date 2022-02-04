package token

import "context"

type (
	TokenCreator interface {
		GenerateAccessToken(ctx context.Context, ID, email, name string) (accessToken string, err error)
	}

	TokenVerificator interface{}

	TokenCreatorVerificator interface {
		TokenCreator
		TokenVerificator
	}
)
