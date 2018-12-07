package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getRoleByID(id int) (*model.Role, error) {
	if p, ok := db.roles[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

//GetRoleByID return author by ID
func (db *Local) GetRoleByID(id int) (*model.Role, error) {
	db.Lock()
	defer db.Unlock()
	return db.getRoleByID(id)
}

func (db *Local) getRoleByName(name string) (*model.Role, error) {
	for i := range db.roles {
		if db.roles[i].Name == name {
			return db.roles[i], nil
		}
	}
	return nil, model.ErrNotFound
}

//GetRoleByName return author by name
func (db *Local) GetRoleByName(name string) (*model.Role, error) {
	db.Lock()
	defer db.Unlock()
	return db.getRoleByName(name)
}

//GetRoleList list roles
func (db *Local) GetRoleList() ([]*model.Role, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.roles))
	i := 0
	for id := range db.roles {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]*model.Role, len(ids))
	for i, id := range ids {
		list[i] = db.roles[id]
	}
	return list, nil
}

//InsertRole insert author
func (db *Local) InsertRole(name string) (*model.Role, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getRoleByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.roles) + 1
	v := &model.Role{ID: int64(idx), Name: name}
	db.roles[idx] = v
	return v, nil
}

//UpdateRole update author
func (db *Local) UpdateRole(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	v, err := db.getRoleByID(id)
	if err != nil {
		return err
	}
	v.Name = name
	return nil
}

//DeleteRole delete author
func (db *Local) DeleteRole(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.roles[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.roles, id)
	return nil
}
