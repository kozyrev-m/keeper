package httpserver

import "errors"

var (
	errIncorrectLoginOrPassword = errors.New("incorrect login or password")
	errNotAuthenticated = errors.New("not authenticated")
)