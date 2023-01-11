package error

import "errors"

//TODO

var (
	// ErrInvalidToken returns when the token is invalid
	ErrInvalidToken = errors.New("invalid token")
)
