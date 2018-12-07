package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getEditionTypeByID(id int) (*model.EditionType, error) {
	if p, ok := db.editionTypes[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetEditionTypeByID return author by ID
func (db *Local) GetEditionTypeByID(id int) (*model.EditionType, error) {
	db.Lock()
	defer db.Unlock()
	return db.getEditionTypeByID(id)
}

func (db *Local) getEditionTypeByName(name string) (*model.EditionType, error) {
	for i := range db.editionTypes {
		if db.editionTypes[i].Name == name {
			return db.editionTypes[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetEditionTypeByName return author by name
func (db *Local) GetEditionTypeByName(name string) (*model.EditionType, error) {
	db.Lock()
	defer db.Unlock()
	return db.getEditionTypeByName(name)
}

//GetEditionTypeList list editionTypes
func (db *Local) GetEditionTypeList() ([]*model.EditionType, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.editionTypes))
	i := 0
	for id := range db.editionTypes {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.EditionType, len(ids))
	for i, id := range ids {
		list[i] = db.editionTypes[id]
	}
	return list, nil
}

//InsertEditionType insert author
func (db *Local) InsertEditionType(name string) (*model.EditionType, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getEditionTypeByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.editionTypes) + 1
	v := &model.EditionType{ID: int64(idx), Name: name}
	db.editionTypes[idx] = v
	return v, nil
}

//UpdateEditionType update author
func (db *Local) UpdateEditionType(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	v, err := db.getEditionTypeByID(id)
	if err != nil {
		return err
	}
	v.Name = name
	return nil
}

//DeleteEditionType delete author
func (db *Local) DeleteEditionType(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.editionTypes[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.editionTypes, id)
	return nil
}
