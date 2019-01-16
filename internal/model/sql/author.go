package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetAuthorByID           = `SELECT id, name FROM author WHERE id = ?`
	qryGetAuthorByName         = `SELECT id, name FROM author WHERE name = ?`
	qryListAuthors             = `SELECT id, name FROM author`
	qryListAuthorsByCategoryID = `
SELECT distinct a.id, a.name
FROM author a
JOIN authorship au ON (a.id = au.author_id)
JOIN book b ON (au.book_id = b.id)
JOIN series s ON (b.series_id = s.id)
JOIN category c ON (s.category_id = c.id)
WHERE c.id = ?`
	qryListAuthorsByRoleID = `
SELECT distinct a.id, a.name
FROM author a
JOIN authorship au ON (a.id = au.author_id)
JOIN role r ON (au.role_id = r.id)
WHERE r.id = ?`
	qryInsertAuthor = `
INSERT INTO author (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
	qryUpdateAuthor = `UPDATE author SET name = ?, update_ts = NOW() WHERE id = ?`
	qryDeleteAuthor = `DELETE FROM author WHERE id = ?`
)

func (m *Model) getAuthor(query string, params ...interface{}) (*model.Author, error) {
	var v = model.Author{}
	err := m.db.QueryRow(query, params...).Scan(&v.ID, &v.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &v, nil
	}
}

func (m *Model) listAuthors(query string, params ...interface{}) ([]*model.Author, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Author{}
	for rows.Next() {
		v := model.Author{}
		if err := rows.Scan(&v.ID, &v.Name); err != nil {
			return nil, err
		}
		l = append(l, &v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// GetAuthorByID returns author by ID
func (m *Model) GetAuthorByID(id int) (*model.Author, error) {
	return m.getAuthor(qryGetAuthorByID, id)
}

// GetAuthorByName returns author by name
func (m *Model) GetAuthorByName(name string) (*model.Author, error) {
	return m.getAuthor(qryGetAuthorByName, name)
}

// ListAuthors returns author list
func (m *Model) ListAuthors() ([]*model.Author, error) {
	return m.listAuthors(qryListAuthors)
}

// ListAuthorsByCategoryID returns author list by category id
func (m *Model) ListAuthorsByCategoryID(categoryID int) ([]*model.Author, error) {
	return m.listAuthors(qryListAuthorsByCategoryID, categoryID)
}

// ListAuthorsByRoleID returns author list by role id
func (m *Model) ListAuthorsByRoleID(roleID int) ([]*model.Author, error) {
	return m.listAuthors(qryListAuthorsByRoleID, roleID)
}

// InsertAuthor inserts author
func (m *Model) InsertAuthor(name string) (*model.Author, error) {
	res, err := m.db.Exec(qryInsertAuthor, name)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Author{ID: id, Name: name}, nil
}

// UpdateAuthor update author
func (m *Model) UpdateAuthor(id int, name string) error {
	return m.exec(qryUpdateAuthor, name, id)
}

// DeleteAuthor delete author
func (m *Model) DeleteAuthor(id int) error {
	return m.exec(qryDeleteAuthor, id)
}
