package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetClassByID   = `SELECT id, name FROM class WHERE id = ?`
	qryGetClassByName = `SELECT id, name FROM class WHERE name = ?`
	qryInsertClass    = `
INSERT INTO class (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
	qryUpdateClass = `UPDATE class SET name = ?, update_ts = NOW() WHERE id = ?`
	qryDeleteClass = `DELETE FROM class WHERE id = ?`
)

func (m *Model) getClass(query string, params ...interface{}) (*model.Class, error) {
	var v = model.Class{}
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

func (m *Model) listClasses(query string, params ...interface{}) ([]*model.Class, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Class{}
	for rows.Next() {
		v := model.Class{}
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

// GetClassByID returns class by ID
func (m *Model) GetClassByID(id int) (*model.Class, error) {
	return m.getClass(qryGetClassByID, id)
}

// GetClassByName returns class by name
func (m *Model) GetClassByName(name string) (*model.Class, error) {
	return m.getClass(qryGetClassByName, name)
}

//ListClassesByIDs list classes by author id or category id or series id
func (m *Model) ListClassesByIDs(authorID, categoryID, seriesID *int) ([]*model.Class, error) {
	qry := `SELECT DISTINCT class.id, class.name from class`
	where := []string{"1"}
	vals := []interface{}{}

	if authorID != nil || categoryID != nil || seriesID != nil {
		qry += ` JOIN classification ON (classification.class_id = class.id)`
		if authorID != nil || categoryID != nil {
			qry += ` JOIN series ON (classification.series_id = series.id)`
			if authorID != nil {
				qry += ` JOIN book ON (book.series_id = series.id) JOIN authorship ON (authorship.book_id = book.id)`
				where = append(where, `authorship.author_id = ?`)
				vals = append(vals, *authorID)
			}
			if categoryID != nil {
				where = append(where, `series.category_id = ?`)
				vals = append(vals, *categoryID)
			}
		}
		if seriesID != nil {
			where = append(where, `classification.series_id = ?`)
			vals = append(vals, *seriesID)
		}
	}
	return m.listClasses(qry+` WHERE `+strings.Join(where, " AND "), vals...)
}

// InsertClass inserts class
func (m *Model) InsertClass(name string) (*model.Class, error) {
	res, err := m.db.Exec(qryInsertClass, name)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Class{ID: id, Name: name}, nil
}

// UpdateClass update class
func (m *Model) UpdateClass(id int, name string) error {
	return m.exec(qryUpdateClass, name, id)
}

// DeleteClass delete class
func (m *Model) DeleteClass(id int) error {
	return m.exec(qryDeleteClass, id)
}
