package local

import (
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

//InsertClass insert author
func (db *Local) InsertClass(name string) (*model.Class, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getClassByName(name)
	if err == nil {
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
	_, err = db.getClassByName(name)
	if err == nil {
		return model.ErrDuplicateKey
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

//ListClassesByIDs returns filtered class list
func (db *Local) ListClassesByIDs(authorID, categoryID, seriesID *int) ([]*model.Class, error) {
	db.Lock()
	defer db.Unlock()

	if authorID == nil && categoryID == nil && seriesID == nil {
		ret := []*model.Class{}
		for k := range db.classes {
			ret = append(ret, db.classes[k])
		}
		return ret, nil
	}

	var seriesIDs map[int]struct{}
	if authorID != nil {
		seriesIDs = make(map[int]struct{})
		bookIDs := make(map[int]struct{})
		for i := range db.authorships {
			if db.authorships[i].AuthorID == int64(*authorID) {
				bookIDs[int(db.authorships[i].BookID)] = struct{}{}
			}
		}
		for i := range bookIDs {
			if _, ok := db.books[i]; ok {
				seriesIDs[int(db.books[i].SeriesID)] = struct{}{}
			}
		}
	}

	finalSeriesIDs := make(map[int]struct{})
	for i := range db.series {
		if seriesIDs != nil {
			if _, ok := seriesIDs[i]; !ok {
				continue
			}
		}
		if (categoryID == nil || db.series[i].CategoryID == int64(*categoryID)) &&
			(seriesID == nil || i == *seriesID) {
			finalSeriesIDs[i] = struct{}{}
		}
	}
	var classIDs = make(map[int]struct{})
	for i := range db.classifications {
		if _, ok := finalSeriesIDs[db.classifications[i].SeriesID]; ok {
			classIDs[db.classifications[i].ClassID] = struct{}{}
		}
	}

	ret := []*model.Class{}
	for k := range classIDs {
		if _, ok := db.classes[k]; ok {
			ret = append(ret, db.classes[k])
		}
	}
	return ret, nil
}
