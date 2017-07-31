package appmodel

import (
	"github.com/pkg/errors"
	"os"
)

// List of errors
var (
	ErrNotFound           = errors.New("not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidID          = errors.New("invalid id")
	ErrDuplicateKey       = errors.New("duplicate key")
)

// User struct for users
type User struct {
	ID         int64
	Email      string
	Nickname   string
	IsVerified bool
	IsAdmin    bool
	password   []byte
	salt       []byte
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

// Modeler interface for model
type Modeler interface {
	Close() error
	GetUserList() ([]User, error)
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByNickname(nickname string) (*User, error)
	GetUserByEmailOrNickname(email, nickname string) (*User, error)
	GetAuthenticatedUser(login, password string) (*User, error)
	InsertUser(email, nickname, password string) (*User, error)
	UpdateUserNickname(id int, nickname string) error
	UpdateUserPassword(id int, password string) error
	UpdateUserActivation(id int, activated bool) error
	DeleteUser(id int) error
	InsertBook(isbn, name string) (*Book, error)
	GetBookByID(id int) (*Book, error)
	GetBookByName(name string) (*Book, error)
	GetBookList() ([]Book, error)
	UpdateBook(id int, name string) error
	DeleteBook(id int) error
	GetOwnershipList(userID int) ([]Ownership, error)
	GetOwnership(userID, bookID int) (*Ownership, error)
	InsertOwnership(userID, bookID int) (*Ownership, error)
	UpdateOwnership(userID, bookID int) error
	DeleteOwnership(userID, bookID int) error
}

var modeler Modeler

func GetModeler() (Modeler, error) {
	var err error
	if modeler == nil {
		switch os.Getenv("ISDEV") {
		case "1":
			modeler = NewLocal()
		default:
			connString := os.Getenv("DB_CONN_STR")
			if connString == "" {
				panic("no connexion string")
			}
			modeler, err = NewModel(connString)
		}
	}
	return modeler, err
}
