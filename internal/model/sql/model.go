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
		return sql.Open("mysql", connString+"?charset=utf8mb4,utf8")
	}
}

// Model struct for model
type Model struct {
	fConn ConnGetter
	pass  model.Passworder
	db    *sql.DB
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
		return err
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

func filterError(err error) error {
	if strings.HasPrefix(err.Error(), "ERROR 1452") {
		return model.ErrNotFound
	}
	if strings.HasPrefix(err.Error(), "ERROR 1062") {
		return model.ErrDuplicateKey
	}
	return err
}
