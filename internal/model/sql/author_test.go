package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestInsertAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := escapeQuery(qryInsertAuthor)
	mock.ExpectExec(qry).WithArgs("foo").WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Author
		err  error
	}{
		{"duplicate error", nil, model.ErrDuplicateKey},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Author{ID: 123, Name: "foo"}, nil},
	}

	for i := range tests {
		v, err := m.InsertAuthor("foo")
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

func TestGetAuthorByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetAuthorByID)
	mock.ExpectQuery(qry).WithArgs(123).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Author
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Author{ID: 1, Name: "foo"}, nil},
	}

	for i := range tests {
		v, err := m.GetAuthorByID(123)
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

func TestGetAuthorByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetAuthorByName)
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Author
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Author{ID: 1, Name: "foo"}, nil},
	}

	for i := range tests {
		v, err := m.GetAuthorByName("foo")
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

func TestListAuthors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListAuthors)
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Author
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Author{model.Author{ID: 1, Name: "foo"}}, nil},
	}

	for i := range tests {
		v, err := m.ListAuthors()
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

func TestListAuthorsByCategoryID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListAuthorsByCategoryID)
	mock.ExpectQuery(qry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Author
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Author{model.Author{ID: 1, Name: "foo"}}, nil},
	}

	for i := range tests {
		v, err := m.ListAuthorsByCategoryID(1)
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

func TestListAuthorsByRoleID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListAuthorsByRoleID)
	mock.ExpectQuery(qry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Author
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Author{model.Author{ID: 1, Name: "foo"}}, nil},
	}

	for i := range tests {
		v, err := m.ListAuthorsByRoleID(1)
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

func TestUpdateAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryUpdateAuthor)
	mock.ExpectExec(qry).WithArgs("foo", 123).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo", 123).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs("foo", 123).WillReturnResult(sqlmock.NewResult(0, 1))
	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)
	tests := []struct {
		desc string
		err  error
	}{
		{"query error", errors.New("query error")},
		{"no rows", model.ErrNotFound},
		{"valid", nil},
	}
	for i := range tests {
		err := m.UpdateAuthor(123, "foo")
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

func TestDeleteAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryDeleteAuthor)
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
		err := m.DeleteAuthor(123)
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
