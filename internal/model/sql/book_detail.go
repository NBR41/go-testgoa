package sql

import (
	"database/sql"
	"strconv"

	"github.com/NBR41/go-testgoa/internal/api"
	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetBookDetail = `
SELECT book.id, book.isbn, book.name,
series.id, series.name,
category.id, category.name,
edition.id,
print.id, print.name,
collection.id, collection.name,
editor.id, editor.name
FROM book
JOIN series ON (book.series_id = series.id)
JOIN category ON (series.category_id = category.id)
JOIN edition ON (edition.book_id = book.id)
JOIN print ON (edition.print_id = print.id)
JOIN collection ON (edition.collection_id = collection.id)
JOIN editor ON (collection.editor_id = editor.id)
WHERE book.isbn = ?`
)

//GetBookDetail return a Book detail by ISBN
func (m *Model) GetBookDetail(isbn string) (*model.BookDetail, error) {
	var b = model.BookDetail{
		Book: &model.Book{
			Series: &model.Series{
				Category: &model.Category{},
			},
		},
		Edition: &model.Edition{
			Print: &model.Print{},
			Collection: &model.Collection{
				Editor: &model.Editor{},
			},
		},
	}
	err := m.db.QueryRow(qryGetBookDetail, isbn).Scan(
		&b.Book.ID, &b.Book.ISBN, &b.Book.Name,
		&b.Book.Series.ID, &b.Book.Series.Name,
		&b.Book.Series.Category.ID, &b.Book.Series.Category.Name,
		&b.Edition.ID, &b.Edition.Print.ID, &b.Edition.Print.Name,
		&b.Edition.Collection.ID, &b.Edition.Collection.Name,
		&b.Edition.Collection.Editor.ID, &b.Edition.Collection.Editor.Name,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		b.Book.SeriesID = b.Book.Series.ID
		b.Book.Series.CategoryID = b.Book.Series.Category.ID
		b.Edition.BookID = b.Book.ID
		b.Edition.PrintID = b.Edition.Print.ID
		b.Edition.CollectionID = b.Edition.Collection.ID
		b.Edition.Collection.EditorID = b.Edition.Collection.Editor.ID

		seriesID := int(b.Book.SeriesID)
		b.Classes, err = m.ListClassesByIDs(nil, nil, &seriesID)
		if err != nil {
			return nil, err
		}
		b.Authors, err = m.ListAuthorshipsByBookID(int(b.Book.ID))
		if err != nil {
			return nil, err
		}
		return &b, nil
	}
}

//InsertBookDetail insert all values in book detail
func (m *Model) InsertBookDetail(isbn string, bookAPI *api.BookDetail) (*model.BookDetail, error) {
	var series *model.Series
	var editor *model.Editor
	var collection *model.Collection
	var print *model.Print
	var category *model.Category
	var book *model.Book
	var authorships []*model.Authorship
	var err error

	if bookAPI.Print != nil {
		print, err = m.getOrInsertPrint(*bookAPI.Print)
		if err != nil {
			return nil, err
		}
	} else {
		print = &model.Print{ID: model.UndefinedID, Name: model.UndefinedName}
	}

	if bookAPI.Category != nil {
		category, err = m.getOrInsertCategory(*bookAPI.Category)
		if err != nil {
			return nil, err
		}
	} else {
		category = &model.Category{ID: model.UndefinedID, Name: model.UndefinedName}
	}
	if bookAPI.Editor != nil {
		editor, err = m.getOrInsertEditor(*bookAPI.Editor)
		if err != nil {
			return nil, err
		}
	} else {
		editor = &model.Editor{ID: model.UndefinedID, Name: model.UndefinedName}
	}

	if bookAPI.Collection != nil {
		collection, err = m.getOrInsertCollection(*bookAPI.Collection, int(editor.ID))
	} else {
		collection, err = m.getOrInsertCollection(model.UndefinedName, int(editor.ID))
	}
	if err != nil {
		return nil, err
	}
	collection.Editor = editor

	series, err = m.getOrInsertSeries(bookAPI.Title, int(category.ID))
	if err != nil {
		return nil, err
	}
	series.Category = category

	var volume string
	if bookAPI.Volume != nil {
		volume = strconv.Itoa(*bookAPI.Volume)
	} else {
		volume = ""
	}

	book, err = m.InsertBook(isbn, volume, int(series.ID))
	if err != nil {
		return nil, err
	}
	book.Series = series

	edition, err := m.InsertEdition(int(book.ID), int(collection.ID), int(print.ID))
	if err != nil {
		return nil, err
	}
	edition.Print = print

	if len(bookAPI.Authors) > 0 {
		for i := range bookAPI.Authors {
			var role *model.Role
			var author *model.Author
			if bookAPI.Authors[i].Role != nil {
				role, err = m.getOrInsertRole(*bookAPI.Authors[i].Role)
				if err != nil {
					return nil, err
				}
			} else {
				role = &model.Role{ID: model.UndefinedID, Name: model.UndefinedName}
			}

			author, err = m.getOrInsertAuthor(bookAPI.Authors[i].Name)
			if err != nil {
				return nil, err
			}
			authorship, err := m.InsertAuthorship(int(author.ID), int(book.ID), int(role.ID))
			if err != nil {
				return nil, err
			}
			authorship.Role = role
			authorship.Author = author
			authorships = append(authorships, authorship)
			if err != nil {
				return nil, err
			}
		}
	} else {
		authorship, err := m.InsertAuthorship(model.UndefinedID, int(book.ID), model.UndefinedID)
		if err != nil {
			return nil, err
		}
		authorships = append(authorships, authorship)
	}

	return &model.BookDetail{
		Book:    book,
		Edition: edition,
		Classes: []*model.Class{},
		Authors: authorships,
	}, nil
}
