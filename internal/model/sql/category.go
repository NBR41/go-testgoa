package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
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
	return m.getCategory(`SELECT id, name FROM category where id = ?`, id)
}

// GetCategoryByName returns category by name
func (m *Model) GetCategoryByName(name string) (*model.Category, error) {
	return m.getCategory(`SELECT id, name FROM category where name = ?`, name)
}

// ListCategories list categories
func (m *Model) ListCategories() ([]*model.Category, error) {
	return m.listCategories(`SELECT id, name FROM category`)
}

// ListCategoriesByAuthorID list categories by author id
func (m *Model) ListCategoriesByAuthorID(authorID int) ([]*model.Category, error) {
	if _, err := m.GetAuthorByID(authorID); err != nil {
		return nil, err
	}
	return m.listCategories(`
SELECT distinct c.id, c.name
FROM category c
JOIN series s ON (s.category_id = c.id)
JOIN book b ON (b.series_id = s.id)
JOIN authorship a ON (a.book_id = b.id)
WHERE a.author_id = ?`,
		authorID,
	)
}

// ListCategoriesByClassID list categories by class id
func (m *Model) ListCategoriesByClassID(classID int) ([]*model.Category, error) {
	if _, err := m.GetClassByID(classID); err != nil {
		return nil, err
	}
	return m.listCategories(`
SELECT distinct c.id, c.name
FROM category c
JOIN series s ON (s.category_id = c.id)
JOIN classification cl ON (cl.series_id = s.id)
WHERE cl.class_id = ?`,
		classID,
	)
}

// InsertCategory inserts category
func (m *Model) InsertCategory(name string) (*model.Category, error) {
	res, err := m.db.Exec(
		`
INSERT INTO category (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		name,
	)
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
	return m.exec(
		`UPDATE category SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

// DeleteCategory delete category
func (m *Model) DeleteCategory(id int) error {
	return m.exec(`DELETE FROM category where id = ?`, id)
}
