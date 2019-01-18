package local

import (
	"sort"

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

//ListCategories list categories
func (db *Local) ListCategories() ([]*model.Category, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.categories))
	i := 0
	for id := range db.categories {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Category, len(ids))
	for i, id := range ids {
		list[i] = db.categories[id]
	}
	return list, nil
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

//ListCategoriesByAuthorID list categories by author id
func (db *Local) ListCategoriesByAuthorID(authorID int) ([]*model.Category, error) {
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

	categoryIDs := make(map[int]struct{})
	for i := range seriesIDs {
		if _, ok := db.series[i]; ok {
			categoryIDs[int(db.series[i].CategoryID)] = struct{}{}
		}
	}
	return db.buildCategoryList(categoryIDs)
}

//ListCategoriesByClassID list categories by class id
func (db *Local) ListCategoriesByClassID(classID int) ([]*model.Category, error) {
	db.Lock()
	defer db.Unlock()
	seriesIDs := make(map[int]struct{})
	for i := range db.classifications {
		if db.classifications[i].ClassID == classID {
			seriesIDs[int(db.classifications[i].SeriesID)] = struct{}{}
		}
	}

	categoryIDs := make(map[int]struct{})
	for i := range seriesIDs {
		if _, ok := db.series[i]; ok {
			categoryIDs[int(db.series[i].CategoryID)] = struct{}{}
		}
	}

	return db.buildCategoryList(categoryIDs)
}

func (db *Local) buildCategoryList(idsb map[int]struct{}) ([]*model.Category, error) {
	ret := []*model.Category{}
	for k := range idsb {
		if _, ok := db.categories[k]; ok {
			ret = append(ret, db.categories[k])
		}
	}
	return ret, nil
}
