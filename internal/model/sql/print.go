package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetPrintByID   = `SELECT id, name FROM print WHERE id = ?`
	qryGetPrintByName = `SELECT id, name FROM print WHERE name = ?`
	qryListPrints     = `SELECT id, name FROM print`
	qryInsertPrint    = `
INSERT INTO print (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
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

// ListPrints list prints
func (m *Model) ListPrints() ([]*model.Print, error) {
	return m.listPrints(qryListPrints)
}

// ListPrintsByIDs list prints by collection id or series id
func (m *Model) ListPrintsByIDs(collectionID, seriesID *int) ([]*model.Print, error) {
	if collectionID == nil && seriesID == nil {
		return m.ListPrints()
	}
	qry := `SELECT distinct p.id, p.name FROM print p JOIN edition e ON (e.print_id = p.id)`
	where := []string{}
	vals := []interface{}{}
	if collectionID != nil {
		where = append(where, `e.collection_id = ?`)
		vals = append(vals, *collectionID)
	}
	if seriesID != nil {
		qry += ` JOIN book b ON (e.book_id = b.id)`
		where = append(where, `b.series_id = ?`)
		vals = append(vals, *seriesID)
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
