// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "my-inventory": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/NBR41/go-testgoa/design
// --out=$(GOPATH)/src/github.com/NBR41/go-testgoa
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// AuthorsHref returns the resource href.
func AuthorsHref(authorID interface{}) string {
	paramauthorID := strings.TrimLeftFunc(fmt.Sprintf("%v", authorID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/authors/%v", paramauthorID)
}

// AuthorshipsHref returns the resource href.
func AuthorshipsHref(authorshipID interface{}) string {
	paramauthorshipID := strings.TrimLeftFunc(fmt.Sprintf("%v", authorshipID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/authorships/%v", paramauthorshipID)
}

// BooksHref returns the resource href.
func BooksHref(bookID interface{}) string {
	parambookID := strings.TrimLeftFunc(fmt.Sprintf("%v", bookID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/books/%v", parambookID)
}

// CategoriesHref returns the resource href.
func CategoriesHref(categoryID interface{}) string {
	paramcategoryID := strings.TrimLeftFunc(fmt.Sprintf("%v", categoryID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/categories/%v", paramcategoryID)
}

// ClassesHref returns the resource href.
func ClassesHref(classID interface{}) string {
	paramclassID := strings.TrimLeftFunc(fmt.Sprintf("%v", classID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/classes/%v", paramclassID)
}

// EditionsHref returns the resource href.
func EditionsHref(editionID interface{}) string {
	parameditionID := strings.TrimLeftFunc(fmt.Sprintf("%v", editionID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/editions/%v", parameditionID)
}

// EditorsHref returns the resource href.
func EditorsHref(editorID interface{}) string {
	parameditorID := strings.TrimLeftFunc(fmt.Sprintf("%v", editorID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/editors/%v", parameditorID)
}

// CollectionsHref returns the resource href.
func CollectionsHref(editorID, collectionID interface{}) string {
	parameditorID := strings.TrimLeftFunc(fmt.Sprintf("%v", editorID), func(r rune) bool { return r == '/' })
	paramcollectionID := strings.TrimLeftFunc(fmt.Sprintf("%v", collectionID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/editors/%v/collections/%v", parameditorID, paramcollectionID)
}

// OwnershipsHref returns the resource href.
func OwnershipsHref(userID, bookID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	parambookID := strings.TrimLeftFunc(fmt.Sprintf("%v", bookID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/users/%v/ownerships/%v", paramuserID, parambookID)
}

// PrintsHref returns the resource href.
func PrintsHref(printID interface{}) string {
	paramprintID := strings.TrimLeftFunc(fmt.Sprintf("%v", printID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/prints/%v", paramprintID)
}

// RolesHref returns the resource href.
func RolesHref(roleID interface{}) string {
	paramroleID := strings.TrimLeftFunc(fmt.Sprintf("%v", roleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/roles/%v", paramroleID)
}

// SeriesHref returns the resource href.
func SeriesHref(seriesID interface{}) string {
	paramseriesID := strings.TrimLeftFunc(fmt.Sprintf("%v", seriesID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/series/%v", paramseriesID)
}

// ClassificationsHref returns the resource href.
func ClassificationsHref(seriesID, classID interface{}) string {
	paramseriesID := strings.TrimLeftFunc(fmt.Sprintf("%v", seriesID), func(r rune) bool { return r == '/' })
	paramclassID := strings.TrimLeftFunc(fmt.Sprintf("%v", classID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/series/%v/classifications/%v", paramseriesID, paramclassID)
}

// UsersHref returns the resource href.
func UsersHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/users/%v", paramuserID)
}
