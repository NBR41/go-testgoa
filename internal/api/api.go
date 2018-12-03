package api

import (
	"errors"
)

// ErrNoResult error when no result
var ErrNoResult = errors.New("no volume found")

type BookGetter interface {
	GetBookName(isbn string) (string, error)
}
