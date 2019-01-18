package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getEditorByID(id int) (*model.Editor, error) {
	if p, ok := db.editors[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetEditorByID return author by ID
func (db *Local) GetEditorByID(id int) (*model.Editor, error) {
	db.Lock()
	defer db.Unlock()
	return db.getEditorByID(id)
}

func (db *Local) getEditorByName(name string) (*model.Editor, error) {
	for i := range db.editors {
		if db.editors[i].Name == name {
			return db.editors[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetEditorByName return author by name
func (db *Local) GetEditorByName(name string) (*model.Editor, error) {
	db.Lock()
	defer db.Unlock()
	return db.getEditorByName(name)
}

//ListEditors list editors
func (db *Local) ListEditors() ([]*model.Editor, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.editors))
	i := 0
	for id := range db.editors {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Editor, len(ids))
	for i, id := range ids {
		list[i] = db.editors[id]
	}
	return list, nil
}

//InsertEditor insert author
func (db *Local) InsertEditor(name string) (*model.Editor, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getEditorByName(name)
	if err == nil {
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.editors) + 1
	v := &model.Editor{ID: int64(idx), Name: name}
	db.editors[idx] = v
	return v, nil
}

//UpdateEditor update author
func (db *Local) UpdateEditor(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	v, err := db.getEditorByID(id)
	if err != nil {
		return err
	}
	_, err = db.getEditorByName(name)
	if err == nil {
		return model.ErrDuplicateKey
	}
	v.Name = name
	return nil
}

//DeleteEditor delete author
func (db *Local) DeleteEditor(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.editors[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.editors, id)
	return nil
}
