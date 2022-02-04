package model

import (
	"context"
)

type (
	Reader interface {
		GetUserByEmail(ctx context.Context, email string) (user User, err error)
	}

	Writer interface{}

	ReaderWriter interface {
		Reader
		Writer
	}
)
