package local

import (
	"sort"

	"github.com/NBR41/go-testgoa/internal/model"
)

// GetUserList returns user list
func (db *Local) GetUserList() ([]model.User, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.users))
	i := 0
	for id := range db.users {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]model.User, len(ids))
	for i, id := range ids {
		list[i] = *db.users[id]
	}
	return list, nil
}

// GetUserByID returns user by ID
func (db *Local) GetUserByID(id int) (*model.User, error) {
	db.Lock()
	defer db.Unlock()
	if p, ok := db.users[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

// GetUserByEmailOrNickname returns user by email or nickname
func (db *Local) GetUserByEmailOrNickname(email, nickname string) (*model.User, error) {
	db.Lock()
	defer db.Unlock()
	return db.getUserByEmailOrNickname(email, nickname)
}

func (db *Local) getUserByEmailOrNickname(email, nickname string) (*model.User, error) {
	for i := range db.users {
		if db.users[i].Nickname == nickname || db.users[i].Email == email {
			return db.users[i], nil
		}
	}
	return nil, model.ErrNotFound
}

// GetUserByEmail returns user by email
func (db *Local) GetUserByEmail(email string) (*model.User, error) {
	db.Lock()
	defer db.Unlock()
	return db.getUserByEmail(email)
}

func (db *Local) getUserByEmail(email string) (*model.User, error) {
	for i := range db.users {
		if db.users[i].Email == email {
			return db.users[i], nil
		}
	}
	return nil, model.ErrNotFound
}

// GetUserByNickname returns user by nickname
func (db *Local) GetUserByNickname(nickname string) (*model.User, error) {
	db.Lock()
	defer db.Unlock()
	return db.getUserByNickname(nickname)
}

func (db *Local) getUserByNickname(nickname string) (*model.User, error) {
	for i := range db.users {
		if db.users[i].Nickname == nickname {
			return db.users[i], nil
		}
	}
	return nil, model.ErrNotFound
}

// GetAuthenticatedUser returns user if password matches email or nickname
func (db *Local) GetAuthenticatedUser(login, password string) (*model.User, error) {
	u, err := db.getUserByEmailOrNickname(login, login)
	if err != nil {
		return nil, err
	}

	ok, err := db.pass.ComparePassword(password, u.Salt, u.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, model.ErrInvalidCredentials
	}
	return u, nil
}

// InsertUser insert user
func (db *Local) InsertUser(email, nickname, password string) (*model.User, error) {
	salt, hash, err := db.pass.CryptPassword(password)
	if err != nil {
		return nil, err
	}

	db.Lock()
	defer db.Unlock()
	u, err := db.getUserByEmailOrNickname(email, nickname)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		if u.Email == email {
			if u.Nickname == nickname {
				return nil, model.ErrDuplicateKey
			}
			return nil, model.ErrDuplicateEmail
		}
		return nil, model.ErrDuplicateNickname
	}
	idx := len(db.users) + 1
	u = &model.User{ID: int64(idx), Email: email, Nickname: nickname, Salt: salt, Password: hash}
	db.users[idx] = u
	db.ownerships[idx] = []*model.Book{}
	return u, nil
}

// UpdateUserNickname updates user nickname by ID
func (db *Local) UpdateUserNickname(id int, nickname string) error {
	db.Lock()
	defer db.Unlock()
	exU, err := db.getUserByNickname(nickname)
	if err != nil {
		if err == model.ErrNotFound {
			u, ok := db.users[id]
			if !ok {
				return model.ErrNotFound
			}
			u.Nickname = nickname
			return nil
		}
		return err
	}

	if exU.ID != int64(id) {
		return model.ErrDuplicateKey
	}
	return nil
}

// UpdateUserPassword updates user password by ID
func (db *Local) UpdateUserPassword(id int, password string) error {
	salt, hash, err := db.pass.CryptPassword(password)
	if err != nil {
		return err
	}

	db.Lock()
	defer db.Unlock()
	u, ok := db.users[id]
	if !ok {
		return model.ErrNotFound
	}
	u.Salt = salt
	u.Password = hash
	return nil
}

// UpdateUserActivation update user activation by ID
func (db *Local) UpdateUserActivation(id int, validated bool) error {
	db.Lock()
	defer db.Unlock()
	u, ok := db.users[id]
	if !ok {
		return model.ErrNotFound
	}
	u.IsValidated = validated
	return nil
}

// DeleteUser deletes user by ID
func (db *Local) DeleteUser(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.users[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(db.users, id)
	return nil
}
