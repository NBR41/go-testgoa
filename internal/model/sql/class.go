package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetClassByID          = `SELECT id, name FROM class WHERE id = ?`
	qryGetClassByName        = `SELECT id, name FROM class WHERE name = ?`
	qryListClasses           = `SELECT id, name FROM class`
	qryListClassesBySeriesID = `
SELECT distinct cl.id, cl.name
FROM class cl
JOIN classification c ON (c.class_id = cl.id)
WHERE c.series_id = ?`
	qryListClassesByAuthorID = `
SELECT distinct cl.id, cl.name
FROM class cl
JOIN classification c ON (c.class_id = cl.id)
JOIN series s ON (s.series_id = s.id)
JOIN book b ON (b.series_id = s.id)
JOIN authorship a (a.book_id = b.id)
WHERE a.author_id = ?`
	qryListClassesByCategoryID = `
SELECT distinct cl.id, cl.name
FROM class cl
JOIN classification c ON (c.class_id = cl.id)
JOIN series s ON (s.series_id = s.id)
WHERE s.category_id = ?`
	qryInsertClass = `
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

// ListClasses list classes
func (m *Model) ListClasses() ([]*model.Class, error) {
	return m.listClasses(qryListClasses)
}

// ListClassesBySeriesID list classes by series ID
func (m *Model) ListClassesBySeriesID(seriesID int) ([]*model.Class, error) {
	return m.listClasses(qryListClassesBySeriesID, seriesID)
}

// ListClassesByAuthorID list classes by author ID
func (m *Model) ListClassesByAuthorID(authorID int) ([]*model.Class, error) {
	return m.listClasses(qryListClassesByAuthorID, authorID)
}

// ListClassesByCategoryID list classes by category ID
func (m *Model) ListClassesByCategoryID(categoryID int) ([]*model.Class, error) {
	return m.listClasses(qryListClassesByCategoryID, categoryID)
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
