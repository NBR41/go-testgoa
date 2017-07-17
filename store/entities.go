package store

import (
	"github.com/pkg/errors"
	"os"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrBookNotFound       = errors.New("book not found")
	ErrOwnershipNotFound  = errors.New("ownership not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateKey       = errors.New("duplicate key")
)

// User struct for users
type User struct {
	ID         int64
	Email      string
	Nickname   string
	Password   string
	IsVerified bool
	IsAdmin    bool
}

// Book struct for books
type Book struct {
	ID   int64
	Name string
	URL  string
}

// Ownership struct for user book association
type Ownership struct {
	UserID int64
	BookID int64
	Book   *Book
}

type Modeler interface {
	Close() error
	GetUserList() ([]User, error)
	GetUserByID(id int) (*User, error)
	GetUserByEmailOrNickname(email, nickname string) (*User, error)
	GetAuthenticatedUser(password, email, nickname string) (*User, error)
	InsertUser(nickname, email, password string) (*User, error)
	UpdateUserNickname(id int, nickname string) error
	UpdateUserPassword(id int, password string) error
	UpdateUserActivation(id int, activated bool) error
	DeleteUser(id int) error
	InsertBook(name string) (*Book, error)
	GetBookByID(id int) (*Book, error)
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
