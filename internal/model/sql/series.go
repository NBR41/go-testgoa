package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
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
	return m.getSeries(
		`
SELECT s.id, s.name, c.id, c.name
FROM series s
JOIN category c ON (c.id = s.category_id)
WHERE s.id = ?`, id,
	)
}

//GetSeriesByName return a series by name
func (m *Model) GetSeriesByName(name string) (*model.Series, error) {
	return m.getSeries(
		`
SELECT s.id, s.name, c.id, c.name
FROM series s
JOIN category c ON (c.id = s.category_id)
WHERE s.name = ?`, name,
	)
}

//ListSeries list all series
func (m *Model) ListSeries() ([]*model.Series, error) {
	return m.listSeries(
		`SELECT s.id, s.name, c.id, c.name FROM series s JOIN category c ON (c.id = s.series_id)`,
	)
}

//InsertSeries insert a series and return it
func (m *Model) InsertSeries(name string, categoryID int) (*model.Series, error) {
	e, err := m.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}
	res, err := m.db.Exec(
		`
INSERT INTO series (id, name, category_id, create_ts, update_ts)
VALUES (null, ?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		name, categoryID,
	)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Series{ID: id, Name: name, CategoryID: int64(categoryID), Category: e}, nil
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
		if _, err := m.GetCategoryByID(*categoryID); err != nil {
			return err
		}
		parts += "category_id = ?, "
		values = append(values, *categoryID)
	}
	values = append(values, id)
	return m.exec(
		`UPDATE series SET `+parts+`update_ts = NOW() where id = ?`,
		values...,
	)
}

//DeleteSeries delete a series
func (m *Model) DeleteSeries(id int) error {
	return m.exec(`DELETE FROM series where id = ?`, id)
}
