package repository

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUserNotCreated = errors.New("user not created")
)
