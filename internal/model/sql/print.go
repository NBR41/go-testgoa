package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetPrintByID   = `SELECT id, name FROM print WHERE id = ?`
	qryGetPrintByName = `SELECT id, name FROM print WHERE name = ?`
	qryInsertPrint    = `
INSERT INTO print (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())`
	qryUpdatePrint = `UPDATE print SET name = ?, update_ts = NOW() WHERE id = ?`
	qryDeletePrint = `DELETE FROM print WHERE id = ?`
)

func (m *Model) getPrint(query string, params ...interface{}) (*model.Print, error) {
	var v = model.Print{}
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

// ListPrints list prints
func (m *Model) listPrints(query string, params ...interface{}) ([]*model.Print, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Print{}
	for rows.Next() {
		v := model.Print{}
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

// GetPrintByID returns print by ID
func (m *Model) GetPrintByID(id int) (*model.Print, error) {
	return m.getPrint(qryGetPrintByID, id)
}

// GetPrintByName returns print by name
func (m *Model) GetPrintByName(name string) (*model.Print, error) {
	return m.getPrint(qryGetPrintByName, name)
}

// ListPrintsByIDs list prints by collection id or editor id or series id
func (m *Model) ListPrintsByIDs(collectionID, editorID, seriesID *int) ([]*model.Print, error) {
	qry := `SELECT DISTINCT print.id, print.name FROM print`
	where := []string{"1"}
	vals := []interface{}{}
	if collectionID != nil || editorID != nil || seriesID != nil {
		qry += ` JOIN edition ON (edition.print_id = print.id)`

		if collectionID != nil {
			where = append(where, `edition.collection_id = ?`)
			vals = append(vals, *collectionID)
		}
		if editorID != nil {
			qry += ` JOIN collection ON (edition.collection_id = collection.id)`
			where = append(where, `collection.editor_id = ?`)
			vals = append(vals, *editorID)
		}
		if seriesID != nil {
			qry += ` JOIN book ON (edition.book_id = book.id)`
			where = append(where, `book.series_id = ?`)
			vals = append(vals, *seriesID)
		}
	}
	return m.listPrints(qry+` WHERE `+strings.Join(where, " AND "), vals...)
}

// InsertPrint inserts print
func (m *Model) InsertPrint(name string) (*model.Print, error) {
	res, err := m.db.Exec(qryInsertPrint, name)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Print{ID: id, Name: name}, nil
}

// UpdatePrint update print
func (m *Model) UpdatePrint(id int, name string) error {
	return m.exec(qryUpdatePrint, name, id)
}

// DeletePrint delete print
func (m *Model) DeletePrint(id int) error {
	return m.exec(qryDeletePrint, id)
}

func (m *Model) getOrInsertPrint(name string) (*model.Print, error) {
	print, err := m.GetPrintByName(name)
	if err == model.ErrNotFound {
		print, err = m.InsertPrint(name)
	}
	if err != nil {
		return nil, err
	}
	return print, nil
}
