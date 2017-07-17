package store

import (
	"sort"
	"sync"
)

// Local emulates a database driver using in-memory data structures.
type Local struct {
	sync.Mutex
	users      map[int]*User
	books      map[int]*Book
	ownerships map[int][]*Book
}

func NewLocal() *Local {
	book := &Book{ID: 1, Name: "test1"}
	book2 := &Book{ID: 1, Name: "test2"}
	book3 := &Book{ID: 1, Name: "test3"}
	book4 := &Book{ID: 1, Name: "test4"}
	return &Local{
		users: map[int]*User{
			1: &User{ID: 1, Email: `admin@myinventory.com`, Nickname: `admin`, IsVerified: true, IsAdmin: true},
			2: &User{ID: 2, Email: `new@myinventory.com`, Nickname: `new`, IsVerified: false, IsAdmin: false},
			3: &User{ID: 3, Email: `user@myinventory.com`, Nickname: `user`, IsVerified: true, IsAdmin: false},
			4: &User{ID: 4, Email: `nobooks@myinventory.com`, Nickname: `nobooks`, IsVerified: true, IsAdmin: false},
		},
		books: map[int]*Book{
			1: book,
			2: book2,
			3: book3,
			4: book4,
		},
		ownerships: map[int][]*Book{
			1: []*Book{book, book4},
			3: []*Book{book2, book3},
		},
	}
}

// Close close the connextion
func (db *Local) Close() error {
	return nil
}

// GetUserList returns user list
func (db *Local) GetUserList() ([]User, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.users))
	i := 0
	for id := range db.users {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]User, len(ids))
	for i, id := range ids {
		list[i] = *db.users[id]
	}
	return list, nil
}

// GetUserByID returns user by ID
func (db *Local) GetUserByID(id int) (*User, error) {
	db.Lock()
	defer db.Unlock()
	if p, ok := db.users[id]; ok {
		return p, nil
	}
	return nil, ErrUserNotFound
}

// GetUserByEmailOrNickname returns user by email or nickname
func (db *Local) GetUserByEmailOrNickname(email, nickname string) (*User, error) {
	db.Lock()
	defer db.Unlock()
	for i := range db.users {
		if db.users[i].Nickname == nickname || db.users[i].Email == email {
			return db.users[i], nil
		}
	}
	return nil, ErrUserNotFound
}

// GetAuthenticatedUser returns user if password matches email or nickname
func (db *Local) GetAuthenticatedUser(password, email, nickname string) (*User, error) {
	db.Lock()
	defer db.Unlock()
	for i := range db.users {
		if db.users[i].Nickname == nickname || db.users[i].Email == email {
			if db.users[i].Password == password {
				return db.users[i], nil
			}
			return nil, ErrInvalidCredentials
		}
	}
	return nil, ErrUserNotFound
}

// InsertUser insert user
func (db *Local) InsertUser(nickname, email, password string) (*User, error) {
	db.Lock()
	defer db.Unlock()
	idx := len(db.users)
	u := &User{ID: int64(idx), Email: email, Nickname: nickname, Password: password}
	db.users[idx] = u
	return u, nil
}

// UpdateUserNickname updates user nickname by ID
func (db *Local) UpdateUserNickname(id int, nickname string) error {
	db.Lock()
	defer db.Unlock()
	u, ok := db.users[id]
	if !ok {
		return ErrUserNotFound
	}
	u.Nickname = nickname
	return nil
}

// UpdateUserPassword updates user password by ID
func (db *Local) UpdateUserPassword(id int, password string) error {
	db.Lock()
	defer db.Unlock()
	u, ok := db.users[id]
	if !ok {
		return ErrUserNotFound
	}
	u.Password = password
	return nil
}

// UpdateUserActivation update user activation by ID
func (db *Local) UpdateUserActivation(id int, activated bool) error {
	db.Lock()
	defer db.Unlock()
	u, ok := db.users[id]
	if !ok {
		return ErrUserNotFound
	}
	u.IsVerified = activated
	return nil
}

// DeleteUser deletes user by ID
func (db *Local) DeleteUser(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.users[id]
	if !ok {
		return ErrUserNotFound
	}
	delete(db.users, id)
	return nil
}

// InsertBook inserts book
func (db *Local) InsertBook(name string) (*Book, error) {
	db.Lock()
	defer db.Unlock()
	idx := len(db.books)
	b := &Book{ID: int64(idx), Name: name}
	db.books[idx] = b
	return b, nil
}

// GetBookByID returns book by ID
func (db *Local) GetBookByID(id int) (*Book, error) {
	db.Lock()
	defer db.Unlock()
	if p, ok := db.books[id]; ok {
		return p, nil
	}
	return nil, ErrBookNotFound
}

// GetBookList returns book list
func (db *Local) GetBookList() ([]Book, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.books))
	i := 0
	for id := range db.books {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]Book, len(ids))
	for i, id := range ids {
		list[i] = *db.books[id]
	}
	return list, nil
}

// UpdateBook update book infos
func (db *Local) UpdateBook(id int, name string) error {
	db.Lock()
	defer db.Unlock()
	b, ok := db.books[id]
	if !ok {
		return ErrBookNotFound
	}
	b.Name = name
	return nil
}

// DeleteBook delete book by ID
func (db *Local) DeleteBook(id int) error {
	db.Lock()
	defer db.Unlock()
	_, ok := db.books[id]
	if !ok {
		return ErrBookNotFound
	}
	delete(db.books, id)
	return nil
}

// GetOwnershipList returns book list by user ID
func (db *Local) GetOwnershipList(userID int) ([]Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, ErrUserNotFound
	}

	var list = []Ownership{}
	for i := range l {
		list = append(
			list,
			Ownership{
				UserID: int64(userID),
				BookID: l[i].ID,
				Book:   l[i],
			},
		)
	}
	return list, nil
}

// GetOwnership returns user book association
func (db *Local) GetOwnership(userID, bookID int) (*Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, ErrUserNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			return &Ownership{UserID: int64(userID), BookID: int64(bookID), Book: l[i]}, nil
		}
	}
	return nil, ErrOwnershipNotFound
}

// InsertOwnership inserts user book association
func (db *Local) InsertOwnership(userID, bookID int) (*Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, ErrUserNotFound
	}

	b, ok := db.books[bookID]
	if !ok {
		return nil, ErrBookNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			return nil, ErrDuplicateKey
		}
	}

	db.ownerships[userID] = append(l, b)
	return &Ownership{UserID: int64(userID), BookID: int64(bookID), Book: b}, nil
}

func (db *Local) UpdateOwnership(userID, bookID int) error {
	return nil
}

// DeleteOwnership deletes user book association
func (db *Local) DeleteOwnership(userID, bookID int) error {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return ErrUserNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			db.ownerships[userID] = append(l[:i], l[i+1:]...)
			break
		}
	}
	return ErrOwnershipNotFound
}
