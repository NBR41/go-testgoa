package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getPrintByID(id int) (*model.Print, error) {
	if p, ok := db.prints[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetPrintByID return author by ID
func (db *Local) GetPrintByID(id int) (*model.Print, error) {
	db.Lock()
	defer db.Unlock()
	return db.getPrintByID(id)
}

func (db *Local) getPrintByName(name string) (*model.Print, error) {
	for i := range db.prints {
		if db.prints[i].Name == name {
			return db.prints[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetPrintByName return author by name
func (db *Local) GetPrintByName(name string) (*model.Print, error) {
	db.Lock()
	defer db.Unlock()
	return db.getPrintByName(name)
}

//ListPrints list prints
func (db *Local) ListPrints() ([]*model.Print, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.prints))
	i := 0
	for id := range db.prints {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Print, len(ids))
	for i, id := range ids {
		list[i] = db.prints[id]
	}
	return list, nil
}

//InsertPrint insert author
func (db *Local) InsertPrint(name string) (*model.Print, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getPrintByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.prints) + 1
	v := &model.Print{ID: int64(idx), Name: name}
	db.prints[idx] = v
	return v, nil
}

//UpdatePrint update author
func (db *Local) UpdatePrint(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	v, err := db.getPrintByID(id)
	if err != nil {
		return err
	}
	v.Name = name
	return nil
}

//DeletePrint delete author
func (db *Local) DeletePrint(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.prints[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.prints, id)
	return nil
}