package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getClassByID(id int) (*model.Class, error) {
	if p, ok := db.classes[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetClassByID return author by ID
func (db *Local) GetClassByID(id int) (*model.Class, error) {
	db.Lock()
	defer db.Unlock()
	return db.getClassByID(id)
}

func (db *Local) getClassByName(name string) (*model.Class, error) {
	for i := range db.classes {
		if db.classes[i].Name == name {
			return db.classes[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetClassByName return author by name
func (db *Local) GetClassByName(name string) (*model.Class, error) {
	db.Lock()
	defer db.Unlock()
	return db.getClassByName(name)
}

//ListClasss list classes
func (db *Local) ListClasses() ([]*model.Class, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.classes))
	i := 0
	for id := range db.classes {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Class, len(ids))
	for i, id := range ids {
		list[i] = db.classes[id]
	}
	return list, nil
}

//InsertClass insert author
func (db *Local) InsertClass(name string) (*model.Class, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getClassByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.classes) + 1
	v := &model.Class{ID: int64(idx), Name: name}
	db.classes[idx] = v
	return v, nil
}

//UpdateClass update author
func (db *Local) UpdateClass(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	v, err := db.getClassByID(id)
	if err != nil {
		return err
	}
	v.Name = name
	return nil
}

//DeleteClass delete author
func (db *Local) DeleteClass(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.classes[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.classes, id)
	return nil
}
