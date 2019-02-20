package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetEditorByID   = `SELECT id, name FROM editor WHERE id = ?`
	qryGetEditorByName = `SELECT id, name FROM editor WHERE name = ?`
	qryInsertEditor    = `
INSERT INTO editor (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())`
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

//ListEditorsByIDs list editors by print id or series id
func (m *Model) ListEditorsByIDs(printID, seriesID *int) ([]*model.Editor, error) {
	qry := `SELECT DISTINCT editor.id, editor.name FROM editor`
	where := []string{"1"}
	vals := []interface{}{}

	if printID != nil || seriesID != nil {
		qry += ` JOIN collection ON (collection.editor_id = editor.id) JOIN edition ON (edition.collection_id = collection.id)`
		if printID != nil {
			where = append(where, `edition.print_id = ?`)
			vals = append(vals, *printID)
		}
		if seriesID != nil {
			qry += ` JOIN book ON (edition.book_id = book.id)`
			where = append(where, `book.series_id = ?`)
			vals = append(vals, *seriesID)
		}
	}
	return m.listEditors(qry+` WHERE `+strings.Join(where, " AND "), vals...)
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
