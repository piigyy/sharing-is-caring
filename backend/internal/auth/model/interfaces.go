package model

import (
	"context"
)

type (
	RepositoryReader interface {
		GetUserByEmail(ctx context.Context, email string) (user User, err error)
	}

	RepositoryWriter interface{}

	RepositoryReaderWriter interface {
		RepositoryReader
		RepositoryWriter
	}
)

type (
	ServiceReader interface {
		Login(ctx context.Context, payload LoginRequest) (response LoginResponse, err error)
	}

	ServiceWriter interface{}

	ServiceReaderWriter interface {
		ServiceReader
		ServiceWriter
	}
)
