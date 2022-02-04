package model

import "errors"

var (
	ErrUserNotFound error = errors.New("user not found")
	ErrUnAuthorized error = errors.New("authorization is failed")
)
