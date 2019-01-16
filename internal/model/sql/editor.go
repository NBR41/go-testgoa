package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetEditorByID   = `SELECT id, name FROM editor WHERE id = ?`
	qryGetEditorByName = `SELECT id, name FROM editor WHERE name = ?`
	qryListEditors     = `SELECT id, name FROM editor`
	qryInsertEditor    = `
INSERT INTO editor (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
	qryUpdateEditor = `UPDATE editor SET name = ?, update_ts = NOW() WHERE id = ?`
	qryDeleteEditor = `DELETE FROM editor WHERE id = ?`
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

func (m *Model) listEditors(query string, params ...interface{}) ([]*model.Editor, error) {
	rows, err := m.db.Query(query, params...)
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

// GetEditorByID returns editor by ID
func (m *Model) GetEditorByID(id int) (*model.Editor, error) {
	return m.getEditor(qryGetEditorByID, id)
}

// GetEditorByName returns editor by name
func (m *Model) GetEditorByName(name string) (*model.Editor, error) {
	return m.getEditor(qryGetEditorByName, name)
}

// ListEditors list editors
func (m *Model) ListEditors() ([]*model.Editor, error) {
	return m.listEditors(qryListEditors)
}

// InsertEditor inserts editor
func (m *Model) InsertEditor(name string) (*model.Editor, error) {
	res, err := m.db.Exec(qryInsertEditor, name)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Editor{ID: id, Name: name}, nil
}

// UpdateEditor update editor
func (m *Model) UpdateEditor(id int, name string) error {
	return m.exec(qryUpdateEditor, name, id)
}

// DeleteEditor delete editor
func (m *Model) DeleteEditor(id int) error {
	return m.exec(qryDeleteEditor, id)
}
