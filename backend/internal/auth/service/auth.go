package service

import (
	"context"
	"errors"
	"time"

	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"github.com/piigyy/sharing-is-caring/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type auth struct {
	authRepository model.RepositoryReaderWriter
	tokenCreator   token.TokenCreator
}

func NewAuthService(authRepository model.RepositoryReaderWriter, tokenCreator token.TokenCreator) *auth {
	return &auth{
		authRepository: authRepository,
		tokenCreator:   tokenCreator,
	}
}

func (s *auth) Login(ctx context.Context, payload model.LoginRequest) (response model.LoginResponse, err error) {
	var user model.User
	var accessToken string

	user, err = s.authRepository.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			err = model.ErrUserNotFound
			return
		}
		return
	}

	if comparePasErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); comparePasErr != nil {
		err = model.ErrUnAuthorized
		return
	}

	accessToken, err = s.tokenCreator.GenerateAccessToken(ctx, user.ID.Hex(), user.Name, user.Name)
	if err != nil {
		return
	}

	return model.LoginResponse{
		ID:                   user.ID.Hex(),
		AccessToken:          accessToken,
		AccessTokenExpiredAt: time.Now().Add(2 * time.Hour),
		Email:                payload.Email,
	}, nil
}
