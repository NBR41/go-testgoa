package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getEditor(query string, params ...interface{}) (*model.Editor, error) {
	var v = model.Editor{}
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

func (m *Model) GetEditorByID(id int) (*model.Editor, error) {
	return m.getEditor(`SELECT id, name FROM editor where id = ?`, id)
}

func (m *Model) GetEditorByName(name string) (*model.Editor, error) {
	return m.getEditor(`SELECT id, name FROM editor where name = ?`, name)
}

func (m *Model) ListEditors() ([]*model.Editor, error) {
	rows, err := m.db.Query(`SELECT id, name FROM editor`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Editor{}
	for rows.Next() {
		v := model.Editor{}
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

func (m *Model) InsertEditor(name string) (*model.Editor, error) {
	_, err := m.GetEditorByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO editor (id, name, create_ts, update_ts)
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
	return &model.Editor{ID: id, Name: name}, nil
}

func (m *Model) UpdateEditor(id int, name string) error {
	return m.exec(
		`UPDATE editor SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

func (m *Model) DeleteEditor(id int) error {
	return m.exec(`DELETE FROM editor where id = ?`, id)
}
