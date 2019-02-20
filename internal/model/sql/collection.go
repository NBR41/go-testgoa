package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetCollectionByID = `
SELECT collection.id, collection.name, editor.id, editor.name
FROM collection
JOIN editor ON (editor.id = collection.editor_id)
WHERE collection.id = ?`
	qryGetCollectionByName = `
SELECT collection.id, collection.name, editor.id, editor.name
FROM collection
JOIN editor ON (editor.id = collection.editor_id)
WHERE collection.name = ?`
	qryInsertCollection = `
INSERT INTO collection (id, name, editor_id, create_ts, update_ts)
VALUES (null, ?, ?, NOW(), NOW())`
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

//ListCollectionsByIDs list classes by editor id or print id or series id
func (m *Model) ListCollectionsByIDs(editorID, printID, seriesID *int) ([]*model.Collection, error) {
	qry := `SELECT DISTINCT collection.id, collection.name, editor.id, editor.name FROM collection JOIN editor ON (collection.editor_id = editor.id)`
	where := []string{"1"}
	vals := []interface{}{}
	if editorID != nil {
		where = append(where, `editor.id = ?`)
		vals = append(vals, *editorID)
	}

	if printID != nil || seriesID != nil {
		qry += ` JOIN edition ON (edition.collection_id = collection.id)`
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
	return m.listCollections(qry+` WHERE `+strings.Join(where, " AND "), vals...)
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
