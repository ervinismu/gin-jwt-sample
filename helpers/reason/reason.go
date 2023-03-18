package helpers

import "errors"

var (
	ErrFailedRegisterUser = errors.New("failed to register user")
	ErrFailedLoginUser = errors.New("failed to login user")
)
