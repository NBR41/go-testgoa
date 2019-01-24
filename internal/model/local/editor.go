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

//ListEditorsByIDs returns a filtered editor list
func (db *Local) ListEditorsByIDs(printID, seriesID *int) ([]*model.Editor, error) {
	db.Lock()
	defer db.Unlock()

	if printID == nil && seriesID == nil {
		ret := []*model.Editor{}
		for k := range db.editors {
			ret = append(ret, db.editors[k])
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

	var collectionIDs = make(map[int]struct{})
	for i := range db.editions {
		if bookIDs != nil {
			if _, ok := bookIDs[int(db.editions[i].BookID)]; !ok {
				continue
			}
		}
		if printID == nil || int64(*printID) == db.editions[i].PrintID {
			collectionIDs[int(db.editions[i].CollectionID)] = struct{}{}
		}
	}

	var editorIDs = make(map[int]struct{})
	for i := range db.collections {
		if _, ok := collectionIDs[i]; ok {
			editorIDs[int(db.collections[i].EditorID)] = struct{}{}
		}
	}

	ret := []*model.Editor{}
	for k := range editorIDs {
		if _, ok := db.editors[k]; ok {
			ret = append(ret, db.editors[k])
		}
	}
	return ret, nil
}
