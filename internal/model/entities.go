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

//Class struct for class
type Class struct {
	ID   int64
	Name string
}

//Role struct for role
type Role struct {
	ID   int64
	Name string
}

//Print struct for print
type Print struct {
	ID   int64
	Name string
}

//Editor struct for editor
type Editor struct {
	ID   int64
	Name string
}

//Collection struct for collection
type Collection struct {
	ID       int64
	Name     string
	EditorID int64
	Editor   *Editor
}

//Authorship struct for authorship
type Authorship struct {
	ID       int64
	AuthorID int64
	Author   *Author
	BookID   int64
	Book     *Book
	RoleID   int64
	Role     *Role
}

//Edition struct for edition
type Edition struct {
	ID           int64
	BookID       int64
	Book         *Book
	CollectionID int64
	Collection   *Collection
	PrintID      int64
	Print        *Print
}
