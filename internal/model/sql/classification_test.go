package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestGetClassification(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetClassification)
	mock.ExpectQuery(qry).WithArgs(1, 2).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(1, 2).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
	mock.ExpectQuery(qry).WithArgs(1, 2).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(qry).WithArgs(1, 2).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(1, 2).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Class
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Class{ID: 1, Name: "foo"}, nil},
	}

	for i := range tests {
		v, err := m.GetClassification(1, 2)
		if err != nil {
			if tests[i].err == nil {
				t.Errorf("unexpected error for [%s], [%v]", tests[i].desc, err)
				continue
			}
			if tests[i].err.Error() != err.Error() {
				t.Errorf("unexpected error for [%s], exp [%v] got [%v]", tests[i].desc, tests[i].err, err)
				continue
			}
			continue
		}
		if tests[i].err != nil {
			t.Errorf("expecting error for [%s]", tests[i].desc)
		}
		if diff := pretty.Compare(v, tests[i].exp); diff != "" {
			t.Errorf("unexpected value for [%s]\n%s", tests[i].desc, diff)
		}
	}
}

func TestInsertClassification(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := escapeQuery(qryInsertClassification)
	mock.ExpectExec(qry).WithArgs(1, 2).WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(qry).WithArgs(1, 2).WillReturnError(errors.New(fkErr))
	mock.ExpectExec(qry).WithArgs(1, 2).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(2, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Class
		err  error
	}{
		{"duplicate", nil, model.ErrDuplicateKey},
		{"invalid id", nil, model.ErrInvalidID},
		{"query error", nil, errors.New("query error")},
		{"valid", &model.Class{ID: 2}, nil},
	}

	for i := range tests {
		v, err := m.InsertClassification(1, 2)
		if err != nil {
			if tests[i].err == nil {
				t.Errorf("unexpected error for [%s], [%v]", tests[i].desc, err)
				continue
			}
			if tests[i].err.Error() != err.Error() {
				t.Errorf("unexpected error for [%s], exp [%v] got [%v]", tests[i].desc, tests[i].err, err)
				continue
			}
			continue
		}
		if tests[i].err != nil {
			t.Errorf("expecting error for [%s]", tests[i].desc)
		}
		if diff := pretty.Compare(v, tests[i].exp); diff != "" {
			t.Errorf("unexpected value for [%s]\n%s", tests[i].desc, diff)
		}
	}
}

func TestDeleteClassification(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryDeleteClassification)
	mock.ExpectExec(qry).WithArgs(1, 2).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(1, 2).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(0, 1))
	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)
	tests := []struct {
		desc string
		err  error
	}{
		{"query error", errors.New("query error")},
		{"result error", errors.New("result error")},
		{"no rows", model.ErrNotFound},
		{"valid", nil},
	}
	for i := range tests {
		err := m.DeleteClassification(1, 2)
		if err != nil {
			if tests[i].err == nil {
				t.Errorf("unexpected error for [%s], [%v]", tests[i].desc, err)
				continue
			}
			if tests[i].err.Error() != err.Error() {
				t.Errorf("unexpected error for [%s], exp [%v] got [%v]", tests[i].desc, tests[i].err, err)
				continue
			}
			continue
		}
		if tests[i].err != nil {
			t.Errorf("expecting error for [%s]", tests[i].desc)
		}
	}
}
