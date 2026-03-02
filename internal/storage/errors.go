package storage

import (
	"errors"
)

var (
	ErrNotFound     = errors.New("Short link not found")
	ErrInvalidShort = errors.New("Invalid short link")
)
