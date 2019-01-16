package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestGetEditionByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetEditionByID)
	mock.ExpectQuery(qry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"eid", "bid", "bname", "cid", "cname", "pid", "pname"}))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"eid", "bid", "bname", "cid", "cname", "pid", "pname"}).AddRow("foo", "abar", "abaz", "rbar", "rbaz", "bbar", "bbaz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"eid", "bid", "bname", "cid", "cname", "pid", "pname"}).AddRow(1, 2, "foo", 3, "bar", 4, "baz").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"eid", "bid", "bname", "cid", "cname", "pid", "pname"}).AddRow(1, 2, "foo", 3, "bar", 4, "baz"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Edition
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "eid": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Edition{ID: 1, BookID: 2, CollectionID: 3, PrintID: 4, Book: &model.Book{ID: 2, Name: "foo"}, Collection: &model.Collection{ID: 3, Name: "bar"}, Print: &model.Print{ID: 4, Name: "baz"}}, nil},
	}

	for i := range tests {
		v, err := m.GetEditionByID(1)
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

func TestListEditions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListEditions)
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"eid", "bid", "bname", "cid", "cname", "pid", "pname"}))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"eid", "bid", "bname", "cid", "cname", "pid", "pname"}).AddRow("foo", "abar", "abaz", "rbar", "rbaz", "bbar", "bbaz"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"eid", "bid", "bname", "cid", "cname", "pid", "pname"}).AddRow(1, 2, "foo", 3, "bar", 4, "baz").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"eid", "bid", "bname", "cid", "cname", "pid", "pname"}).AddRow(1, 2, "foo", 3, "bar", 4, "baz"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Edition
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", []*model.Edition{}, nil},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "eid": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Edition{&model.Edition{ID: 1, BookID: 2, CollectionID: 3, PrintID: 4, Book: &model.Book{ID: 2, Name: "foo"}, Collection: &model.Collection{ID: 3, Name: "bar"}, Print: &model.Print{ID: 4, Name: "baz"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListEditions()
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

func TestInsertEdition(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := escapeQuery(qryInsertEdition)
	mock.ExpectExec(qry).WithArgs(2, 3, 4).WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(qry).WithArgs(2, 3, 4).WillReturnError(errors.New(fkErr))
	mock.ExpectExec(qry).WithArgs(2, 3, 4).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(2, 3, 4).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs(2, 3, 4).WillReturnResult(sqlmock.NewResult(1, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Edition
		err  error
	}{
		{"duplicate error", nil, model.ErrDuplicateKey},
		{"duplicate error", nil, model.ErrInvalidID},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Edition{ID: 1, BookID: 2, CollectionID: 3, PrintID: 4}, nil},
	}

	for i := range tests {
		v, err := m.InsertEdition(2, 3, 4)
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

func TestDeleteEdition(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryDeleteEdition)
	mock.ExpectExec(qry).WithArgs(123).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(123).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs(123).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs(123).WillReturnResult(sqlmock.NewResult(0, 1))
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
		err := m.DeleteEdition(123)
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
