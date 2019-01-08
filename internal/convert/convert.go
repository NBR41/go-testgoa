package convert

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/model"
)

//ToAuthorMedia converts an author model into an author media type
func ToAuthorMedia(a *model.Author) *app.Author {
	if a == nil {
		return nil
	}
	return &app.Author{
		Href:       app.AuthorsHref(a.ID),
		AuthorID:   int(a.ID),
		AuthorName: a.Name,
	}
}

//ToCategoryMedia converts a category model into a category media type
func ToCategoryMedia(a *model.Category) *app.Category {
	if a == nil {
		return nil
	}
	return &app.Category{
		Href:         app.CategoriesHref(a.ID),
		CategoryID:   int(a.ID),
		CategoryName: a.Name,
	}
}

//ToPrintMedia converts a print model into a print media type
func ToPrintMedia(a *model.Print) *app.Print {
	if a == nil {
		return nil
	}
	return &app.Print{
		Href:      app.PrintsHref(a.ID),
		PrintID:   int(a.ID),
		PrintName: a.Name,
	}
}

//ToEditorMedia converts an editor model into an editor media type
func ToEditorMedia(a *model.Editor) *app.Editor {
	if a == nil {
		return nil
	}
	return &app.Editor{
		Href:       app.EditorsHref(a.ID),
		EditorID:   int(a.ID),
		EditorName: a.Name,
	}
}

//ToClassMedia converts a class model into a class media type
func ToClassMedia(a *model.Class) *app.Class {
	if a == nil {
		return nil
	}
	return &app.Class{
		Href:      app.ClassesHref(a.ID),
		ClassID:   int(a.ID),
		ClassName: a.Name,
	}
}

//ToRoleMedia converts a role model into a role media type
func ToRoleMedia(a *model.Role) *app.Role {
	if a == nil {
		return nil
	}
	return &app.Role{
		Href:     app.RolesHref(a.ID),
		RoleID:   int(a.ID),
		RoleName: a.Name,
	}
}

// ToAuthTokenMedia converts a user model and token into a auth token media type
func ToAuthTokenMedia(a *model.User, accToken, refToken string) *app.Authtoken {
	if a == nil {
		return nil
	}
	return &app.Authtoken{
		User:         ToUserMedia(a),
		AccessToken:  accToken,
		RefreshToken: refToken,
	}
}

// ToBookMedia converts a book model into a book media type
func ToBookMedia(a *model.Book) *app.Book {
	if a == nil {
		return nil
	}
	return &app.Book{
		Href:     app.BooksHref(a.ID),
		BookID:   int(a.ID),
		BookName: a.Name,
		BookIsbn: a.ISBN,
	}
}

// ToBookLinkMedia converts a book model into a book link media type
func ToBookLinkMedia(a *model.Book) *app.BookLink {
	if a == nil {
		return nil
	}
	return &app.BookLink{
		Href:   app.BooksHref(a.ID),
		BookID: int(a.ID),
	}
}

// ToOwnershipMedia converts a book model into a book media type
func ToOwnershipMedia(a *model.Ownership) *app.Ownership {
	if a == nil {
		return nil
	}
	return &app.Ownership{
		Book:   ToBookMedia(a.Book),
		BookID: int(a.BookID),
		Href:   app.OwnershipsHref(a.UserID, a.BookID),
		UserID: int(a.UserID),
	}
}

// ToUserMedia converts a user model into a user media type
func ToUserMedia(a *model.User) *app.User {
	if a == nil {
		return nil
	}
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
	if a == nil {
		return nil
	}
	return &app.UserTiny{
		Href:     app.UsersHref(a.ID),
		UserID:   int(a.ID),
		Nickname: a.Nickname,
	}
}

//ToAuthorshipMedia converts an authorship model into an authorship media type
func ToAuthorshipMedia(a *model.Authorship) *app.Authorship {
	if a == nil {
		return nil
	}
	return &app.Authorship{
		AuthorshipID: int(a.ID),
		AuthorID:     int(a.AuthorID),
		Author:       ToAuthorMedia(a.Author),
		BookID:       int(a.BookID),
		Book:         ToBookMedia(a.Book),
		RoleID:       int(a.RoleID),
		Role:         ToRoleMedia(a.Role),
		Href:         app.AuthorshipsHref(a.ID),
	}
}

//ToAuthorshipMedia converts an authorship model into an authorship media type
func ToEditionMedia(a *model.Edition) *app.Edition {
	if a == nil {
		return nil
	}
	return &app.Edition{
		EditionID:    int(a.ID),
		BookID:       int(a.BookID),
		Book:         ToBookMedia(a.Book),
		CollectionID: int(a.CollectionID),
		Collection:   ToCollectionMedia(a.Collection),
		PrintID:      int(a.PrintID),
		Print:        ToPrintMedia(a.Print),
		Href:         app.EditionsHref(a.ID),
	}
}

// ToClassificationMedia converts a classification model into a classification media type
func ToClassificationMedia(seriesID int, a *model.Class) *app.Classification {
	if a == nil {
		return nil
	}
	return &app.Classification{
		Class: ToClassMedia(a),
		Href:  app.ClassificationsHref(seriesID, a.ID),
	}
}

// ToCollectionMedia converts a collection model into a collection media type
func ToCollectionMedia(a *model.Collection) *app.Collection {
	if a == nil {
		return nil
	}
	return &app.Collection{
		CollectionID:   int(a.ID),
		CollectionName: a.Name,
		EditorID:       int(a.EditorID),
		Editor:         ToEditorMedia(a.Editor),
		Href:           app.CollectionsHref(a.ID),
	}
}

// ToSeriesMedia converts a series model into a series media type
func ToSeriesMedia(a *model.Series) *app.Series {
	if a == nil {
		return nil
	}
	return &app.Series{
		SeriesID:   int(a.ID),
		SeriesName: a.Name,
		CategoryID: int(a.CategoryID),
		Category:   ToCategoryMedia(a.Category),
		Href:       app.SeriesHref(a.ID),
	}
}
