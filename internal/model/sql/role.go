package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getRole(query string, params ...interface{}) (*model.Role, error) {
	var v = model.Role{}
	err := m.db.QueryRow(query, params...).Scan(&v.ID, &v.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &v, nil
	}
}

func (m *Model) listRoles(query string, params ...interface{}) ([]*model.Role, error) {
	rows, err := m.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Role{}
	for rows.Next() {
		v := model.Role{}
		if err := rows.Scan(&v.ID, &v.Name); err != nil {
			return nil, err
		}
		l = append(l, &v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// GetRoleByID returns role by ID
func (m *Model) GetRoleByID(id int) (*model.Role, error) {
	return m.getRole(`SELECT id, name FROM role where id = ?`, id)
}

// GetRoleByName returns role by name
func (m *Model) GetRoleByName(name string) (*model.Role, error) {
	return m.getRole(`SELECT id, name FROM role where name = ?`, name)
}

// ListRoles list roles
func (m *Model) ListRoles() ([]*model.Role, error) {
	return m.listRoles(`SELECT id, name FROM role`)
}

func (m *Model) ListRolesByAuthorID(authorID int) ([]*model.Role, error) {
	if _, err := m.GetAuthorByID(authorID); err != nil {
		return nil, err
	}
	return m.listRoles(`SELECT r.id, r.name FROM role r JOIN authorship a ON (a.role_id = r.id) where a.author_id = ?`, authorID)
}

// InsertRole inserts role
func (m *Model) InsertRole(name string) (*model.Role, error) {
	_, err := m.GetRoleByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO role (id, name, create_ts, update_ts)
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
	return &model.Role{ID: id, Name: name}, nil
}

// UpdateRole update role
func (m *Model) UpdateRole(id int, name string) error {
	return m.exec(
		`UPDATE role SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

// DeleteRole delete role
func (m *Model) DeleteRole(id int) error {
	return m.exec(`DELETE FROM role where id = ?`, id)
}
