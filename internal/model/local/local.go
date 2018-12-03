package local

import (
	"sort"
	"sync"

	"github.com/NBR41/go-testgoa/internal/model"
)

// Local emulates a database driver using in-memory data structures.
type Local struct {
	pass model.Passworder
	sync.Mutex
	users      map[int]*model.User
	books      map[int]*model.Book
	ownerships map[int][]*model.Book
}

// New returns new instance of Local storage
func New(pass model.Passworder) *Local {
	book := &model.Book{ID: 1, ISBN: "isbn-123", Name: "test1"}
	book2 := &model.Book{ID: 2, ISBN: "isbn-456", Name: "test2"}
	book3 := &model.Book{ID: 3, ISBN: "isbn-789", Name: "test3"}
	book4 := &model.Book{ID: 4, ISBN: "isbn-135", Name: "test4"}
	return &Local{
		pass: pass,
		users: map[int]*model.User{
			3: &model.User{ID: 3, Email: `user@myinventory.com`, Nickname: `user`, IsValidated: true, IsAdmin: false},
			2: &model.User{ID: 2, Email: `new@myinventory.com`, Nickname: `new`, IsValidated: false, IsAdmin: false},
			1: &model.User{
				ID:          1,
				Email:       `admin@myinventory.com`,
				Nickname:    `admin`,
				IsValidated: true,
				IsAdmin:     true,
				Salt:        []byte("\xd6\xe8\u007f Yg\xbc\xe7@\x8b\xe4E\x9b\xb8\xc3\xeepZ\xe0\x90Z\xe4C\xd5%\xe7RP9a(\xfb"),
				Password:    []byte("'\xeb\xbe\x1f\xbaaG\xe1&>\x9f \u007f\xc94^\xdf\xca*\xdb\xf6<\x05\x05A8q\x94\xd0k\xc23\xf9\xd5\xdb-\x8f\x1c\f\xa5\xa1\xcf\xcf\xe1\t\xde\xf4\x89\x81B\x06\x16\x0ecQ\x94*\xa0D\x82\x1dUeJ"),
			},
			4: &model.User{ID: 4, Email: `nobooks@myinventory.com`, Nickname: `nobooks`, IsValidated: true, IsAdmin: false},
		},
		books: map[int]*model.Book{
			1: book,
			2: book2,
			3: book3,
			4: book4,
		},
		ownerships: map[int][]*model.Book{
			1: []*model.Book{book, book4},
			2: []*model.Book{},
			3: []*model.Book{book2, book3},
		},
	}
}

// Close close the connextion
func (db *Local) Close() error {
	return nil
}

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

// InsertBook inserts book
func (db *Local) InsertBook(isbn, name string) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	_, err := db.getBookByISBN(isbn)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	idx := len(db.books) + 1
	b := &model.Book{ID: int64(idx), ISBN: isbn, Name: name}
	db.books[idx] = b
	return b, nil
}

// GetBookByID returns book by ID
func (db *Local) GetBookByID(id int) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	return db.getBookByID(id)
}

func (db *Local) getBookByID(id int) (*model.Book, error) {
	if p, ok := db.books[id]; ok {
		return p, nil
	}
	return nil, model.ErrNotFound
}

// GetBookByName returns book by name
func (db *Local) GetBookByName(name string) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	for i := range db.books {
		if db.books[i].Name == name {
			return db.books[i], nil
		}
	}
	return nil, model.ErrNotFound
}

// GetBookByISBN returns book by isbn
func (db *Local) GetBookByISBN(isbn string) (*model.Book, error) {
	db.Lock()
	defer db.Unlock()
	return db.getBookByISBN(isbn)
}

func (db *Local) getBookByISBN(isbn string) (*model.Book, error) {
	for i := range db.books {
		if db.books[i].ISBN == isbn {
			return db.books[i], nil
		}
	}
	return nil, model.ErrNotFound
}

// GetBookList returns book list
func (db *Local) GetBookList() ([]model.Book, error) {
	db.Lock()
	defer db.Unlock()
	ids := make([]int, len(db.books))
	i := 0
	for id := range db.books {
		ids[i] = id
		i++
	}
	sort.Ints(ids)
	list := make([]model.Book, len(ids))
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
		return model.ErrNotFound
	}
	delete(db.books, id)
	return nil
}

// GetOwnershipList returns book list by user ID
func (db *Local) GetOwnershipList(userID int) ([]*model.Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, model.ErrNotFound
	}

	var list = []*model.Ownership{}
	for i := range l {
		list = append(
			list,
			&model.Ownership{
				UserID: int64(userID),
				BookID: l[i].ID,
				Book:   l[i],
			},
		)
	}
	return list, nil
}

// GetOwnership returns user book association
func (db *Local) GetOwnership(userID, bookID int) (*model.Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, model.ErrNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			return &model.Ownership{UserID: int64(userID), BookID: int64(bookID), Book: l[i]}, nil
		}
	}
	return nil, model.ErrNotFound
}

// InsertOwnership inserts user book association
func (db *Local) InsertOwnership(userID, bookID int) (*model.Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, model.ErrNotFound
	}

	b, ok := db.books[bookID]
	if !ok {
		return nil, model.ErrNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			return nil, model.ErrDuplicateKey
		}
	}

	db.ownerships[userID] = append(l, b)
	return &model.Ownership{UserID: int64(userID), BookID: int64(bookID), Book: b}, nil
}

//UpdateOwnership update the ownership
func (db *Local) UpdateOwnership(userID, bookID int) error {
	return nil
}

// DeleteOwnership deletes user book association
func (db *Local) DeleteOwnership(userID, bookID int) error {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return model.ErrNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			db.ownerships[userID] = append(l[:i], l[i+1:]...)
			return nil
		}
	}
	return model.ErrNotFound
}
