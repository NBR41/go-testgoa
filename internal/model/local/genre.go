package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getGenreByID(id int) (*model.Genre, error) {
	if p, ok := db.genres[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetGenreByID return author by ID
func (db *Local) GetGenreByID(id int) (*model.Genre, error) {
	db.Lock()
	defer db.Unlock()
	return db.getGenreByID(id)
}

func (db *Local) getGenreByName(name string) (*model.Genre, error) {
	for i := range db.genres {
		if db.genres[i].Name == name {
			return db.genres[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetGenreByName return author by name
func (db *Local) GetGenreByName(name string) (*model.Genre, error) {
	db.Lock()
	defer db.Unlock()
	return db.getGenreByName(name)
}

//GetGenreList list genres
func (db *Local) GetGenreList() ([]*model.Genre, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.genres))
	i := 0
	for id := range db.genres {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Genre, len(ids))
	for i, id := range ids {
		list[i] = db.genres[id]
	}
	return list, nil
}

//InsertGenre insert author
func (db *Local) InsertGenre(name string) (*model.Genre, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getGenreByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.genres) + 1
	v := &model.Genre{ID: int64(idx), Name: name}
	db.genres[idx] = v
	return v, nil
}

//UpdateGenre update author
func (db *Local) UpdateGenre(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	v, err := db.getGenreByID(id)
	if err != nil {
		return err
	}
	v.Name = name
	return nil
}

//DeleteGenre delete author
func (db *Local) DeleteGenre(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.genres[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.genres, id)
	return nil
}
