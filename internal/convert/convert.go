package convert

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/model"
)

// ToAuthTokenMedia converts a user model and token into a auth token media type
func ToAuthTokenMedia(a *model.User, accToken, refToken string) *app.Authtoken {
	return &app.Authtoken{
		User:         ToUserMedia(a),
		AccessToken:  accToken,
		RefreshToken: refToken,
	}
}

// ToBookMedia converts a book model into a book media type
func ToBookMedia(a *model.Book) *app.Book {
	return &app.Book{
		Href:     app.BooksHref(a.ID),
		BookID:   int(a.ID),
		BookName: a.Name,
		BookIsbn: a.ISBN,
	}
}

// ToBookLinkMedia converts a book model into a book link media type
func ToBookLinkMedia(a *model.Book) *app.BookLink {
	return &app.BookLink{
		Href:     app.BooksHref(a.ID),
		BookID:   int(a.ID),
		BookName: a.Name,
	}
}

// ToOwnershipMedia converts a book model into a book media type
func ToOwnershipMedia(a *model.Ownership) *app.Ownership {
	return &app.Ownership{
		Book:   ToBookMedia(a.Book),
		BookID: int(a.BookID),
		Href:   app.OwnershipsHref(a.UserID, a.BookID),
		UserID: int(a.UserID),
	}
}

// ToUserMedia converts a user model into a user media type
func ToUserMedia(a *model.User) *app.User {
	return &app.User{
		Email:       a.Email,
		Href:        app.UsersHref(a.ID),
		UserID:      int(a.ID),
		Nickname:    a.Nickname,
		IsAdmin:     a.IsAdmin,
		IsValidated: a.IsValidated,
	}
}

// ToUserTinyMedia converts a user model into a user media type
func ToUserTinyMedia(a *model.User) *app.UserTiny {
	return &app.UserTiny{
		Href:     app.UsersHref(a.ID),
		UserID:   int(a.ID),
		Nickname: a.Nickname,
	}
}
