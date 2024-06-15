package storage

import "errors"


var (
	ErrUserExists = errors.New("user alreaby exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound = errors.New("app not found")
) 