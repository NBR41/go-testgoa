package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getSeriesByID(id int) (*model.Series, error) {
	if p, ok := db.series[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetSeriesByID return series by ID
func (db *Local) GetSeriesByID(id int) (*model.Series, error) {
	db.Lock()
	defer db.Unlock()
	return db.getSeriesByID(id)
}

func (db *Local) getSeriesByName(name string) (*model.Series, error) {
	for i := range db.series {
		if db.series[i].Name == name {
			return db.series[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetSeriesByName return series by name
func (db *Local) GetSeriesByName(name string) (*model.Series, error) {
	db.Lock()
	defer db.Unlock()
	return db.getSeriesByName(name)
}

//ListSeries list series
func (db *Local) ListSeries() ([]*model.Series, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.series))
	i := 0
	for id := range db.series {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Series, len(ids))
	for i, id := range ids {
		list[i] = db.series[id]
	}
	return list, nil
}

//ListSeriesByIDs list series by different filters
func (db *Local) ListSeriesByIDs(authorID, roleID, categoryID, classID *int) ([]*model.Series, error) {
	db.Lock()
	defer db.Unlock()

	bookIDs := make(map[int]struct{})
	for i := range db.authorships {
		if (authorID == nil || *authorID == int(db.authorships[i].AuthorID)) &&
			(roleID == nil || *roleID == int(db.authorships[i].RoleID)) {
			bookIDs[int(db.authorships[i].BookID)] = struct{}{}
		}
	}

	seriesIDs := make(map[int]struct{})
	for i := range db.books {
		if _, ok := bookIDs[i]; ok {
			seriesIDs[int(db.books[i].SeriesID)] = struct{}{}
		}
	}

	var finalIDs map[int]struct{}
	if classID != nil {
		finalIDs = make(map[int]struct{})
		for i := range db.classifications {
			if _, ok := seriesIDs[int(db.classifications[i].SeriesID)]; ok && int(db.classifications[i].ClassID) == *classID {
				finalIDs[int(db.classifications[i].SeriesID)] = struct{}{}
			}
		}
	} else {
		finalIDs = seriesIDs
	}

	ret := []*model.Series{}
	for k := range finalIDs {
		if _, ok := db.series[k]; ok {
			if categoryID == nil || db.series[k].CategoryID == int64(*categoryID) {
				ret = append(ret, db.series[k])
			}
		}
	}
	return ret, nil
}

//ListSeriesByCollectionID list series by collection id
func (db *Local) ListSeriesByCollectionID(collectionID int) ([]*model.Series, error) {
	db.Lock()
	defer db.Unlock()
	bookIDs := make(map[int]struct{})
	for i := range db.editions {
		if int64(collectionID) == db.editions[i].CollectionID {
			bookIDs[int(db.editions[i].BookID)] = struct{}{}
		}
	}
	seriesIDs := make(map[int]struct{})
	for i := range bookIDs {
		if _, ok := db.books[i]; ok {
			seriesIDs[int(db.books[i].SeriesID)] = struct{}{}
		}
	}
	return db.buildSeriesList(seriesIDs)
}

//ListSeriesByCollectionIDPrintID list series by collection and print ids
func (db *Local) ListSeriesByCollectionIDPrintID(collectionID, printID int) ([]*model.Series, error) {
	db.Lock()
	defer db.Unlock()
	bookIDs := make(map[int]struct{})
	for i := range db.editions {
		if int64(collectionID) == db.editions[i].CollectionID && int64(printID) == db.editions[i].PrintID {
			bookIDs[int(db.editions[i].BookID)] = struct{}{}
		}
	}
	seriesIDs := make(map[int]struct{})
	for i := range bookIDs {
		if _, ok := db.books[i]; ok {
			seriesIDs[int(db.books[i].SeriesID)] = struct{}{}
		}
	}
	return db.buildSeriesList(seriesIDs)
}

func (db *Local) buildSeriesList(idsb map[int]struct{}) ([]*model.Series, error) {
	ret := []*model.Series{}
	for k := range idsb {
		if _, ok := db.series[k]; ok {
			ret = append(ret, db.series[k])
		}
	}
	return ret, nil
}

//InsertSeries insert series
func (db *Local) InsertSeries(name string, categoryID int) (*model.Series, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getSeriesByName(name)
	if err == nil {
		return nil, model.ErrDuplicateKey
	}
	_, err = db.getCategoryByID(categoryID)
	if err != nil {
		return nil, model.ErrInvalidID
	}
	idx := len(db.series) + 1
	v := &model.Series{ID: int64(idx), Name: name, CategoryID: int64(categoryID)}
	db.series[idx] = v
	return v, nil
}

//UpdateSeries update series
func (db *Local) UpdateSeries(id int, name *string, categoryID *int) error {
	db.Lock()
	defer db.Unlock()
	b, err := db.getSeriesByID(id)
	if err != nil {
		return err
	}
	if name != nil {
		_, err := db.getSeriesByName(*name)
		if err == nil {
			return model.ErrDuplicateKey
		}
		b.Name = *name
	}
	if categoryID != nil {
		_, err := db.getCategoryByID(*categoryID)
		if err == model.ErrNotFound {
			return model.ErrInvalidID
		}
		b.CategoryID = int64(*categoryID)
	}
	return nil
}

//DeleteSeries delete series
func (db *Local) DeleteSeries(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.series[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.series, id)
	return nil
}
