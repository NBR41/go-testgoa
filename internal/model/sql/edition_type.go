package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getEditionType(query string, params ...interface{}) (*model.EditionType, error) {
	var v = model.EditionType{}
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

func (m *Model) GetEditionTypeByID(id int) (*model.EditionType, error) {
	return m.getEditionType(`SELECT id, name FROM edition_type where id = ?`, id)
}

func (m *Model) GetEditionTypeByName(name string) (*model.EditionType, error) {
	return m.getEditionType(`SELECT id, name FROM edition_type where name = ?`, name)
}

func (m *Model) GetEditionTypeList() ([]*model.EditionType, error) {
	rows, err := m.db.Query(`SELECT id, name FROM edition_type`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.EditionType{}
	for rows.Next() {
		v := model.EditionType{}
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

func (m *Model) InsertEditionType(name string) (*model.EditionType, error) {
	_, err := m.GetEditionTypeByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO edition_type (id, name, create_ts, update_ts)
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
	return &model.EditionType{ID: id, Name: name}, nil
}

func (m *Model) UpdateEditionType(id int, name string) error {
	return m.exec(
		`UPDATE edition_type SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

func (m *Model) DeleteEditionType(id int) error {
	return m.exec(`DELETE FROM edition_type where id = ?`, id)
}
