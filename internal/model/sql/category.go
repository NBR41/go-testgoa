package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetCategoryByID          = `SELECT id, name FROM category WHERE id = ?`
	qryGetCategoryByName        = `SELECT id, name FROM category WHERE name = ?`
	qryListCategories           = `SELECT id, name FROM category`
	qryListCategoriesByAuthorID = `
SELECT distinct c.id, c.name
FROM category c
JOIN series s ON (s.category_id = c.id)
JOIN book b ON (b.series_id = s.id)
JOIN authorship a ON (a.book_id = b.id)
WHERE a.author_id = ?`
	qryListCategoriesByClassID = `
SELECT distinct c.id, c.name
FROM category c
JOIN series s ON (s.category_id = c.id)
JOIN classification cl ON (cl.series_id = s.id)
WHERE cl.class_id = ?`
	qryInsertCategory = `
INSERT INTO category (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
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

// ListCategories list categories
func (m *Model) ListCategories() ([]*model.Category, error) {
	return m.listCategories(qryListCategories)
}

// ListCategoriesByAuthorID list categories by author id
func (m *Model) ListCategoriesByAuthorID(authorID int) ([]*model.Category, error) {
	return m.listCategories(qryListCategoriesByAuthorID, authorID)
}

// ListCategoriesByClassID list categories by class id
func (m *Model) ListCategoriesByClassID(classID int) ([]*model.Category, error) {
	return m.listCategories(qryListCategoriesByClassID, classID)
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
