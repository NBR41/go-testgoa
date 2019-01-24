package local

import (
	"github.com/NBR41/go-testgoa/internal/model"
	"sort"
)

func (db *Local) getAuthorByID(id int) (*model.Author, error) {
	if p, ok := db.authors[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetAuthorByID return author by ID
func (db *Local) GetAuthorByID(id int) (*model.Author, error) {
	db.Lock()
	defer db.Unlock()
	return db.getAuthorByID(id)
}

func (db *Local) getAuthorByName(name string) (*model.Author, error) {
	for i := range db.authors {
		if db.authors[i].Name == name {
			return db.authors[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetAuthorByName return author by name
func (db *Local) GetAuthorByName(name string) (*model.Author, error) {
	db.Lock()
	defer db.Unlock()
	return db.getAuthorByName(name)
}

//ListAuthors list authors
func (db *Local) ListAuthors() ([]*model.Author, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.authors))
	i := 0
	for id := range db.authors {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Author, len(ids))
	for i, id := range ids {
		list[i] = db.authors[id]
	}
	return list, nil
}

//InsertAuthor insert author
func (db *Local) InsertAuthor(name string) (*model.Author, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getAuthorByName(name)
	if err == nil {
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.authors) + 1
	v := &model.Author{ID: int64(idx), Name: name}
	db.authors[idx] = v
	return v, nil
}

//UpdateAuthor update author
func (db *Local) UpdateAuthor(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	v, err := db.getAuthorByID(id)
	if err != nil {
		return err
	}

	_, err = db.getAuthorByName(name)
	if err == nil {
		return model.ErrDuplicateKey
	}

	v.Name = name
	return nil
}

//DeleteAuthor delete author
func (db *Local) DeleteAuthor(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.authors[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.authors, id)
	return nil
}

//ListAuthorsByIDs returns filtered author list
func (db *Local) ListAuthorsByIDs(categoryID, roleID *int) ([]*model.Author, error) {
	db.Lock()
	defer db.Unlock()
	var bookIDs map[int]struct{}

	if categoryID != nil {
		seriesIDs := []int{}
		for i := range db.series {
			if db.series[i].CategoryID == int64(*categoryID) {
				seriesIDs = append(seriesIDs, i)
			}
		}

		bookIDs = make(map[int]struct{})
		for i := range seriesIDs {
			for j := range db.books {
				if db.books[j].SeriesID == int64(seriesIDs[i]) {
					bookIDs[j] = struct{}{}
				}
			}
		}
	}

	authorIDs := make(map[int]struct{})
	for i := range db.authorships {
		if bookIDs != nil {
			if _, ok := bookIDs[int(db.authorships[i].BookID)]; !ok {
				continue
			}
		}
		if roleID == nil || db.authorships[i].RoleID == int64(*roleID) {
			authorIDs[int(db.authorships[i].AuthorID)] = struct{}{}
		}
	}

	ret := []*model.Author{}
	for k := range authorIDs {
		if _, ok := db.authors[k]; ok {
			ret = append(ret, db.authors[k])
		}
	}
	return ret, nil
}
