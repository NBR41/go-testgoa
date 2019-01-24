package local

import (
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

//InsertPrint insert author
func (db *Local) InsertPrint(name string) (*model.Print, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getPrintByName(name)
	if err == nil {
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
	_, err = db.getPrintByName(name)
	if err == nil {
		return model.ErrDuplicateKey
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

//ListPrintsByIDs list prints by ids
func (db *Local) ListPrintsByIDs(collectionID, editorID, seriesID *int) ([]*model.Print, error) {
	db.Lock()
	defer db.Unlock()

	if collectionID == nil && editorID == nil && seriesID == nil {
		ret := []*model.Print{}
		for k := range db.prints {
			ret = append(ret, db.prints[k])
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
	if editorID != nil {
		collectionIDs = make(map[int]struct{})
		for i := range db.collections {
			if db.collections[i].EditorID == int64(*editorID) {
				collectionIDs[i] = struct{}{}
			}
		}
	}

	printIDs := make(map[int]struct{})
	for i := range db.editions {
		if bookIDs != nil {
			if _, ok := bookIDs[int(db.editions[i].BookID)]; !ok {
				continue
			}
		}
		if collectionIDs != nil {
			if _, ok := collectionIDs[int(db.editions[i].CollectionID)]; !ok {
				continue
			}
		}
		if collectionID == nil || db.editions[i].CollectionID == int64(*collectionID) {
			printIDs[int(db.editions[i].PrintID)] = struct{}{}
		}
	}

	ret := []*model.Print{}
	for k := range printIDs {
		if _, ok := db.prints[k]; ok {
			ret = append(ret, db.prints[k])
		}
	}
	return ret, nil
}
