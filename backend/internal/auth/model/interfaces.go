package model

import (
	"context"
)

type (
	RepositoryReader interface {
		GetUserByEmail(ctx context.Context, email string) (user User, err error)
		DuplicateError(ctx context.Context, err error) bool
	}

	RepositoryWriter interface {
		CreateUser(ctx context.Context, entity User) (userID string, err error)
	}

	RepositoryReaderWriter interface {
		RepositoryReader
		RepositoryWriter
	}
)

type (
	ServiceReader interface {
		Login(ctx context.Context, payload LoginRequest) (response LoginResponse, err error)
		GetUserDetailByEmail(ctx context.Context, email string) (user User, err error)
	}

	ServiceWriter interface {
		RegisterUser(ctx context.Context, payload RegisterUserRequest) (response LoginResponse, err error)
	}

	ServiceReaderWriter interface {
		ServiceReader
		ServiceWriter
	}
)
