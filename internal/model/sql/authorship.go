package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetAuthorshipByID = `
SELECT authorship.id, author.id, author.name, role.id, role.name, book.id, book.name
FROM authorship au
JOIN author ON (authorship.author_id = author.id)
JOIN role ON (authorship.role_id = role.id)
JOIN book ON (authorship.book_id = book.id)
WHERE authorship.id = ?`
	qryListAuthorships = `
SELECT authorship.id, author.id, author.name, role.id, role.name, book.id, book.name
FROM authorship
JOIN author ON (authorship.author_id = author.id)
JOIN role ON (authorship.role_id = role.id)
JOIN book ON (authorship.book_id = book.id)`
	qryInsertAuthorship = `
INSERT INTO authorship (id, author_id, book_id, role_id, create_ts, update_ts)
VALUES (null, ?, ?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
	qryDeleteAuthorship = `DELETE FROM authorship WHERE id = ?`
)

func (m *Model) getAuthorship(query string, params ...interface{}) (*model.Authorship, error) {
	var v = model.Authorship{Author: &model.Author{}, Role: &model.Role{}, Book: &model.Book{}}
	err := m.db.QueryRow(query, params...).Scan(&v.ID, &v.Author.ID, &v.Author.Name, &v.Role.ID, &v.Role.Name, &v.Book.ID, &v.Book.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		v.AuthorID = v.Author.ID
		v.RoleID = v.Role.ID
		v.BookID = v.Book.ID
		return &v, nil
	}
}

func (m *Model) listAuthorships(query string, params ...interface{}) ([]*model.Authorship, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Authorship{}
	for rows.Next() {
		v := &model.Authorship{Author: &model.Author{}, Role: &model.Role{}, Book: &model.Book{}}
		if err := rows.Scan(&v.ID, &v.Author.ID, &v.Author.Name, &v.Role.ID, &v.Role.Name, &v.Book.ID, &v.Book.Name); err != nil {
			return nil, err
		}
		v.AuthorID = v.Author.ID
		v.RoleID = v.Role.ID
		v.BookID = v.Book.ID
		l = append(l, v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

//GetAuthorshipByID returns an authorship by id
func (m *Model) GetAuthorshipByID(id int) (*model.Authorship, error) {
	return m.getAuthorship(qryGetAuthorshipByID, id)
}

//ListAuthorships list all authorships
func (m *Model) ListAuthorships() ([]*model.Authorship, error) {
	return m.listAuthorships(qryListAuthorships)
}

//InsertAuthorship insert an authorship
func (m *Model) InsertAuthorship(authorID, bookID, roleID int) (*model.Authorship, error) {
	res, err := m.db.Exec(qryInsertAuthorship, authorID, bookID, roleID)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Authorship{ID: id, AuthorID: int64(authorID), RoleID: int64(roleID), BookID: int64(bookID)}, nil
}

//DeleteAuthorship delete an authorship
func (m *Model) DeleteAuthorship(id int) error {
	return m.exec(qryDeleteAuthorship, id)
}
