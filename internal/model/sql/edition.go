package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetEditionByID = `
SELECT edition.id, book.id, book.name, collection.id, collection.name, print.id, print.name
FROM edition
JOIN book ON (edition.book_id = book.id)
JOIN collection ON (edition.collection_id = collection.id)
JOIN print ON (edition.print_id = print.id)
WHERE edition.id = ?`
	qryListEditions = `
SELECT edition.id, book.id, book.name, collection.id, collection.name, print.id, print.name
FROM edition
JOIN book ON (edition.book_id = book.id)
JOIN collection ON (edition.collection_id = collection.id)
JOIN print ON (edition.print_id = print.id)`
	qryInsertEdition = `
INSERT INTO edition (id, book_id, collection_id, print_id, create_ts, update_ts)
VALUES (null, ?, ?, ?, NOW(), NOW())`
	qryDeleteEdition = `DELETE FROM edition WHERE id = ?`
)

func (m *Model) getEdition(query string, params ...interface{}) (*model.Edition, error) {
	var v = model.Edition{Book: &model.Book{}, Collection: &model.Collection{}, Print: &model.Print{}}
	err := m.db.QueryRow(query, params...).Scan(&v.ID, &v.Book.ID, &v.Book.Name, &v.Collection.ID, &v.Collection.Name, &v.Print.ID, &v.Print.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		v.BookID = v.Book.ID
		v.CollectionID = v.Collection.ID
		v.PrintID = v.Print.ID
		return &v, nil
	}
}

func (m *Model) listEditions(query string, params ...interface{}) ([]*model.Edition, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Edition{}
	for rows.Next() {
		v := &model.Edition{Book: &model.Book{}, Collection: &model.Collection{}, Print: &model.Print{}}
		if err := rows.Scan(&v.ID, &v.Book.ID, &v.Book.Name, &v.Collection.ID, &v.Collection.Name, &v.Print.ID, &v.Print.Name); err != nil {
			return nil, err
		}
		v.BookID = v.Book.ID
		v.CollectionID = v.Collection.ID
		v.PrintID = v.Print.ID
		l = append(l, v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

//GetEditionByID returns an Edition by id
func (m *Model) GetEditionByID(id int) (*model.Edition, error) {
	return m.getEdition(qryGetEditionByID, id)
}

//ListEditions list all Editions
func (m *Model) ListEditions() ([]*model.Edition, error) {
	return m.listEditions(qryListEditions)
}

//InsertEdition insert an Edition
func (m *Model) InsertEdition(bookID, collectionID, printID int) (*model.Edition, error) {
	res, err := m.db.Exec(qryInsertEdition, bookID, collectionID, printID)
	if err != nil {
		return nil, filterError(err)
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.Edition{ID: id, BookID: int64(bookID), CollectionID: int64(collectionID), PrintID: int64(printID)}, nil
}

//DeleteEdition delete an Edition
func (m *Model) DeleteEdition(id int) error {
	return m.exec(qryDeleteEdition, id)
}
