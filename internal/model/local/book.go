package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

// InsertBook inserts book
func (db *Local) InsertBook(isbn, name string) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getBookByISBN(isbn)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.books) + 1
	b := &model.Book{ID: int64(idx), ISBN: isbn, Name: name}
	db.books[idx] = b
	return b, nil
}

// GetBookByID returns book by ID
func (db *Local) GetBookByID(id int) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	return db.getBookByID(id)
}

func (db *Local) getBookByID(id int) (*model.Book, error) {
	if p, ok := db.books[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

// GetBookByName returns book by name
func (db *Local) GetBookByName(name string) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	for i := range db.books {
		if db.books[i].Name == name {
			return db.books[i], nil
		}
	}
	return nil, model.ErrNotFound
}

// GetBookByISBN returns book by isbn
func (db *Local) GetBookByISBN(isbn string) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	return db.getBookByISBN(isbn)
}

func (db *Local) getBookByISBN(isbn string) (*model.Book, error) {
	for i := range db.books {
		if db.books[i].ISBN == isbn {
			return db.books[i], nil
		}
	}
	return nil, model.ErrNotFound
}

// GetBookList returns book list
func (db *Local) GetBookList() ([]model.Book, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.books))
	i := 0
	for id := range db.books {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]model.Book, len(ids))
	for i, id := range ids {
		list[i] = *db.books[id]
	}
	return list, nil
}

// UpdateBook update book infos
func (db *Local) UpdateBook(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	b, err := db.getBookByID(id)
	if err != nil {
		return err
	}
	b.Name = name
	return nil

}

// DeleteBook delete book by ID
func (db *Local) DeleteBook(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.books[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.books, id)
	return nil
}
