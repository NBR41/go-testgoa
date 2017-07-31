package appmodel

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Model struct for model
type Model struct {
	db *sql.DB
}

// NewModel returns new instance of model
func NewModel(connString string) (*Model, error) {
	db, err := sql.Open("mysql", connString+"?charset=utf8mb4,utf8")
	if err != nil {
		return nil, err
	}
	return &Model{db: db}, nil
}

// Close close the connextion
func (m *Model) Close() error {
	return m.db.Close()
}

func (m *Model) exec(query string, params ...interface{}) error {
	res, err := m.db.Exec(query, params...)
	if err != nil {
		return err
	}

	nb, err := res.RowsAffected()
	switch {
	case err != nil:
		return err
	case nb == 0:
		return ErrNotFound
	default:
		return nil
	}
}

func (m *Model) getUser(query string, params ...interface{}) (*User, error) {
	var u = User{}
	err := m.db.QueryRow(query, params...).Scan(
		&u.ID, &u.Nickname, &u.Email, &u.IsVerified, &u.IsAdmin,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &u, nil
	}
}

func (m *Model) getBook(query string, params ...interface{}) (*Book, error) {
	var b = Book{}
	err := m.db.QueryRow(query, params...).Scan(&b.ID, &b.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &b, nil
	}
}

// GetUserList returns user list
func (m *Model) GetUserList() ([]User, error) {
	rows, err := m.db.Query(`SELECT user_id, nickname, email, verified, admin FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []User{}
	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.IsVerified, &u.IsAdmin); err != nil {
			return nil, err
		}
		l = append(l, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// GetUserByID returns user by ID
func (m *Model) GetUserByID(id int) (*User, error) {
	return m.getUser(
		`SELECT user_id, nickname, email, activated, admin FROM users WHERE user_id = ?`,
		id,
	)
}

// GetUserByEmailOrNickname returns user by email or nickname
func (m *Model) GetUserByEmailOrNickname(email, nickname string) (*User, error) {
	return m.getUser(
		`SELECT user_id, nickname, email, activated, admin FROM users WHERE email = ? OR nickname = ?`,
		email, nickname,
	)
}

// GetUserByEmail returns user by email
func (m *Model) GetUserByEmail(email string) (*User, error) {
	return m.getUser(
		`SELECT user_id, nickname, email, activated, admin FROM users WHERE email = ?`,
		email,
	)
}

// GetUserByNickname returns user by email
func (m *Model) GetUserByNickname(nickname string) (*User, error) {
	return m.getUser(
		`SELECT user_id, nickname, email, activated, admin FROM users WHERE nickname = ?`,
		nickname,
	)
}

// GetAuthenticatedUser returns user if password matches email or nickname
func (m *Model) GetAuthenticatedUser(login, password string) (*User, error) {
	return m.getUser(
		`
SELECT user_id, nickname, email, activated, admin
FROM users
WHERE password = ? AND (email = ? OR nickname =?)`,
		password, login, login,
	)
}

// InsertUser insert user
func (m *Model) InsertUser(email, nickname, password string) (*User, error) {
	_, err := m.GetUserByEmailOrNickname(email, nickname)
	switch {
	case err != nil && err != ErrNotFound:
		return nil, err
	case err == nil:
		return nil, ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO users (user_id, nickname, email, password, activated, admin, create_ts, update_ts)
VALUES (null, ?, ?, ?, 0, 0, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		nickname, email, password,
	)
	if err != nil {
		return nil, err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &User{ID: id, Email: email, Nickname: nickname}, nil
}

// UpdateUserNickname updates user nickname by ID
func (m *Model) UpdateUserNickname(id int, nickname string) error {
	return m.exec(
		`UPDATE users set nickname = ?, update_ts = NOW() where user_id = ?`,
		nickname, id,
	)
}

// UpdateUserPassword updates user password by ID
func (m *Model) UpdateUserPassword(id int, password string) error {
	return m.exec(
		`UPDATE users set password = ?, update_ts = NOW() where user_id = ?`,
		password, id,
	)
}

// UpdateUserActivation update user activation by ID
func (m *Model) UpdateUserActivation(id int, activated bool) error {
	return m.exec(
		`UPDATE users set activated = ?, update_ts = NOW() where user_id = ?`,
		activated, id,
	)
}

// DeleteUser deletes user by ID
func (m *Model) DeleteUser(id int) error {
	return m.exec(`DELETE FROM users where user_id = ?`, id)
}

// InsertBook inserts book
func (m *Model) InsertBook(isbn, name string) (*Book, error) {
	_, err := m.GetBookByName(name)
	switch {
	case err != nil && err != ErrNotFound:
		return nil, err
	case err == nil:
		return nil, ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO books (book_id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		name,
	)
	if err != nil {
		return nil, err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &Book{ID: id, Name: name}, nil
}

// GetBookByID returns book by ID
func (m *Model) GetBookByID(id int) (*Book, error) {
	return m.getBook(`SELECT book_id, name from books where id = ?`, id)
}

// GetBookByID returns book by ID
func (m *Model) GetBookByISBN(isbn string) (*Book, error) {
	return m.getBook(`SELECT book_id, name from books where id = ?`, isbn)
}

// GetBookByName returns book by name
func (m *Model) GetBookByName(name string) (*Book, error) {
	return m.getBook(`SELECT book_id, name from books where name = ?`, name)
}

// GetBookList returns book list
func (m *Model) GetBookList() ([]Book, error) {
	rows, err := m.db.Query(`SELECT book_id, name FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []Book{}
	for rows.Next() {
		b := Book{}
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			return nil, err
		}
		l = append(l, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// UpdateBook update book infos
func (m *Model) UpdateBook(id int, name string) error {
	return m.exec(
		`UPDATE books set name = ?, update_ts = NOW() where book_id = ?`,
		name, id,
	)
}

// DeleteBook delete book by ID
func (m *Model) DeleteBook(id int) error {
	return m.exec(`DELETE FROM books where book_id = ?`, id)
}

// GetOwnershipList returns book list by user ID
func (m *Model) GetOwnershipList(userID int) ([]Ownership, error) {
	rows, err := m.db.Query(
		`SELECT b.book_id, b.name FROM ownerships u JOIN books b USING(book_id) where user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []Ownership{}
	for rows.Next() {
		b := Book{}
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			return nil, err
		}
		l = append(
			l,
			Ownership{
				UserID: int64(userID),
				BookID: b.ID,
				Book:   &b,
			},
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// GetOwnership returns user book association
func (m *Model) GetOwnership(userID, bookID int) (*Ownership, error) {
	b, err := m.getBook(
		`
SELECT b.book_id, b.name
FROM ownerships u
JOIN books b USING(book_id)
where u.user_id = ? and b.book_id = ?`,
		userID, bookID,
	)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, nil
	}
	return &Ownership{UserID: int64(userID), BookID: int64(bookID), Book: b}, nil
}

// InsertOwnership inserts user book association
func (m *Model) InsertOwnership(userID, bookID int) (*Ownership, error) {
	_, err := m.db.Exec(
		`
INSERT INTO ownerships (user_id, book_id, create_ts, update_ts)
VALUES (?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		userID, bookID,
	)
	if err != nil {
		return nil, err
	}
	return &Ownership{UserID: int64(userID), BookID: int64(bookID)}, nil
}

func (m *Model) UpdateOwnership(userID, bookID int) error {
	return m.exec(
		`UPDATE ownerships set update_ts = NOW() where user_id = ? and book_id = ?`,
		userID, bookID,
	)
}

// DeleteOwnership deletes user book association
func (m *Model) DeleteOwnership(userID, bookID int) error {
	return m.exec(`DELETE FROM ownerships where user_id = ? and book_id = ?`, userID, bookID)
}
