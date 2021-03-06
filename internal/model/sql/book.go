package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetBookByID   = `SELECT id, isbn, name, series_id from book WHERE id = ?`
	qryGetBookByISBN = `SELECT id, isbn, name, series_id from book WHERE isbn = ?`
	qryGetBookByName = `SELECT id, isbn, name, series_id from book WHERE name = ?`
	qryListBooks     = `SELECT id, isbn, name, series_id FROM book`
	qryInsertBook    = `
INSERT INTO book (id, isbn, name, series_id, create_ts, update_ts)
VALUES (null, ?, ?, ?, NOW(), NOW())`
	qryDeleteBook = `DELETE FROM book WHERE id = ?`
)

func (m *Model) getBook(query string, params ...interface{}) (*model.Book, error) {
	var b = model.Book{}
	err := m.db.QueryRow(query, params...).Scan(&b.ID, &b.ISBN, &b.Name, &b.SeriesID)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &b, nil
	}
}

func (m *Model) listBooks(query string, params ...interface{}) ([]*model.Book, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Book{}
	for rows.Next() {
		b := &model.Book{}
		if err := rows.Scan(&b.ID, &b.ISBN, &b.Name, &b.SeriesID); err != nil {
			return nil, err
		}
		l = append(l, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// GetBookByID returns book by ID
func (m *Model) GetBookByID(id int) (*model.Book, error) {
	return m.getBook(qryGetBookByID, id)
}

// GetBookByISBN returns book by ISBN
func (m *Model) GetBookByISBN(isbn string) (*model.Book, error) {
	return m.getBook(qryGetBookByISBN, isbn)
}

// GetBookByName returns book by name
func (m *Model) GetBookByName(name string) (*model.Book, error) {
	return m.getBook(qryGetBookByName, name)
}

// ListBooksByIDs returns book list filtered by collection or editor or print or series
func (m *Model) ListBooksByIDs(collectionID, editorID, printID, seriesID *int) ([]*model.Book, error) {
	qry := `SELECT DISTINCT book.id, book.isbn, book.name, book.series_id FROM book`
	where := []string{"1"}
	vals := []interface{}{}
	if collectionID != nil || editorID != nil || printID != nil {
		qry += ` JOIN edition ON (edition.book_id = book.id)`
		if collectionID != nil {
			where = append(where, `edition.collection_id = ?`)
			vals = append(vals, *collectionID)
		}
		if editorID != nil {
			qry += ` JOIN collection ON (edition.collection_id = collection.id)`
			where = append(where, `collection.editor_id = ?`)
			vals = append(vals, *editorID)
		}
		if printID != nil {
			where = append(where, `edition.print_id = ?`)
			vals = append(vals, *printID)
		}
	}
	if seriesID != nil {
		where = append(where, `book.series_id = ?`)
		vals = append(vals, *seriesID)
	}
	return m.listBooks(qry+` WHERE `+strings.Join(where, " AND "), vals...)
}

// InsertBook inserts book
func (m *Model) InsertBook(isbn, name string, seriesID int) (*model.Book, error) {
	res, err := m.db.Exec(qryInsertBook, isbn, name, seriesID)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Book{ID: id, ISBN: isbn, Name: name, SeriesID: int64(seriesID)}, nil
}

// UpdateBook update book infos
func (m *Model) UpdateBook(id int, name *string, seriesID *int) error {
	if name == nil && seriesID == nil {
		return nil
	}
	cols := []string{}
	vals := []interface{}{}
	if name != nil {
		cols = append(cols, "name = ?")
		vals = append(vals, *name)
	}
	if seriesID != nil {
		cols = append(cols, "series_id = ?")
		vals = append(vals, *seriesID)
	}
	vals = append(vals, id)
	return m.exec(
		`UPDATE book SET `+strings.Join(cols, ", ")+`, update_ts = NOW() WHERE id = ?`,
		vals...,
	)
}

// DeleteBook delete book by ID
func (m *Model) DeleteBook(id int) error {
	return m.exec(qryDeleteBook, id)
}
