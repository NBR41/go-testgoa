package local

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

// InsertBook inserts book
func (db *Local) InsertBook(isbn, name string, seriesID int) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getBookByISBN(isbn)
	if err == nil {
		return nil, model.ErrDuplicateKey
	}
	_, err = db.getBookByName(name)
	if err == nil {
		return nil, model.ErrDuplicateKey
	}
	_, err = db.getSeriesByID(seriesID)
	if err == model.ErrNotFound {
		return nil, model.ErrInvalidID
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

func (db *Local) getBookByName(name string) (*model.Book, error) {
	for i := range db.books {
		if db.books[i].Name == name {
			return db.books[i], nil
		}
	}
	return nil, model.ErrNotFound
}

// GetBookByName returns book by name
func (db *Local) GetBookByName(name string) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	return db.getBookByName(name)
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

// UpdateBook update book infos
func (db *Local) UpdateBook(id int, name *string, seriesID *int) error {
	db.Lock()
	defer db.Unlock()
	b, err := db.getBookByID(id)
	if err != nil {
		return err
	}
	if name != nil {
		_, err := db.getBookByName(*name)
		if err == nil {
			return model.ErrDuplicateKey
		}
		b.Name = *name
	}
	if seriesID != nil {
		_, err := db.getSeriesByID(*seriesID)
		if err == model.ErrNotFound {
			return model.ErrInvalidID
		}
		b.SeriesID = int64(*seriesID)
	}
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

//ListBooksByIDs list books by ids
func (db *Local) ListBooksByIDs(collectionID, editorID, printID, seriesID *int) ([]*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	bookIDs := make(map[int]struct{})

	for i := range db.editions {
		if (collectionID == nil || db.editions[i].CollectionID == int64(*collectionID)) &&
			(printID == nil || db.editions[i].PrintID == int64(*printID)) {
			if editorID != nil {
				if _, ok := db.collections[int(db.editions[i].CollectionID)]; ok {
					if db.collections[int(db.editions[i].CollectionID)].EditorID != int64(*editorID) {
						continue
					}
				}
			}
			bookIDs[int(db.editions[i].BookID)] = struct{}{}
		}
	}

	ret := []*model.Book{}
	for i := range db.books {
		if seriesID != nil && db.books[i].SeriesID != int64(*seriesID) {
			continue
		}
		if _, ok := bookIDs[int(db.books[i].ID)]; ok {
			ret = append(ret, db.books[i])
		}
	}
	return ret, nil
}
