package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getUser(query string, params ...interface{}) (*model.User, error) {
	var u = model.User{}
	err := m.db.QueryRow(query, params...).Scan(
		&u.ID, &u.Nickname, &u.Email, &u.IsValidated, &u.IsAdmin,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &u, nil
	}
}

// GetUserList returns user list
func (m *Model) GetUserList() ([]model.User, error) {
	rows, err := m.db.Query(`SELECT id, nickname, email, verified, admin FROM user`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []model.User{}
	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.IsValidated, &u.IsAdmin); err != nil {
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
func (m *Model) GetUserByID(id int) (*model.User, error) {
	return m.getUser(
		`SELECT id, nickname, email, activated, admin FROM user WHERE id = ?`,
		id,
	)
}

// GetUserByEmailOrNickname returns user by email or nickname
func (m *Model) GetUserByEmailOrNickname(email, nickname string) (*model.User, error) {
	return m.getUser(
		`SELECT id, nickname, email, activated, admin FROM user WHERE email = ? OR nickname = ?`,
		email, nickname,
	)
}

// GetUserByEmail returns user by email
func (m *Model) GetUserByEmail(email string) (*model.User, error) {
	return m.getUser(
		`SELECT id, nickname, email, activated, admin FROM user WHERE email = ?`,
		email,
	)
}

// GetUserByNickname returns user by email
func (m *Model) GetUserByNickname(nickname string) (*model.User, error) {
	return m.getUser(
		`SELECT id, nickname, email, activated, admin FROM user WHERE nickname = ?`,
		nickname,
	)
}

// GetAuthenticatedUser returns user if password matches email or nickname
func (m *Model) GetAuthenticatedUser(login, password string) (*model.User, error) {
	var u = model.User{}
	var salt, hash []byte
	err := m.db.QueryRow(
		`
SELECT id, nickname, email, activated, admin, salt, password
FROM user
WHERE email = ? OR nickname =?`,
		login, login,
	).Scan(
		&u.ID, &u.Nickname, &u.Email, &u.IsValidated, &u.IsAdmin, &salt, &hash,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	}

	ok, err := m.pass.ComparePassword(password, salt, hash)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, model.ErrInvalidCredentials
	}
	return &u, nil
}

// InsertUser insert user
func (m *Model) InsertUser(email, nickname, password string) (*model.User, error) {
	_, err := m.GetUserByEmailOrNickname(email, nickname)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}

	salt, hash, err := m.pass.CryptPassword(password)
	if err != nil {
		return nil, err
	}

	res, err := m.db.Exec(
		`
INSERT INTO user (id, nickname, email, salt, password, activated, admin, create_ts, update_ts)
VALUES (null, ?, ?, ?, ?, 0, 0, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		nickname, email, salt, hash,
	)
	if err != nil {
		return nil, err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.User{ID: id, Email: email, Nickname: nickname}, nil
}

// UpdateUserNickname updates user nickname by ID
func (m *Model) UpdateUserNickname(id int, nickname string) error {
	return m.exec(
		`UPDATE user set nickname = ?, update_ts = NOW() where id = ?`,
		nickname, id,
	)
}

// UpdateUserPassword updates user password by ID
func (m *Model) UpdateUserPassword(id int, password string) error {
	salt, hash, err := m.pass.CryptPassword(password)
	if err != nil {
		return err
	}
	return m.exec(
		`UPDATE user set salt = ?, password = ?, update_ts = NOW() where id = ?`,
		salt, hash, id,
	)
}

// UpdateUserActivation update user activation by ID
func (m *Model) UpdateUserActivation(id int, activated bool) error {
	return m.exec(
		`UPDATE user set activated = ?, update_ts = NOW() where id = ?`,
		activated, id,
	)
}

// DeleteUser deletes user by ID
func (m *Model) DeleteUser(id int) error {
	return m.exec(`DELETE FROM user where id = ?`, id)
}
