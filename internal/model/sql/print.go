package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getPrint(query string, params ...interface{}) (*model.Print, error) {
	var v = model.Print{}
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

func (m *Model) GetPrintByID(id int) (*model.Print, error) {
	return m.getPrint(`SELECT id, name FROM print where id = ?`, id)
}

func (m *Model) GetPrintByName(name string) (*model.Print, error) {
	return m.getPrint(`SELECT id, name FROM print where name = ?`, name)
}

func (m *Model) ListPrints() ([]*model.Print, error) {
	rows, err := m.db.Query(`SELECT id, name FROM print`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Print{}
	for rows.Next() {
		v := model.Print{}
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

func (m *Model) InsertPrint(name string) (*model.Print, error) {
	_, err := m.GetPrintByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO print (id, name, create_ts, update_ts)
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
	return &model.Print{ID: id, Name: name}, nil
}

func (m *Model) UpdatePrint(id int, name string) error {
	return m.exec(
		`UPDATE print SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

func (m *Model) DeletePrint(id int) error {
	return m.exec(`DELETE FROM print where id = ?`, id)
}
