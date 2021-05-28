package service

import "errors"

type Error error

var (
	// ErrBadRequest
	ErrBadRequest = errors.New("bad request")
	// ErrRequestInValid
	ErrRequestInValid = errors.New("request invalid")
)
