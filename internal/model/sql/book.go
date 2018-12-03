package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getBook(query string, params ...interface{}) (*model.Book, error) {
	var b = model.Book{}
	err := m.db.QueryRow(query, params...).Scan(&b.ID, &b.ISBN, &b.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &b, nil
	}
}

// InsertBook inserts book
func (m *Model) InsertBook(isbn, name string) (*model.Book, error) {
	_, err := m.GetBookByISBN(isbn)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO books (book_id, isbn, name, create_ts, update_ts)
VALUES (null, ?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		isbn, name,
	)
	if err != nil {
		return nil, err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Book{ID: id, ISBN: isbn, Name: name}, nil
}

// GetBookByID returns book by ID
func (m *Model) GetBookByID(id int) (*model.Book, error) {
	return m.getBook(`SELECT book_id, isbn, name from books where id = ?`, id)
}

// GetBookByISBN returns book by ISBN
func (m *Model) GetBookByISBN(isbn string) (*model.Book, error) {
	return m.getBook(`SELECT book_id, isbn, name from books where isbn = ?`, isbn)
}

// GetBookByName returns book by name
func (m *Model) GetBookByName(name string) (*model.Book, error) {
	return m.getBook(`SELECT book_id, isbn, name from books where name = ?`, name)
}

// GetBookList returns book list
func (m *Model) GetBookList() ([]model.Book, error) {
	rows, err := m.db.Query(`SELECT book_id, isbn, name FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []model.Book{}
	for rows.Next() {
		b := model.Book{}
		if err := rows.Scan(&b.ID, &b.ISBN, &b.Name); err != nil {
			return nil, err
		}
		l = append(l, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// UpdateBook update book infos
func (m *Model) UpdateBook(id int, name string) error {
	return m.exec(
		`UPDATE books set name = ?, update_ts = NOW() where book_id = ?`,
		name, id,
	)
}

// DeleteBook delete book by ID
func (m *Model) DeleteBook(id int) error {
	return m.exec(`DELETE FROM books where book_id = ?`, id)
}
