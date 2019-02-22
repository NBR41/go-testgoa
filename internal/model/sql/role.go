package sql

import (
	"database/sql"
	"strings"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetRoleByID   = `SELECT id, name FROM role WHERE id = ?`
	qryGetRoleByName = `SELECT id, name FROM role WHERE name = ?`
	qryInsertRole    = `
INSERT INTO role (id, name, create_ts, update_ts)
VALUES (null, ?, NOW(), NOW())`
	qryUpdateRole = `UPDATE role SET name = ?, update_ts = NOW() WHERE id = ?`
	qryDeleteRole = `DELETE FROM role WHERE id = ?`
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
	return m.getRole(qryGetRoleByID, id)
}

// GetRoleByName returns role by name
func (m *Model) GetRoleByName(name string) (*model.Role, error) {
	return m.getRole(qryGetRoleByName, name)
}

// ListRolesByIDs list roles by author id
func (m *Model) ListRolesByIDs(authorID *int) ([]*model.Role, error) {
	qry := `SELECT DISTINCT role.id, role.name FROM role`
	where := []string{"1"}
	vals := []interface{}{}
	if authorID != nil {
		qry += ` JOIN authorship ON (authorship.role_id = role.id)`
		where = append(where, `authorship.author_id = ?`)
		vals = append(vals, *authorID)
	}
	return m.listRoles(qry+` WHERE `+strings.Join(where, " AND "), vals...)
}

// InsertRole inserts role
func (m *Model) InsertRole(name string) (*model.Role, error) {
	res, err := m.db.Exec(qryInsertRole, name)
	if err != nil {
		return nil, filterError(err)
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
	return m.exec(qryUpdateRole, name, id)
}

// DeleteRole delete role
func (m *Model) DeleteRole(id int) error {
	return m.exec(qryDeleteRole, id)
}

func (m *Model) getOrInsertRole(name string) (*model.Role, error) {
	role, err := m.GetRoleByName(name)
	if err == model.ErrNotFound {
		role, err = m.InsertRole(name)
	}
	if err != nil {
		return nil, err
	}
	return role, nil
}
