package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getAuthor(query string, params ...interface{}) (*model.Author, error) {
	var v = model.Author{}
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

func (m *Model) GetAuthorByID(id int) (*model.Author, error) {
	return m.getAuthor(`SELECT id, name FROM author where id = ?`, id)
}

func (m *Model) GetAuthorByName(name string) (*model.Author, error) {
	return m.getAuthor(`SELECT id, name FROM author where name = ?`, name)
}

func (m *Model) GetAuthorList() ([]*model.Author, error) {
	rows, err := m.db.Query(`SELECT id, name FROM author`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Author{}
	for rows.Next() {
		v := model.Author{}
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

func (m *Model) InsertAuthor(name string) (*model.Author, error) {
	_, err := m.GetAuthorByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO author (id, name, create_ts, update_ts)
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
	return &model.Author{ID: id, Name: name}, nil
}

func (m *Model) UpdateAuthor(id int, name string) error {
	return m.exec(
		`UPDATE author SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

func (m *Model) DeleteAuthor(id int) error {
	return m.exec(`DELETE FROM author where id = ?`, id)
}
