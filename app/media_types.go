// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "my-inventory": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/NBR41/go-testgoa/design
// --out=$(GOPATH)/src/github.com/NBR41/go-testgoa
// --version=v1.3.0

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// An auth token (default view)
//
// Identifier: application/vnd.authtoken+json; view=default
type Authtoken struct {
	// Unique user ID
	Token string `form:"token" json:"token" xml:"token"`
	// user struct
	User *User `form:"user" json:"user" xml:"user"`
}

// Validate validates the Authtoken media type instance.
func (mt *Authtoken) Validate() (err error) {
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	if utf8.RuneCountInString(mt.Token) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.token`, mt.Token, utf8.RuneCountInString(mt.Token), 1, true))
	}
	if mt.User != nil {
		if err2 := mt.User.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// A Book (default view)
//
// Identifier: application/vnd.book+json; view=default
type Book struct {
	// API href for making requests on the book
	Href string `form:"href" json:"href" xml:"href"`
	// Unique Book ID
	ID int `form:"id" json:"id" xml:"id"`
	// Book ISBN
	Isbn *string `form:"isbn,omitempty" json:"isbn,omitempty" xml:"isbn,omitempty"`
	// Book Name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Book media type instance.
func (mt *Book) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.ID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 1, true))
	}
	if mt.Isbn != nil {
		if utf8.RuneCountInString(*mt.Isbn) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.isbn`, *mt.Isbn, utf8.RuneCountInString(*mt.Isbn), 1, true))
		}
	}
	if mt.Isbn != nil {
		if utf8.RuneCountInString(*mt.Isbn) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.isbn`, *mt.Isbn, utf8.RuneCountInString(*mt.Isbn), 128, false))
		}
	}
	if utf8.RuneCountInString(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 1, true))
	}
	if utf8.RuneCountInString(mt.Name) > 128 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 128, false))
	}
	return
}

// A Book (link view)
//
// Identifier: application/vnd.book+json; view=link
type BookLink struct {
	// API href for making requests on the book
	Href string `form:"href" json:"href" xml:"href"`
	// Unique Book ID
	ID int `form:"id" json:"id" xml:"id"`
	// Book ISBN
	Isbn *string `form:"isbn,omitempty" json:"isbn,omitempty" xml:"isbn,omitempty"`
	// Book Name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the BookLink media type instance.
func (mt *BookLink) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.ID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 1, true))
	}
	if mt.Isbn != nil {
		if utf8.RuneCountInString(*mt.Isbn) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.isbn`, *mt.Isbn, utf8.RuneCountInString(*mt.Isbn), 1, true))
		}
	}
	if mt.Isbn != nil {
		if utf8.RuneCountInString(*mt.Isbn) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.isbn`, *mt.Isbn, utf8.RuneCountInString(*mt.Isbn), 128, false))
		}
	}
	if utf8.RuneCountInString(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 1, true))
	}
	if utf8.RuneCountInString(mt.Name) > 128 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 128, false))
	}
	return
}

// BookCollection is the media type for an array of Book (default view)
//
// Identifier: application/vnd.book+json; type=collection; view=default
type BookCollection []*Book

// Validate validates the BookCollection media type instance.
func (mt BookCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// BookCollection is the media type for an array of Book (link view)
//
// Identifier: application/vnd.book+json; type=collection; view=link
type BookLinkCollection []*BookLink

// Validate validates the BookLinkCollection media type instance.
func (mt BookLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A User ownership (default view)
//
// Identifier: application/vnd.ownership+json; view=default
type Ownership struct {
	// book struct
	Book *Book `form:"book,omitempty" json:"book,omitempty" xml:"book,omitempty"`
	// Unique Book ID
	BookID int `form:"book_id" json:"book_id" xml:"book_id"`
	// API href for making requests on the ownership
	Href string `form:"href" json:"href" xml:"href"`
	// Unique User ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the Ownership media type instance.
func (mt *Ownership) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Book != nil {
		if err2 := mt.Book.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.BookID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.book_id`, mt.BookID, 1, true))
	}
	if mt.UserID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.user_id`, mt.UserID, 1, true))
	}
	return
}

// OwnershipCollection is the media type for an array of Ownership (default view)
//
// Identifier: application/vnd.ownership+json; type=collection; view=default
type OwnershipCollection []*Ownership

// Validate validates the OwnershipCollection media type instance.
func (mt OwnershipCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A token (default view)
//
// Identifier: application/vnd.token+json; view=default
type Token struct {
	// Unique user ID
	Token string `form:"token" json:"token" xml:"token"`
}

// Validate validates the Token media type instance.
func (mt *Token) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	if utf8.RuneCountInString(mt.Token) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.token`, mt.Token, utf8.RuneCountInString(mt.Token), 1, true))
	}
	return
}

// A User (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// user email
	Email string `form:"email" json:"email" xml:"email"`
	// API href for making requests on the user
	Href string `form:"href" json:"href" xml:"href"`
	// Unique User ID
	ID          int  `form:"id" json:"id" xml:"id"`
	IsAdmin     bool `form:"is_admin" json:"is_admin" xml:"is_admin"`
	IsValidated bool `form:"is_validated" json:"is_validated" xml:"is_validated"`
	// user nickname
	Nickname string `form:"nickname" json:"nickname" xml:"nickname"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Nickname == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "nickname"))
	}

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	if mt.ID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 1, true))
	}
	if utf8.RuneCountInString(mt.Nickname) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.nickname`, mt.Nickname, utf8.RuneCountInString(mt.Nickname), 1, true))
	}
	if utf8.RuneCountInString(mt.Nickname) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.nickname`, mt.Nickname, utf8.RuneCountInString(mt.Nickname), 32, false))
	}
	return
}

// A User (tiny view)
//
// Identifier: application/vnd.user+json; view=tiny
type UserTiny struct {
	// API href for making requests on the user
	Href string `form:"href" json:"href" xml:"href"`
	// Unique User ID
	ID int `form:"id" json:"id" xml:"id"`
	// user nickname
	Nickname string `form:"nickname" json:"nickname" xml:"nickname"`
}

// Validate validates the UserTiny media type instance.
func (mt *UserTiny) Validate() (err error) {

	if mt.Nickname == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "nickname"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.ID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 1, true))
	}
	if utf8.RuneCountInString(mt.Nickname) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.nickname`, mt.Nickname, utf8.RuneCountInString(mt.Nickname), 1, true))
	}
	if utf8.RuneCountInString(mt.Nickname) > 32 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.nickname`, mt.Nickname, utf8.RuneCountInString(mt.Nickname), 32, false))
	}
	return
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// UserCollection is the media type for an array of User (tiny view)
//
// Identifier: application/vnd.user+json; type=collection; view=tiny
type UserTinyCollection []*UserTiny

// Validate validates the UserTinyCollection media type instance.
func (mt UserTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
