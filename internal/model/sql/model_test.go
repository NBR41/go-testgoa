package sql

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var escaper = strings.NewReplacer(`?`, `\?`, `(`, `\(`, `)`, `\)`)

func escapeQuery(qry string) string {
	return escaper.Replace(qry)
}

func TestNew(t *testing.T) {
	m, err := New(GetConnGetter(""), nil)
	if m != nil {
		t.Error("unpected value")
	}
	if err == nil {
		t.Error("expecting error")
	}

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	m, err = New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)
	if m == nil {
		t.Error("expecting value")
	} else {
		m.Close()
	}
	if err != nil {
		t.Error("unexpected error")
	}
}
