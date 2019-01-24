package local

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getCategoryByID(id int) (*model.Category, error) {
	if p, ok := db.categories[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetCategoryByID return author by ID
func (db *Local) GetCategoryByID(id int) (*model.Category, error) {
	db.Lock()
	defer db.Unlock()
	return db.getCategoryByID(id)
}

func (db *Local) getCategoryByName(name string) (*model.Category, error) {
	for i := range db.categories {
		if db.categories[i].Name == name {
			return db.categories[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetCategoryByName return author by name
func (db *Local) GetCategoryByName(name string) (*model.Category, error) {
	db.Lock()
	defer db.Unlock()
	return db.getCategoryByName(name)
}

//InsertCategory insert author
func (db *Local) InsertCategory(name string) (*model.Category, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getCategoryByName(name)
	if err == nil {
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.categories) + 1
	v := &model.Category{ID: int64(idx), Name: name}
	db.categories[idx] = v
	return v, nil
}

//UpdateCategory update author
func (db *Local) UpdateCategory(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	v, err := db.getCategoryByID(id)
	if err != nil {
		return err
	}
	_, err = db.getCategoryByName(name)
	if err == nil {
		return model.ErrDuplicateKey
	}
	v.Name = name
	return nil
}

//DeleteCategory delete author
func (db *Local) DeleteCategory(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.categories[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.categories, id)
	return nil
}

//ListCategoriesByIDs list categories filtered by author id, class id
func (db *Local) ListCategoriesByIDs(authorID, classID *int) ([]*model.Category, error) {
	db.Lock()
	defer db.Unlock()

	categoryIDs := make(map[int]struct{})
	if authorID == nil && classID == nil {
		ret := []*model.Category{}
		for k := range db.categories {
			ret = append(ret, db.categories[k])
		}
		return ret, nil
	}

	seriesIDs := make(map[int]struct{})
	switch {
	case authorID != nil && classID != nil:
		bookIDs := make(map[int]struct{})
		for i := range db.authorships {
			if db.authorships[i].AuthorID == int64(*authorID) {
				bookIDs[int(db.authorships[i].BookID)] = struct{}{}
			}
		}

		for i := range bookIDs {
			if _, ok := db.books[i]; ok {
				for j := range db.classifications {
					if int(db.books[i].SeriesID) == db.classifications[j].SeriesID &&
						db.classifications[j].ClassID == *classID {
						seriesIDs[int(db.classifications[i].SeriesID)] = struct{}{}
					}
				}
			}
		}
	case authorID != nil:
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
	case classID != nil:
		for i := range db.classifications {
			if db.classifications[i].ClassID == *classID {
				seriesIDs[int(db.classifications[i].SeriesID)] = struct{}{}
			}
		}
	}
	for i := range seriesIDs {
		if _, ok := db.series[i]; ok {
			categoryIDs[int(db.series[i].CategoryID)] = struct{}{}
		}
	}

	ret := []*model.Category{}
	for k := range categoryIDs {
		if _, ok := db.categories[k]; ok {
			ret = append(ret, db.categories[k])
		}
	}
	return ret, nil
}
