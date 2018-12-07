package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) getGenre(query string, params ...interface{}) (*model.Genre, error) {
	var v = model.Genre{}
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

func (m *Model) GetGenreByID(id int) (*model.Genre, error) {
	return m.getGenre(`SELECT id, name FROM genre where id = ?`, id)
}

func (m *Model) GetGenreByName(name string) (*model.Genre, error) {
	return m.getGenre(`SELECT id, name FROM genre where name = ?`, name)
}

func (m *Model) GetGenreList() ([]*model.Genre, error) {
	rows, err := m.db.Query(`SELECT id, name FROM genre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Genre{}
	for rows.Next() {
		v := model.Genre{}
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

func (m *Model) InsertGenre(name string) (*model.Genre, error) {
	_, err := m.GetGenreByName(name)
	switch {
	case err != nil && err != model.ErrNotFound:
		return nil, err
	case err == nil:
		return nil, model.ErrDuplicateKey
	}
	res, err := m.db.Exec(
		`
INSERT INTO genre (id, name, create_ts, update_ts)
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
	return &model.Genre{ID: id, Name: name}, nil
}

func (m *Model) UpdateGenre(id int, name string) error {
	return m.exec(
		`UPDATE genre SET name = ?, update_ts = NOW() WHERE id = ?`,
		name, id,
	)
}

func (m *Model) DeleteGenre(id int) error {
	return m.exec(`DELETE FROM genre where id = ?`, id)
}
