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
