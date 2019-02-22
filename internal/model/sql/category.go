package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetCategoryByID   = `SELECT id, name FROM category WHERE id = ?`
	qryGetCategoryByName = `SELECT id, name FROM category WHERE name = ?`
	qryInsertCategory    = `
INSERT INTO category (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())`
	qryUpdateCategory = `UPDATE category SET name = ?, update_ts = NOW() WHERE id = ?`
	qryDeleteCategory = `DELETE FROM category WHERE id = ?`
)

func (m *Model) getCategory(query string, params ...interface{}) (*model.Category, error) {
	var v = model.Category{}
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

func (m *Model) listCategories(query string, params ...interface{}) ([]*model.Category, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Category{}
	for rows.Next() {
		v := model.Category{}
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

// GetCategoryByID returns category by ID
func (m *Model) GetCategoryByID(id int) (*model.Category, error) {
	return m.getCategory(qryGetCategoryByID, id)
}

// GetCategoryByName returns category by name
func (m *Model) GetCategoryByName(name string) (*model.Category, error) {
	return m.getCategory(qryGetCategoryByName, name)
}

//ListCategoriesByIDs list categories by class id or author id
func (m *Model) ListCategoriesByIDs(authorID, classID *int) ([]*model.Category, error) {
	qry := `SELECT DISTINCT category.id, category.name FROM category`
	where := []string{"1"}
	vals := []interface{}{}
	if authorID != nil || classID != nil {
		qry += ` JOIN series ON (series.category_id = category.id)`
		if authorID != nil {
			qry += ` JOIN book ON (book.series_id = series.id) JOIN authorship ON (authorship.book_id = book.id)`
			where = append(where, `authorship.author_id = ?`)
			vals = append(vals, *authorID)
		}
		if classID != nil {
			qry += ` JOIN classification ON (classification.series_id = series.id)`
			where = append(where, `classification.class_id = ?`)
			vals = append(vals, *classID)
		}
	}
	return m.listCategories(qry+` WHERE `+strings.Join(where, " AND "), vals...)
}

// InsertCategory inserts category
func (m *Model) InsertCategory(name string) (*model.Category, error) {
	res, err := m.db.Exec(qryInsertCategory, name)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Category{ID: id, Name: name}, nil
}

// UpdateCategory update category
func (m *Model) UpdateCategory(id int, name string) error {
	return m.exec(qryUpdateCategory, name, id)
}

// DeleteCategory delete category
func (m *Model) DeleteCategory(id int) error {
	return m.exec(qryDeleteCategory, id)
}

func (m *Model) getOrInsertCategory(name string) (*model.Category, error) {
	category, err := m.GetCategoryByName(name)
	if err == model.ErrNotFound {
		category, err = m.InsertCategory(name)
	}
	if err != nil {
		return nil, err
	}
	return category, nil
}
