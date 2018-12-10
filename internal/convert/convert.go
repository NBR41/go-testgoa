package convert

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/model"
)

//ToAuthorMedia converts an author model into an author media type
func ToAuthorMedia(a *model.Author) *app.Author {
	return &app.Author{
		Href:       app.AuthorsHref(a.ID),
		AuthorID:   int(a.ID),
		AuthorName: a.Name,
	}
}

//ToCategoryMedia converts a category model into a category media type
func ToCategoryMedia(a *model.Category) *app.Category {
	return &app.Category{
		Href:         app.CategoriesHref(a.ID),
		CategoryID:   int(a.ID),
		CategoryName: a.Name,
	}
}

//ToEditionTypeMedia converts an edition type model into an edition type media type
func ToEditionTypeMedia(a *model.EditionType) *app.Editiontype {
	return &app.Editiontype{
		Href:            app.EditionTypesHref(a.ID),
		EditionTypeID:   int(a.ID),
		EditionTypeName: a.Name,
	}
}

//ToEditorMedia converts an editor model into an editor media type
func ToEditorMedia(a *model.Editor) *app.Editor {
	return &app.Editor{
		Href:       app.EditorsHref(a.ID),
		EditorID:   int(a.ID),
		EditorName: a.Name,
	}
}

//ToGenreMedia converts a genre model into a genre media type
func ToGenreMedia(a *model.Genre) *app.Genre {
	return &app.Genre{
		Href:      app.GenresHref(a.ID),
		GenreID:   int(a.ID),
		GenreName: a.Name,
	}
}

//ToRoleMedia converts a role model into a role media type
func ToRoleMedia(a *model.Role) *app.Role {
	return &app.Role{
		Href:     app.GenresHref(a.ID),
		RoleID:   int(a.ID),
		RoleName: a.Name,
	}
}

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
