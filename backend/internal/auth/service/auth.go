package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"github.com/piigyy/sharing-is-caring/pkg/token"
	"go.mongodb.org/mongo-driver/mongo"
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

	accessToken, err = s.tokenCreator.GenerateAccessToken(ctx, user.ID.Hex(), user.Email, user.Name)
	if err != nil {
		return
	}

	return model.LoginResponse{
		ID:                   user.ID.Hex(),
		AccessToken:          accessToken,
		AccessTokenExpiredAt: time.Now().Add(2 * time.Hour),
		Email:                payload.Email,
		Name:                 user.Name,
	}, nil
}

func (s *auth) RegisterUser(ctx context.Context, payload model.RegisterUserRequest) (response model.LoginResponse, err error) {
	var (
		passwordHashed      []byte
		userID, accessToken string
	)

	log.Println("hashing user password")
	passwordHashed, err = bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
	if err != nil {
		return
	}

	user := model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Phone:    payload.Phone,
		Password: string(passwordHashed),
		CreateAt: time.Now(),
		Updated:  time.Now(),
	}

	log.Println("saving user entity")
	userID, err = s.authRepository.CreateUser(ctx, user)
	if err != nil {
		log.Printf("error trying to save user entity to mongodb: %v\n", err)
		if s.authRepository.DuplicateError(ctx, err) {
			err = model.ErrUserDuplicated
			return
		}

		return
	}

	log.Println("generating user access token")
	accessToken, err = s.tokenCreator.GenerateAccessToken(ctx, userID, user.Email, user.Name)
	if err != nil {
		return
	}

	return model.LoginResponse{
		ID:                   userID,
		Email:                user.Email,
		AccessToken:          accessToken,
		AccessTokenExpiredAt: time.Now().Add(12 * time.Hour),
	}, nil
}

func (s *auth) GetUserDetailByEmail(ctx context.Context, email string) (user model.User, err error) {
	user, err = s.authRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return
	}
	return
}

func (s *auth) UpdatePassword(ctx context.Context, payload model.UpdatePasswordRequest) error {
	user, err := s.GetUserDetailByEmail(ctx, payload.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.ErrUnAuthorized
		}
		log.Printf("unexpected error auth.UpdatePassword: %v\n", err)
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.OldPassword)); err != nil {
		return model.ErrUnAuthorized
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(payload.NewPassword), bcrypt.MinCost)
	if err != nil {
		log.Printf("unexpected error auth.UpdatePassword: %v\n", err)
		return err
	}

	return s.authRepository.UpdatePassword(ctx, payload.Email, string(newPassword))
}
