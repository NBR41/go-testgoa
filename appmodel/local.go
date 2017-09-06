package appmodel

import (
	"sort"
	"sync"

	"github.com/NBR41/go-testgoa/appsec"
)

// Local emulates a database driver using in-memory data structures.
type Local struct {
	sync.Mutex
	users      map[int]*User
	books      map[int]*Book
	ownerships map[int][]*Book
}

// NewLocal returns new instance of Local storage
func NewLocal() *Local {
	book := &Book{ID: 1, ISBN: "isbn-123", Name: "test1"}
	book2 := &Book{ID: 2, ISBN: "isbn-456", Name: "test2"}
	book3 := &Book{ID: 3, ISBN: "isbn-789", Name: "test3"}
	book4 := &Book{ID: 4, ISBN: "isbn-135", Name: "test4"}
	return &Local{
		users: map[int]*User{
			3: &User{ID: 3, Email: `user@myinventory.com`, Nickname: `user`, IsValidated: true, IsAdmin: false},
			2: &User{ID: 2, Email: `new@myinventory.com`, Nickname: `new`, IsValidated: false, IsAdmin: false},
			1: &User{
				ID:          1,
				Email:       `admin@myinventory.com`,
				Nickname:    `admin`,
				IsValidated: true,
				IsAdmin:     true,
				salt:        []byte("\xd6\xe8\u007f Yg\xbc\xe7@\x8b\xe4E\x9b\xb8\xc3\xeepZ\xe0\x90Z\xe4C\xd5%\xe7RP9a(\xfb"),
				password:    []byte("\"ya\xba4ґ},W9\x81p:\x81\x1d\xa9p\xb7\xa21qc\xf5i\x9d\x14\x92\x96\x97n\x91E\xf7\xf8>L\x86\xb3@Y\xe6)\xe7.\xb1A\x8b\xa5\x06d\\`\x03\xb3e\x1a\xc8\xd2P\xad\x965\x8e"),
			},
			4: &User{ID: 4, Email: `nobooks@myinventory.com`, Nickname: `nobooks`, IsValidated: true, IsAdmin: false},
		},
		books: map[int]*Book{
			1: book,
			2: book2,
			3: book3,
			4: book4,
		},
		ownerships: map[int][]*Book{
			1: []*Book{book, book4},
			2: []*Book{},
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
	return nil, ErrNotFound
}

// GetUserByEmailOrNickname returns user by email or nickname
func (db *Local) GetUserByEmailOrNickname(email, nickname string) (*User, error) {
	db.Lock()
	defer db.Unlock()
	return db.getUserByEmailOrNickname(email, nickname)
}

func (db *Local) getUserByEmailOrNickname(email, nickname string) (*User, error) {
	for i := range db.users {
		if db.users[i].Nickname == nickname || db.users[i].Email == email {
			return db.users[i], nil
		}
	}
	return nil, ErrNotFound
}

// GetUserByEmail returns user by email
func (db *Local) GetUserByEmail(email string) (*User, error) {
	db.Lock()
	defer db.Unlock()
	return db.getUserByEmail(email)
}

func (db *Local) getUserByEmail(email string) (*User, error) {
	for i := range db.users {
		if db.users[i].Email == email {
			return db.users[i], nil
		}
	}
	return nil, ErrNotFound
}

// GetUserByNickname returns user by nickname
func (db *Local) GetUserByNickname(nickname string) (*User, error) {
	db.Lock()
	defer db.Unlock()
	return db.getUserByNickname(nickname)
}

func (db *Local) getUserByNickname(nickname string) (*User, error) {
	for i := range db.users {
		if db.users[i].Nickname == nickname {
			return db.users[i], nil
		}
	}
	return nil, ErrNotFound
}

// GetAuthenticatedUser returns user if password matches email or nickname
func (db *Local) GetAuthenticatedUser(login, password string) (*User, error) {
	u, err := db.getUserByEmailOrNickname(login, login)
	if err != nil {
		return nil, err
	}

	ok, err := appsec.ComparePassword(password, u.salt, u.password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrInvalidCredentials
	}
	return u, nil
}

// InsertUser insert user
func (db *Local) InsertUser(email, nickname, password string) (*User, error) {
	salt, hash, err := appsec.CryptPassword(password)
	if err != nil {
		return nil, err
	}

	db.Lock()
	defer db.Unlock()
	u, err := db.getUserByEmailOrNickname(email, nickname)
	switch {
	case err != nil && err != ErrNotFound:
		return nil, err
	case err == nil:
		if u.Email == email {
			if u.Nickname == nickname {
				return nil, ErrDuplicateKey
			}
			return nil, ErrDuplicateEmail
		}
		return nil, ErrDuplicateNickname
	}
	idx := len(db.users) + 1
	u = &User{ID: int64(idx), Email: email, Nickname: nickname, salt: salt, password: hash}
	db.users[idx] = u
	db.ownerships[idx] = []*Book{}
	return u, nil
}

// UpdateUserNickname updates user nickname by ID
func (db *Local) UpdateUserNickname(id int, nickname string) error {
	db.Lock()
	defer db.Unlock()
	exU, err := db.getUserByNickname(nickname)
	if err != nil {
		if err == ErrNotFound {
			u, ok := db.users[id]
			if !ok {
				return ErrNotFound
			}
			u.Nickname = nickname
			return nil
		}
		return err
	}

	if exU.ID != int64(id) {
		return ErrDuplicateKey
	}
	return nil
}

// UpdateUserPassword updates user password by ID
func (db *Local) UpdateUserPassword(id int, password string) error {
	salt, hash, err := appsec.CryptPassword(password)
	if err != nil {
		return err
	}

	db.Lock()
	defer db.Unlock()
	u, ok := db.users[id]
	if !ok {
		return ErrNotFound
	}
	u.salt = salt
	u.password = hash
	return nil
}

// UpdateUserActivation update user activation by ID
func (db *Local) UpdateUserActivation(id int, validated bool) error {
	db.Lock()
	defer db.Unlock()
	u, ok := db.users[id]
	if !ok {
		return ErrNotFound
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
		return ErrNotFound
	}
	delete(db.users, id)
	return nil
}

// InsertBook inserts book
func (db *Local) InsertBook(isbn, name string) (*Book, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getBookByISBN(isbn)
	switch {
	case err != nil && err != ErrNotFound:
		return nil, err
	case err == nil:
		return nil, ErrDuplicateKey
	}
	idx := len(db.books) + 1
	b := &Book{ID: int64(idx), ISBN: isbn, Name: name}
	db.books[idx] = b
	return b, nil
}

// GetBookByID returns book by ID
func (db *Local) GetBookByID(id int) (*Book, error) {
	db.Lock()
	defer db.Unlock()
	return db.getBookByID(id)
}

func (db *Local) getBookByID(id int) (*Book, error) {
	if p, ok := db.books[id]; ok {
		return p, nil
	}
	return nil, ErrNotFound
}

// GetBookByName returns book by name
func (db *Local) GetBookByName(name string) (*Book, error) {
	db.Lock()
	defer db.Unlock()
	for i := range db.books {
		if db.books[i].Name == name {
			return db.books[i], nil
		}
	}
	return nil, ErrNotFound
}

// GetBookByISBN returns book by isbn
func (db *Local) GetBookByISBN(isbn string) (*Book, error) {
	db.Lock()
	defer db.Unlock()
	return db.getBookByISBN(isbn)
}

func (db *Local) getBookByISBN(isbn string) (*Book, error) {
	for i := range db.books {
		if db.books[i].ISBN == isbn {
			return db.books[i], nil
		}
	}
	return nil, ErrNotFound
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
	b, err := db.getBookByID(id)
	if err != nil {
		return err
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
		return ErrNotFound
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
		return nil, ErrNotFound
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
		return nil, ErrNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			return &Ownership{UserID: int64(userID), BookID: int64(bookID), Book: l[i]}, nil
		}
	}
	return nil, ErrNotFound
}

// InsertOwnership inserts user book association
func (db *Local) InsertOwnership(userID, bookID int) (*Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, ErrNotFound
	}

	b, ok := db.books[bookID]
	if !ok {
		return nil, ErrNotFound
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
		return ErrNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			db.ownerships[userID] = append(l[:i], l[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
