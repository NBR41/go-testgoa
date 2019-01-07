package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getClass(query string, params ...interface{}) (*model.Class, error) {
	var v = model.Class{}
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

func (m *Model) GetClassByID(id int) (*model.Class, error) {
	return m.getClass(`SELECT id, name FROM class where id = ?`, id)
}

func (m *Model) GetClassByName(name string) (*model.Class, error) {
	return m.getClass(`SELECT id, name FROM class where name = ?`, name)
}

func (m *Model) ListClasses() ([]*model.Class, error) {
	rows, err := m.db.Query(`SELECT id, name FROM class`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Class{}
	for rows.Next() {
		v := model.Class{}
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

func (m *Model) InsertClass(name string) (*model.Class, error) {
	_, err := m.GetClassByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO class (id, name, create_ts, update_ts)
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
	return &model.Class{ID: id, Name: name}, nil
}

func (m *Model) UpdateClass(id int, name string) error {
	return m.exec(
		`UPDATE class SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

func (m *Model) DeleteClass(id int) error {
	return m.exec(`DELETE FROM class where id = ?`, id)
}
