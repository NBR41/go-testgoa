package controllers

import (
	"context"
	"github.com/NBR41/go-testgoa/internal/model"
)

//CtxKey type for the keys stored in context
type CtxKey string

//Fmodeler function to get a Modeler
type Fmodeler func() (Modeler, error)

// Modeler interface for model
type Modeler interface {
	Close() error

	GetAuthorByID(id int) (*model.Author, error)
	GetAuthorByName(name string) (*model.Author, error)
	ListAuthorsByIDs(categoryID, roleID *int) ([]*model.Author, error)
	InsertAuthor(name string) (*model.Author, error)
	UpdateAuthor(id int, name string) error
	DeleteAuthor(id int) error

	GetBookByID(id int) (*model.Book, error)
	GetBookByISBN(isbn string) (*model.Book, error)
	GetBookByName(name string) (*model.Book, error)
	ListBooksByIDs(collectionID, editorID, printID, seriesID *int) ([]*model.Book, error)
	InsertBook(isbn, name string, seriesID int) (*model.Book, error)
	UpdateBook(id int, name *string, seriesID *int) error
	DeleteBook(id int) error

	GetCategoryByID(id int) (*model.Category, error)
	GetCategoryByName(name string) (*model.Category, error)
	ListCategoriesByIDs(authorID, classID *int) ([]*model.Category, error)
	InsertCategory(name string) (*model.Category, error)
	UpdateCategory(id int, name string) error
	DeleteCategory(id int) error

	GetClassByID(id int) (*model.Class, error)
	GetClassByName(name string) (*model.Class, error)
	ListClassesByIDs(authorID, categoryID, seriesID *int) ([]*model.Class, error)
	InsertClass(name string) (*model.Class, error)
	UpdateClass(id int, name string) error
	DeleteClass(id int) error

	GetCollectionByID(id int) (*model.Collection, error)
	GetCollectionByName(name string) (*model.Collection, error)
	ListCollectionsByIDs(editorID, printID, seriesID *int) ([]*model.Collection, error)
	InsertCollection(name string, editorID int) (*model.Collection, error)
	UpdateCollection(id int, name *string, editorID *int) error
	DeleteCollection(id int) error

	GetEditorByID(id int) (*model.Editor, error)
	GetEditorByName(name string) (*model.Editor, error)
	ListEditorsByIDs(printID, seriesID *int) ([]*model.Editor, error)
	InsertEditor(name string) (*model.Editor, error)
	UpdateEditor(id int, name string) error
	DeleteEditor(id int) error

	GetPrintByID(id int) (*model.Print, error)
	GetPrintByName(name string) (*model.Print, error)
	ListPrintsByIDs(collectionID, editorID, seriesID *int) ([]*model.Print, error)
	InsertPrint(name string) (*model.Print, error)
	UpdatePrint(id int, name string) error
	DeletePrint(id int) error

	GetRoleByID(id int) (*model.Role, error)
	GetRoleByName(name string) (*model.Role, error)
	ListRolesByIDs(authorID *int) ([]*model.Role, error)
	InsertRole(name string) (*model.Role, error)
	UpdateRole(id int, name string) error
	DeleteRole(id int) error

	GetSeriesByID(id int) (*model.Series, error)
	GetSeriesByName(name string) (*model.Series, error)
	ListSeriesByIDs(authorID, categoryID, classID, roleID *int) ([]*model.Series, error)
	ListSeriesByEditionIDs(collectionID, editorID, printID *int) ([]*model.Series, error)
	InsertSeries(name string, categoryID int) (*model.Series, error)
	UpdateSeries(id int, name *string, categoryID *int) error
	DeleteSeries(id int) error

	GetAuthorshipByID(id int) (*model.Authorship, error)
	ListAuthorships() ([]*model.Authorship, error)
	InsertAuthorship(authorID, bookID, roleID int) (*model.Authorship, error)
	DeleteAuthorship(id int) error

	GetClassification(seriesID, classID int) (*model.Class, error)
	InsertClassification(seriesID, classID int) (*model.Class, error)
	DeleteClassification(seriesID, classID int) error

	GetEditionByID(id int) (*model.Edition, error)
	ListEditions() ([]*model.Edition, error)
	InsertEdition(bookID, collectionID, printID int) (*model.Edition, error)
	DeleteEdition(id int) error

	GetOwnership(userID, bookID int) (*model.Ownership, error)
	ListOwnershipsByUserID(userID int) ([]*model.Ownership, error)
	InsertOwnership(userID, bookID int) (*model.Ownership, error)
	UpdateOwnership(userID, bookID int) error
	DeleteOwnership(userID, bookID int) error

	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByNickname(nickname string) (*model.User, error)
	GetUserByEmailOrNickname(email, nickname string) (*model.User, error)
	GetAuthenticatedUser(login, password string) (*model.User, error)
	ListUsers() ([]*model.User, error)
	InsertUser(email, nickname, password string) (*model.User, error)
	UpdateUserNickname(id int, nickname string) error
	UpdateUserPassword(id int, password string) error
	UpdateUserActivation(id int, activated bool) error
	DeleteUser(id int) error
}

//TokenHelper interface helper for tokens
type TokenHelper interface {
	GetPasswordToken(userID int64, email string) (string, error)
	ValidatePasswordToken(token string) (int64, string, error)
	GetValidationToken(userID int64, email string) (string, error)
	ValidateValidationToken(token string) (int64, string, error)
	GetAuthToken(userID int64, isAdmin bool) (string, error)
	GetRefreshToken(userID int64, isAdmin bool) (string, error)
	ValidateRefreshToken(token string) (int64, error)
}

//MailSender interface to send mails
type MailSender interface {
	SendResetPasswordMail(email, token string) error
	SendPasswordUpdatedMail(email string) error
	SendNewUserMail(u *model.User, token string) error
	SendActivationMail(u *model.User, token string) error
	SendUserActivatedMail(email string) error
}

//APIHelper interface to get book informations
type APIHelper interface {
	GetBookName(isbn string) (string, error)
}

//Lister interface to process list requests
type Lister interface {
	ListAuthors(ctx context.Context, fm Fmodeler, rCtx authorsResponse, categoryID, roleID *int) error
	ListBooks(ctx context.Context, fm Fmodeler, rCtx booksResponse, collectionID, editorID, printID, seriesID *int) error
	ListCategories(ctx context.Context, fm Fmodeler, rCtx categoriesResponse, authorID, classID *int) error
	ListClasses(ctx context.Context, fm Fmodeler, rCtx classesResponse, authorID, categoryID, seriesID *int) error
	ListCollections(ctx context.Context, fm Fmodeler, rCtx collectionsResponse, editorID, printID, seriesID *int) error
	ListEditors(ctx context.Context, fm Fmodeler, rCtx editorsResponse, printID, seriesID *int) error
	ListPrints(ctx context.Context, fm Fmodeler, rCtx printsResponse, collectionID, editorID, seriesID *int) error
	ListSeries(ctx context.Context, fm Fmodeler, rCtx seriesResponse, authorID, categoryID, classID, roleID *int) error
	ListSeriesByEditionIDs(ctx context.Context, fm Fmodeler, rCtx seriesResponse, collectionID, editorID, printID *int) error
}
