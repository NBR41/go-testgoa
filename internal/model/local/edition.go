package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getEditionByID(id int) (*model.Edition, error) {
	if p, ok := db.editions[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetEditionByID returns edition by id
func (db *Local) GetEditionByID(id int) (*model.Edition, error) {
	db.Lock()
	defer db.Unlock()
	return db.getEditionByID(id)
}

//ListEditions lsit editions
func (db *Local) ListEditions() ([]*model.Edition, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.editions))
	i := 0
	for id := range db.editions {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Edition, len(ids))
	for i, id := range ids {
		list[i] = db.editions[id]
	}
	return list, nil
}

//InsertEdition insert an edition
func (db *Local) InsertEdition(bookID, collectionID, printID int) (*model.Edition, error) {
	db.Lock()
	defer db.Unlock()
	for _, v := range db.editions {
		if v.BookID == int64(bookID) && v.CollectionID == int64(collectionID) && v.PrintID == int64(printID) {
			return nil, model.ErrDuplicateKey
		}
	}

	if _, ok := db.books[bookID]; !ok {
		return nil, model.ErrInvalidID
	}
	if _, ok := db.collections[collectionID]; !ok {
		return nil, model.ErrInvalidID
	}
	if _, ok := db.prints[printID]; !ok {
		return nil, model.ErrInvalidID
	}

	idx := len(db.editions) + 1
	v := &model.Edition{ID: int64(idx), BookID: int64(bookID), CollectionID: int64(collectionID), PrintID: int64(printID)}
	db.editions[idx] = v
	return v, nil
}

//DeleteEdition delete an edition
func (db *Local) DeleteEdition(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.editions[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.editions, id)
	return nil
}
