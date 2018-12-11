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

func (m *Model) GetCategoryByID(id int) (*model.Category, error) {
	return m.getCategory(`SELECT id, name FROM category where id = ?`, id)
}

func (m *Model) GetCategoryByName(name string) (*model.Category, error) {
	return m.getCategory(`SELECT id, name FROM category where name = ?`, name)
}

func (m *Model) ListCategories() ([]*model.Category, error) {
	rows, err := m.db.Query(`SELECT id, name FROM category`)
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

func (m *Model) InsertCategory(name string) (*model.Category, error) {
	_, err := m.GetCategoryByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO category (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		name,
	)
	if err != nil {
		return nil, err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Category{ID: id, Name: name}, nil
}

func (m *Model) UpdateCategory(id int, name string) error {
	return m.exec(
		`UPDATE category SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

func (m *Model) DeleteCategory(id int) error {
	return m.exec(`DELETE FROM category where id = ?`, id)
}
