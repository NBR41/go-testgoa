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

//ListClasses list classes
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

//ListClassesBySeriesID list classes by series id
func (db *Local) ListClassesBySeriesID(seriesID int) ([]*model.Class, error) {
	db.Lock()
	defer db.Unlock()
	classIDs := make(map[int]struct{})
	for i := range db.classifications {
		if db.classifications[i].SeriesID == seriesID {
			classIDs[db.classifications[i].ClassID] = struct{}{}
		}
	}
	return db.buildClassList(classIDs)
}

//ListClassesByAuthorID list classes by author id
func (db *Local) ListClassesByAuthorID(authorID int) ([]*model.Class, error) {
	db.Lock()
	defer db.Unlock()

	bookIDs := make(map[int]struct{})
	for i := range db.authorships {
		if db.authorships[i].AuthorID == int64(authorID) {
			bookIDs[int(db.authorships[i].BookID)] = struct{}{}
		}
	}

	seriesIDs := make(map[int]struct{})
	for i := range bookIDs {
		if _, ok := db.books[i]; ok {
			seriesIDs[int(db.books[i].SeriesID)] = struct{}{}
		}
	}

	classIDs := make(map[int]struct{})
	for i := range seriesIDs {
		for j := range db.classifications {
			if i == db.classifications[j].SeriesID {
				classIDs[db.classifications[j].ClassID] = struct{}{}
			}
		}
	}
	return db.buildClassList(classIDs)
}

//ListClassesByCategoryID list classes by category id
func (db *Local) ListClassesByCategoryID(categoryID int) ([]*model.Class, error) {
	db.Lock()
	defer db.Unlock()

	seriesIDs := make(map[int]struct{})
	for i := range db.series {
		if db.series[i].CategoryID == int64(categoryID) {
			seriesIDs[int(db.series[i].ID)] = struct{}{}
		}
	}

	classIDs := make(map[int]struct{})
	for i := range seriesIDs {
		for j := range db.classifications {
			if i == db.classifications[j].SeriesID {
				classIDs[db.classifications[j].ClassID] = struct{}{}
			}
		}
	}
	return db.buildClassList(classIDs)
}

func (db *Local) buildClassList(idsb map[int]struct{}) ([]*model.Class, error) {
	ret := []*model.Class{}
	for k := range idsb {
		if _, ok := db.classes[k]; ok {
			ret = append(ret, db.classes[k])
		}
	}
	return ret, nil
}
