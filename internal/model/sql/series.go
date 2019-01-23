package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetSeriesByID = `
SELECT series.id, series.name, category.id, category.name
FROM series
JOIN category ON (series.category_id = category.id)
WHERE series.id = ?`
	qryGetSeriesByName = `
SELECT series.id, series.name, category.id, category.name
FROM series
JOIN category ON (series.category_id = category.id)
WHERE series.name = ?`
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

//ListSeriesByIDs list series by filter author ID, role ID, category ID, class ID
func (m *Model) ListSeriesByIDs(authorID, categoryID, classID, roleID *int) ([]*model.Series, error) {
	qry := `SELECT DISTINCT series.id, series.name, category.id, category.name FROM series JOIN category ON (series.category_id = category.id)`
	where := []string{"1"}
	vals := []interface{}{}
	if authorID != nil || roleID != nil {
		qry += ` JOIN book ON (book.series_id = series.id) JOIN authorship ON (authorship.book_id = book.id)`
		if authorID != nil {
			where = append(where, `authorship.author_id = ?`)
			vals = append(vals, *authorID)
		}
		if roleID != nil {
			where = append(where, `authorship.role_id = ?`)
			vals = append(vals, *roleID)
		}
	}
	if classID != nil {
		qry += ` JOIN classification ON (classification.series_id = series.id)`
		where = append(where, `classification.class_id = ?`)
		vals = append(vals, *classID)
	}
	if categoryID != nil {
		where = append(where, `category.id = ?`)
		vals = append(vals, *categoryID)
	}
	return m.listSeries(qry+` WHERE `+strings.Join(where, " AND "), vals...)
}

// ListSeriesByEditionIDs list series by filter collection ID, editor ID, print ID
func (m *Model) ListSeriesByEditionIDs(collectionID, editorID, printID *int) ([]*model.Series, error) {
	qry := `SELECT DISTINCT series.id, series.name, category.id, category.name FROM series JOIN category ON (series.category_id = category.id)`
	where := []string{"1"}
	vals := []interface{}{}
	if collectionID != nil || editorID != nil || printID != nil {
		qry += ` JOIN book ON (book.series_id = series.id) JOIN edition ON (edition.book_id = book.id)`
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
	return m.listSeries(qry+` WHERE `+strings.Join(where, " AND "), vals...)
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
