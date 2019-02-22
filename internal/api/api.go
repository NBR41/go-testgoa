package api

import (
	"errors"
)

// ErrNoResult error when no result
var ErrNoResult = errors.New("no volume found")

//BookDetail struct for book detail
type BookDetail struct {
	Title       string
	Subtitle    *string
	Authors     []*Author
	Category    *string
	Editor      *string
	Description *string
	Collection  *string
	Print       *string
	Volume      *int
}

//Author Book Author
type Author struct {
	Name string
	Role *string
}

//BookGetter interface to get book detail by API
type BookGetter interface {
	GetBookDetail(isbn string) (*BookDetail, error)
}
