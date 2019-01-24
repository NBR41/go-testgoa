package local

import (
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

//ListSeriesByIDs list series by different filters
func (db *Local) ListSeriesByIDs(authorID, categoryID, classID, roleID *int) ([]*model.Series, error) {
	db.Lock()
	defer db.Unlock()

	if authorID == nil && roleID == nil && categoryID == nil && classID == nil {
		return db.listSeries()
	}

	var bookSeriesIDs map[int]struct{}
	if authorID != nil || roleID != nil {
		bookIDs := make(map[int]struct{})
		for i := range db.authorships {
			if (authorID == nil || *authorID == int(db.authorships[i].AuthorID)) &&
				(roleID == nil || *roleID == int(db.authorships[i].RoleID)) {
				bookIDs[int(db.authorships[i].BookID)] = struct{}{}
			}
		}
		bookSeriesIDs = make(map[int]struct{})
		for i := range bookIDs {
			if _, ok := db.books[i]; ok {
				bookSeriesIDs[int(db.books[i].SeriesID)] = struct{}{}
			}
		}
	}

	var classSeriesIDs map[int]struct{}
	if classID != nil {
		classSeriesIDs = make(map[int]struct{})
		for i := range db.classifications {
			if int(db.classifications[i].ClassID) == *classID {
				classSeriesIDs[int(db.classifications[i].SeriesID)] = struct{}{}
			}
		}
	}

	seriesIDs := make(map[int]struct{})
	for i := range db.series {
		if bookSeriesIDs != nil {
			if _, ok := bookSeriesIDs[i]; !ok {
				continue
			}
		}
		if classSeriesIDs != nil {
			if _, ok := classSeriesIDs[i]; !ok {
				continue
			}
		}
		if categoryID == nil || db.series[i].CategoryID == int64(*categoryID) {
			seriesIDs[i] = struct{}{}
		}
	}
	return db.buildSeriesList(seriesIDs)
}

//ListSeriesByEditionIDs return a filtered series list
func (db *Local) ListSeriesByEditionIDs(collectionID, editorID, printID *int) ([]*model.Series, error) {
	db.Lock()
	defer db.Unlock()

	if collectionID == nil && editorID == nil && printID == nil {
		return db.listSeries()
	}

	var collectionIDs map[int]struct{}
	if editorID != nil {
		collectionIDs = make(map[int]struct{})
		for i := range db.collections {
			if db.collections[i].EditorID == int64(*editorID) {
				collectionIDs[int(db.collections[i].ID)] = struct{}{}
			}
		}
	}

	var bookIDs = make(map[int]struct{})
	for i := range db.editions {
		if collectionIDs != nil {
			if _, ok := collectionIDs[int(db.editions[i].CollectionID)]; !ok {
				continue
			}
		}

		if (printID == nil || db.editions[i].PrintID == int64(*printID)) &&
			(collectionID == nil || db.editions[i].CollectionID == int64(*collectionID)) {
			bookIDs[int(db.editions[i].BookID)] = struct{}{}
		}
	}

	var seriesIDs = make(map[int]struct{})
	for i := range bookIDs {
		if _, ok := db.books[i]; ok {
			seriesIDs[int(db.books[i].SeriesID)] = struct{}{}
		}
	}

	return db.buildSeriesList(seriesIDs)
}

func (db *Local) listSeries() ([]*model.Series, error) {
	ret := []*model.Series{}
	for i := range db.series {
		ret = append(ret, db.series[i])
	}
	return ret, nil
}

func (db *Local) buildSeriesList(ids map[int]struct{}) ([]*model.Series, error) {
	ret := []*model.Series{}
	for i := range ids {
		if _, ok := db.series[i]; ok {
			ret = append(ret, db.series[i])
		}
	}
	return ret, nil
}
