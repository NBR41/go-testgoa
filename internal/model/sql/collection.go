package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetCollectionByID = `
SELECT c.id, c.name, e.id, e.name
FROM collection c
JOIN editors e ON (e.id = c.editor_id)
WHERE c.id = ?`
	qryGetCollectionByName = `
SELECT c.id, c.name, e.id, e.name
FROM collection c
JOIN editors e ON (e.id = c.editor_id)
WHERE c.name = ?`
	qryListCollections           = `SELECT distinct c.id, c.name, e.id, e.name FROM collection c JOIN editors e ON (e.id = c.editor_id)`
	qryListCollectionsByEditorID = `SELECT distinct c.id, c.name, e.id, e.name FROM collection c JOIN editors e ON (e.id = c.editor_id) WHERE e.id = ?`
	qryInsertCollection          = `
INSERT INTO collection (id, name, editor_id, create_ts, update_ts)
VALUES (null, ?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
	qryDeleteCollection = `DELETE FROM collection WHERE id = ?`
)

func (m *Model) getCollection(query string, params ...interface{}) (*model.Collection, error) {
	var a = &model.Collection{Editor: &model.Editor{}}
	err := m.db.QueryRow(query, params...).Scan(&a.ID, &a.Name, &a.Editor.ID, &a.Editor.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		a.EditorID = a.Editor.ID
		return a, nil
	}
}

func (m *Model) listCollections(query string, params ...interface{}) ([]*model.Collection, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Collection{}
	for rows.Next() {
		v := &model.Collection{Editor: &model.Editor{}}
		if err := rows.Scan(&v.ID, &v.Name, &v.Editor.ID, &v.Editor.Name); err != nil {
			return nil, err
		}
		v.EditorID = v.Editor.ID
		l = append(l, v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

//GetCollectionByID return a collection by id
func (m *Model) GetCollectionByID(id int) (*model.Collection, error) {
	return m.getCollection(qryGetCollectionByID, id)
}

//GetCollectionByName return a collection by name
func (m *Model) GetCollectionByName(name string) (*model.Collection, error) {
	return m.getCollection(qryGetCollectionByName, name)
}

//ListCollections list all collections
func (m *Model) ListCollections() ([]*model.Collection, error) {
	return m.listCollections(qryListCollections)
}

//ListCollectionsByEditorID list all collections for an editor id
func (m *Model) ListCollectionsByEditorID(id int) ([]*model.Collection, error) {
	return m.listCollections(qryListCollectionsByEditorID, id)
}

//InsertCollection insert a collection and return it
func (m *Model) InsertCollection(name string, editorID int) (*model.Collection, error) {
	res, err := m.db.Exec(qryInsertCollection, name, editorID)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Collection{ID: id, Name: name, EditorID: int64(editorID)}, nil
}

//UpdateCollection update a collection name or editor id
func (m *Model) UpdateCollection(id int, name *string, editorID *int) error {
	if name == nil && editorID == nil {
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
	if editorID != nil {
		parts += "editor_id = ?, "
		values = append(values, *editorID)
	}
	values = append(values, id)
	return m.exec(
		`UPDATE collection SET `+parts+`update_ts = NOW() WHERE id = ?`,
		values...,
	)
}

//DeleteCollection delete a collection
func (m *Model) DeleteCollection(id int) error {
	return m.exec(qryDeleteCollection, id)
}
