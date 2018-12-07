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

//GetCategoryList list categories
func (db *Local) GetCategoryList() ([]*model.Category, error) {
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
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
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
