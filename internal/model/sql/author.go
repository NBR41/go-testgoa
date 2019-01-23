package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetAuthorByID   = `SELECT id, name FROM author WHERE id = ?`
	qryGetAuthorByName = `SELECT id, name FROM author WHERE name = ?`
	qryInsertAuthor    = `
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

//ListAuthorsByIDs returns filtered author list
func (m *Model) ListAuthorsByIDs(categoryID, roleID *int) ([]*model.Author, error) {
	qry := `SELECT DISTINCT author.id, author.name FROM author`
	where := []string{"1"}
	vals := []interface{}{}
	if categoryID != nil || roleID != nil {
		qry += ` JOIN authorship ON (authorship.author_id = author.id)`
		if categoryID != nil {
			qry += ` JOIN book ON (authorship.book_id = book.id) JOIN series ON (book.series_id = series.id)`
			where = append(where, `series.category_id = ?`)
			vals = append(vals, *categoryID)
		}
		if roleID != nil {
			where = append(where, `authorship.role_id = ?`)
			vals = append(vals, *roleID)
		}
	}
	return m.listAuthors(qry+` WHERE `+strings.Join(where, " AND "), vals...)
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
