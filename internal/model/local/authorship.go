package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getAuthorshipByID(id int) (*model.Authorship, error) {
	if p, ok := db.authorships[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetAuthorshipByID get authorship by id
func (db *Local) GetAuthorshipByID(id int) (*model.Authorship, error) {
	db.Lock()
	defer db.Unlock()
	return db.getAuthorshipByID(id)
}

//ListAuthorships list authorships
func (db *Local) ListAuthorships() ([]*model.Authorship, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.authorships))
	i := 0
	for id := range db.authorships {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Authorship, len(ids))
	for i, id := range ids {
		list[i] = db.authorships[id]
	}
	return list, nil
}

//ListAuthorshipsByBookID list all authorships for a book
func (db *Local) ListAuthorshipsByBookID(bookID int) ([]*model.Authorship, error) {
	// TODO
	return nil, nil
}

//InsertAuthorship insert an authorship
func (db *Local) InsertAuthorship(authorID, bookID, roleID int) (*model.Authorship, error) {
	db.Lock()
	defer db.Unlock()
	for _, v := range db.authorships {
		if v.AuthorID == int64(authorID) && v.BookID == int64(bookID) && v.RoleID == int64(roleID) {
			return nil, model.ErrDuplicateKey
		}
	}

	if _, ok := db.authors[authorID]; !ok {
		return nil, model.ErrInvalidID
	}
	if _, ok := db.books[bookID]; !ok {
		return nil, model.ErrInvalidID
	}
	if _, ok := db.roles[roleID]; !ok {
		return nil, model.ErrInvalidID
	}

	idx := len(db.authorships) + 1
	v := &model.Authorship{ID: int64(idx), AuthorID: int64(authorID), RoleID: int64(roleID), BookID: int64(bookID)}
	db.authorships[idx] = v
	return v, nil
}

//DeleteAuthorship delete an authorship
func (db *Local) DeleteAuthorship(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.authorships[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.authorships, id)
	return nil
}
