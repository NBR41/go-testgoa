package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getCollectionByID(id int) (*model.Collection, error) {
	if p, ok := db.collections[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetCollectionByID return a collection by id
func (db *Local) GetCollectionByID(id int) (*model.Collection, error) {
	db.Lock()
	defer db.Unlock()
	return db.getCollectionByID(id)
}

func (db *Local) getCollectionByName(name string) (*model.Collection, error) {
	for i := range db.collections {
		if db.collections[i].Name == name {
			return db.collections[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetCollectionByName return a collection by name
func (db *Local) GetCollectionByName(name string) (*model.Collection, error) {
	db.Lock()
	defer db.Unlock()
	return db.getCollectionByName(name)
}

//InsertCollection insert a collection and return it
func (db *Local) InsertCollection(name string, editorID int) (*model.Collection, error) {
	db.Lock()
	defer db.Unlock()
	if _, err := db.getCollectionByName(name); err == nil {
		return nil, model.ErrDuplicateKey
	}
	e, err := db.getEditorByID(editorID)
	if err != nil {
		return nil, model.ErrInvalidID
	}
	idx := len(db.collections) + 1
	v := &model.Collection{ID: int64(idx), Name: name, EditorID: int64(editorID), Editor: e}
	db.collections[idx] = v
	return v, nil
}

//UpdateCollection update a collection name or editor id
func (db *Local) UpdateCollection(id int, name *string, editorID *int) error {
	db.Lock()
	defer db.Unlock()
	c, err := db.getCollectionByID(id)
	if err != nil {
		return err
	}
	if name != nil {
		if _, err = db.getCollectionByName(*name); err == nil {
			return model.ErrDuplicateKey
		}
		c.Name = *name
	}
	if editorID != nil {
		e, err := db.getEditorByID(*editorID)
		if err == model.ErrNotFound {
			return model.ErrInvalidID
		}
		c.EditorID = e.ID
		c.Editor = e
	}
	return nil
}

//DeleteCollection delete a collection
func (db *Local) DeleteCollection(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.collections[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.collections, id)
	return nil
}

//ListCollections list all collections
func (db *Local) ListCollections() ([]*model.Collection, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.collections))
	i := 0
	for id := range db.collections {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Collection, len(ids))
	for i, id := range ids {
		list[i] = db.collections[id]
	}
	return list, nil
}

//ListCollectionsByEditorID list all collections for an editor id
func (db *Local) ListCollectionsByEditorID(id int) ([]*model.Collection, error) {
	db.Lock()
	defer db.Unlock()
	ids := []int{}
	for i := range db.collections {
		if db.collections[i].EditorID == int64(id) {
			ids = append(ids, i)
		}
	}
	sort.Ints(ids)
	list := make([]*model.Collection, len(ids))
	for i, id := range ids {
		list[i] = db.collections[id]
	}
	return list, nil
}
