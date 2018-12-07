package controllers

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

type CtxKey string

type Fmodeler func() (Modeler, error)

// Modeler interface for model
type Modeler interface {
	Close() error

	GetAuthorByID(id int) (*model.Author, error)
	GetAuthorByName(name string) (*model.Author, error)
	GetAuthorList() ([]*model.Author, error)
	InsertAuthor(name string) (*model.Author, error)
	UpdateAuthor(id int, name string) error
	DeleteAuthor(id int) error

	GetBookByID(id int) (*model.Book, error)
	GetBookByISBN(isbn string) (*model.Book, error)
	GetBookByName(name string) (*model.Book, error)
	GetBookList() ([]model.Book, error)
	InsertBook(isbn, name string) (*model.Book, error)
	UpdateBook(id int, name string) error
	DeleteBook(id int) error

	GetCategoryByID(id int) (*model.Category, error)
	GetCategoryByName(name string) (*model.Category, error)
	GetCategoryList() ([]*model.Category, error)
	InsertCategory(name string) (*model.Category, error)
	UpdateCategory(id int, name string) error
	DeleteCategory(id int) error

	GetEditionTypeByID(id int) (*model.EditionType, error)
	GetEditionTypeByName(name string) (*model.EditionType, error)
	GetEditionTypeList() ([]*model.EditionType, error)
	InsertEditionType(name string) (*model.EditionType, error)
	UpdateEditionType(id int, name string) error
	DeleteEditionType(id int) error

	GetEditorByID(id int) (*model.Editor, error)
	GetEditorByName(name string) (*model.Editor, error)
	GetEditorList() ([]*model.Editor, error)
	InsertEditor(name string) (*model.Editor, error)
	UpdateEditor(id int, name string) error
	DeleteEditor(id int) error

	GetGenreByID(id int) (*model.Genre, error)
	GetGenreByName(name string) (*model.Genre, error)
	GetGenreList() ([]*model.Genre, error)
	InsertGenre(name string) (*model.Genre, error)
	UpdateGenre(id int, name string) error
	DeleteGenre(id int) error

	GetOwnership(userID, bookID int) (*model.Ownership, error)
	GetOwnershipList(userID int) ([]*model.Ownership, error)
	InsertOwnership(userID, bookID int) (*model.Ownership, error)
	UpdateOwnership(userID, bookID int) error
	DeleteOwnership(userID, bookID int) error

	GetRoleByID(id int) (*model.Role, error)
	GetRoleByName(name string) (*model.Role, error)
	GetRoleList() ([]*model.Role, error)
	InsertRole(name string) (*model.Role, error)
	UpdateRole(id int, name string) error
	DeleteRole(id int) error

	GetUserList() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByNickname(nickname string) (*model.User, error)
	GetUserByEmailOrNickname(email, nickname string) (*model.User, error)
	GetAuthenticatedUser(login, password string) (*model.User, error)
	InsertUser(email, nickname, password string) (*model.User, error)
	UpdateUserNickname(id int, nickname string) error
	UpdateUserPassword(id int, password string) error
	UpdateUserActivation(id int, activated bool) error
	DeleteUser(id int) error
}

type TokenHelper interface {
	GetPasswordToken(userID int64, email string) (string, error)
	ValidatePasswordToken(token string) (int64, string, error)
	GetValidationToken(userID int64, email string) (string, error)
	ValidateValidationToken(token string) (int64, string, error)
	GetAuthToken(userID int64, isAdmin bool) (string, error)
	GetRefreshToken(userID int64, isAdmin bool) (string, error)
	ValidateRefreshToken(token string) (int64, error)
}

type MailSender interface {
	SendResetPasswordMail(email, token string) error
	SendPasswordUpdatedMail(email string) error
	SendNewUserMail(u *model.User, token string) error
	SendActivationMail(u *model.User, token string) error
	SendUserActivatedMail(email string) error
}

type APIHelper interface {
	GetBookName(isbn string) (string, error)
}
