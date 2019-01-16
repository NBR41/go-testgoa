package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetSeriesByID = `
	SELECT s.id, s.name, c.id, c.name
	FROM series s
	JOIN category c ON (c.id = s.category_id)
	WHERE s.id = ?`
	qryGetSeriesByName = `
SELECT s.id, s.name, c.id, c.name
FROM series s
JOIN category c ON (c.id = s.category_id)
WHERE s.name = ?`
	qryListSeries               = `SELECT s.id, s.name, c.id, c.name FROM series s JOIN category c ON (c.id = s.series_id)`
	qryListSeriesByCollectionID = `
SELECT DISTINCT s.id, s.name, c.id, c.name
FROM series s
JOIN category c ON (c.id = s.series_id)
JOIN book b ON (b.seried_id = s.id)
JOIN edition e ON (e.book_id = b.id)
WHERE e.collection_id = ?`
	qryListSeriesByCollectionIDPrintID = `
SELECT DISTINCT s.id, s.name, c.id, c.name
FROM series s
JOIN category c ON (c.id = s.series_id)
JOIN book b ON (b.seried_id = s.id)
JOIN edition e ON (e.book_id = b.id)
WHERE e.collection_id = ? AND e.print_id = ?`
	qryInsertSeries = `
INSERT INTO series (id, name, category_id, create_ts, update_ts)
VALUES (null, ?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
	qryDeleteSeries = `DELETE FROM series WHERE id = ?`
)

func (m *Model) getSeries(query string, params ...interface{}) (*model.Series, error) {
	var a = &model.Series{Category: &model.Category{}}
	err := m.db.QueryRow(query, params...).Scan(&a.ID, &a.Name, &a.Category.ID, &a.Category.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		a.CategoryID = a.Category.ID
		return a, nil
	}
}

func (m *Model) listSeries(query string, params ...interface{}) ([]*model.Series, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Series{}
	for rows.Next() {
		v := &model.Series{Category: &model.Category{}}
		if err := rows.Scan(&v.ID, &v.Name, &v.Category.ID, &v.Category.Name); err != nil {
			return nil, err
		}
		v.CategoryID = v.Category.ID
		l = append(l, v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

//GetSeriesByID return a series by id
func (m *Model) GetSeriesByID(id int) (*model.Series, error) {
	return m.getSeries(qryGetSeriesByID, id)
}

//GetSeriesByName return a series by name
func (m *Model) GetSeriesByName(name string) (*model.Series, error) {
	return m.getSeries(qryGetSeriesByName, name)
}

//ListSeries list all series
func (m *Model) ListSeries() ([]*model.Series, error) {
	return m.listSeries(qryListSeries)
}

//ListSeriesByIDs list series by filter author ID, role ID, category ID, class ID
func (m *Model) ListSeriesByIDs(authorID, roleID, categoryID, classID *int) ([]*model.Series, error) {
	if authorID == nil && roleID == nil && categoryID == nil && classID == nil {
		return m.ListSeries()
	}
	qry := `SELECT DISTINCT s.id, s.name, c.id, c.name FROM series s JOIN category c ON (c.id = s.series_id)`
	where := []string{}
	vals := []interface{}{}
	if authorID != nil || roleID != nil {
		qry += ` JOIN book b ON (b.series_id = s.id) JOIN authorship a ON (a.book_id = b.id)`
		if authorID != nil {
			where = append(where, `a.author_id = ?`)
			vals = append(vals, *authorID)
		}
		if roleID != nil {
			where = append(where, `a.role_id = ?`)
			vals = append(vals, *roleID)
		}
	}
	if classID != nil {
		qry += ` JOIN classification cl ON (cl.series_id = s.id)`
		where = append(where, `cl.class_id = ?`)
		vals = append(vals, *classID)
	}
	if categoryID != nil {
		where = append(where, `c.id = ?`)
		vals = append(vals, *categoryID)
	}
	return m.listSeries(qry+` WHERE `+strings.Join(where, " AND "), vals...)
}

//ListSeriesByCollectionID list series by collection id
func (m *Model) ListSeriesByCollectionID(collectionID int) ([]*model.Series, error) {
	return m.listSeries(qryListSeriesByCollectionID, collectionID)
}

//ListSeriesByCollectionIDPrintID list series by couple collection id print id
func (m *Model) ListSeriesByCollectionIDPrintID(collectionID, printID int) ([]*model.Series, error) {
	return m.listSeries(qryListSeriesByCollectionIDPrintID, collectionID, printID)
}

//InsertSeries insert a series and return it
func (m *Model) InsertSeries(name string, categoryID int) (*model.Series, error) {
	res, err := m.db.Exec(qryInsertSeries, name, categoryID)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Series{ID: id, Name: name, CategoryID: int64(categoryID)}, nil
}

//UpdateSeries update a series name or category id
func (m *Model) UpdateSeries(id int, name *string, categoryID *int) error {
	if name == nil && categoryID == nil {
		return nil
	}
	var (
		values []interface{}
		parts  string
	)
	if name != nil {
		parts += "name = ?, "
		values = append(values, *name)
	}
	if categoryID != nil {
		parts += "category_id = ?, "
		values = append(values, *categoryID)
	}
	values = append(values, id)
	return m.exec(
		`UPDATE series SET `+parts+`update_ts = NOW() WHERE id = ?`,
		values...,
	)
}

//DeleteSeries delete a series
func (m *Model) DeleteSeries(id int) error {
	return m.exec(qryDeleteSeries, id)
}
