package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestGetAuthorshipByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetAuthorshipByID)
	mock.ExpectQuery(qry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "aid", "aname", "rid", "rname", "bid", "bname"}))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "aid", "aname", "rid", "rname", "bid", "bname"}).AddRow("foo", "abar", "abaz", "rbar", "rbaz", "bbar", "bbaz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "aid", "aname", "rid", "rname", "bid", "bname"}).AddRow(1, 2, "foo", 3, "bar", 4, "baz").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "aid", "aname", "rid", "rname", "bid", "bname"}).AddRow(1, 2, "foo", 3, "bar", 4, "baz"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Authorship
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Authorship{ID: 1, AuthorID: 2, RoleID: 3, BookID: 4, Author: &model.Author{ID: 2, Name: "foo"}, Role: &model.Role{ID: 3, Name: "bar"}, Book: &model.Book{ID: 4, Name: "baz"}}, nil},
	}

	for i := range tests {
		v, err := m.GetAuthorshipByID(1)
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

func TestListAuthorships(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListAuthorships)
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "aid", "aname", "rid", "rname", "bid", "bname"}))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "aid", "aname", "rid", "rname", "bid", "bname"}).AddRow("foo", "abar", "abaz", "rbar", "rbaz", "bbar", "bbaz"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "aid", "aname", "rid", "rname", "bid", "bname"}).AddRow(1, 2, "foo", 3, "bar", 4, "baz").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "aid", "aname", "rid", "rname", "bid", "bname"}).AddRow(1, 2, "foo", 3, "bar", 4, "baz"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Authorship
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", []*model.Authorship{}, nil},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Authorship{&model.Authorship{ID: 1, AuthorID: 2, RoleID: 3, BookID: 4, Author: &model.Author{ID: 2, Name: "foo"}, Role: &model.Role{ID: 3, Name: "bar"}, Book: &model.Book{ID: 4, Name: "baz"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListAuthorships()
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

func TestInsertAuthorship(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := escapeQuery(qryInsertAuthorship)
	mock.ExpectExec(qry).WithArgs(2, 4, 3).WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(qry).WithArgs(2, 4, 3).WillReturnError(errors.New(fkErr))
	mock.ExpectExec(qry).WithArgs(2, 4, 3).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(2, 4, 3).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs(2, 4, 3).WillReturnResult(sqlmock.NewResult(1, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Authorship
		err  error
	}{
		{"duplicate error", nil, model.ErrDuplicateKey},
		{"duplicate error", nil, model.ErrInvalidID},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Authorship{ID: 1, AuthorID: 2, RoleID: 3, BookID: 4}, nil},
	}

	for i := range tests {
		v, err := m.InsertAuthorship(2, 4, 3)
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

func TestDeleteAuthorship(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryDeleteAuthorship)
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
		err := m.DeleteAuthorship(123)
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
