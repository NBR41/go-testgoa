package model

import (
	"github.com/pkg/errors"
)

// List of errors
var (
	ErrNotFound           = errors.New("not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidID          = errors.New("invalid id")
	ErrDuplicateKey       = errors.New("duplicate key")
	ErrDuplicateEmail     = errors.New("duplicate email")
	ErrDuplicateNickname  = errors.New("duplicate nickname")
)

type Passworder interface {
	CryptPassword(password string) ([]byte, []byte, error)
	ComparePassword(password string, salt, hash []byte) (bool, error)
}

// User struct for users
type User struct {
	ID          int64
	Email       string
	Nickname    string
	IsValidated bool
	IsAdmin     bool
	Password    []byte
	Salt        []byte
}

// Book struct for books
type Book struct {
	ID   int64
	ISBN string
	Name string
	URL  string
}

// Ownership struct for user book association
type Ownership struct {
	UserID int64
	BookID int64
	Book   *Book
}

//Category struct for category
type Category struct {
	ID   int64
	Name string
}

//Author struct for author
type Author struct {
	ID   int64
	Name string
}

//Genre struct for genre
type Genre struct {
	ID   int64
	Name string
}

//Role struct for role
type Role struct {
	ID   int64
	Name string
}

//EditionType struct for edition type
type EditionType struct {
	ID   int64
	Name string
}

//Editor struct for editor
type Editor struct {
	ID   int64
	Name string
}

type Collection struct {
	ID       int64
	Name     string
	EditorID int64
	Editor   *Editor
}
