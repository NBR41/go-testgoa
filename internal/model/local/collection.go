package local

import (
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

//ListCollectionsByIDs returns filtered collection list
func (db *Local) ListCollectionsByIDs(editorID, printID, seriesID *int) ([]*model.Collection, error) {
	db.Lock()
	defer db.Unlock()

	if editorID == nil && printID == nil && seriesID == nil {
		ret := []*model.Collection{}
		for k := range db.collections {
			ret = append(ret, db.collections[k])
		}
		return ret, nil
	}

	var bookIDs map[int]struct{}
	if seriesID != nil {
		bookIDs = make(map[int]struct{})
		for i := range db.books {
			if db.books[i].SeriesID == int64(*seriesID) {
				bookIDs[i] = struct{}{}
			}
		}
	}

	var collectionIDs map[int]struct{}
	if seriesID != nil || printID != nil {
		collectionIDs = make(map[int]struct{})
		for i := range db.editions {
			if bookIDs != nil {
				if _, ok := bookIDs[int(db.editions[i].BookID)]; !ok {
					continue
				}
			}
			if printID == nil || db.editions[i].PrintID == int64(*printID) {
				collectionIDs[int(db.editions[i].CollectionID)] = struct{}{}
			}
		}
	}

	var finalCollectionIDs = make(map[int]struct{})
	for i := range db.collections {
		if collectionIDs != nil {
			if _, ok := collectionIDs[i]; !ok {
				continue
			}
		}
		if editorID == nil || db.collections[i].EditorID == int64(*editorID) {
			finalCollectionIDs[i] = struct{}{}
		}
	}

	ret := []*model.Collection{}
	for k := range finalCollectionIDs {
		if _, ok := db.collections[k]; ok {
			ret = append(ret, db.collections[k])
		}
	}
	return ret, nil
}
