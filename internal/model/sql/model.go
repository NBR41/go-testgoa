package sql

import (
	"database/sql"
	"strings"

	// enable mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/NBR41/go-testgoa/internal/model"
)

//ConnGetter type for a conn getter func
type ConnGetter func() (*sql.DB, error)

//GetConnGetter return a connection getter func
func GetConnGetter(connString string) ConnGetter {
	return func() (*sql.DB, error) {
		db, err := sql.Open("mysql", connString+"?charset=utf8mb4,utf8")
		if err != nil {
			return nil, err
		}
		if err = db.Ping(); err != nil {
			db.Close()
			return nil, err
		}
		return db, nil
	}
}

// Model struct for model
type Model struct {
	pass model.Passworder
	db   *sql.DB
}

// New returns new instance of model
func New(f ConnGetter, pass model.Passworder) (*Model, error) {
	db, err := f()
	if err != nil {
		return nil, err
	}
	return &Model{db: db, pass: pass}, nil
}

// Close close the connextion
func (m *Model) Close() error {
	return m.db.Close()
}

func (m *Model) exec(query string, params ...interface{}) error {
	res, err := m.db.Exec(query, params...)
	if err != nil {
		return filterError(err)
	}
	nb, err := res.RowsAffected()
	switch {
	case err != nil:
		return filterError(err)
	case nb == 0:
		return model.ErrNotFound
	default:
		return nil
	}
}

const (
	duplicateErr = "Error 1062"
	fkErr        = "Error 1452"
)

func filterError(err error) error {
	if strings.HasPrefix(err.Error(), duplicateErr) {
		switch {
		case strings.Contains(err.Error(), "email"):
			return model.ErrDuplicateEmail
		case strings.Contains(err.Error(), "nickname"):
			return model.ErrDuplicateNickname
		default:
			return model.ErrDuplicateKey
		}
	}
	if strings.HasPrefix(err.Error(), fkErr) {
		return model.ErrInvalidID
	}
	return err
}
