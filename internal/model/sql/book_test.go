package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestInsertBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := `
INSERT INTO book \(id, isbn, name, create_ts, update_ts\)
VALUES \(null, \?, \?, NOW\(\), NOW\(\)\)
ON DUPLICATE KEY UPDATE update_ts = VALUES\(update_ts\)`

	isbnqry := `SELECT id, isbn, name from book where isbn = \?`
	mock.
		ExpectQuery(isbnqry).
		WithArgs("foo").
		WillReturnError(errors.New("duplicate error"))
	mock.ExpectQuery(isbnqry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar"))
	mock.ExpectQuery(isbnqry).WithArgs("foo").WillReturnError(model.ErrNotFound)
	mock.ExpectExec(qry).WithArgs("foo", "bar").WillReturnError(errors.New("query error"))
	mock.ExpectQuery(isbnqry).WithArgs("foo").WillReturnError(model.ErrNotFound)
	mock.ExpectExec(qry).WithArgs("foo", "bar").WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectQuery(isbnqry).WithArgs("foo").WillReturnError(model.ErrNotFound)
	mock.ExpectExec(qry).WithArgs("foo", "bar").WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Book
		err  error
	}{
		{"duplicate error", nil, errors.New("duplicate error")},
		{"duplicate", nil, model.ErrDuplicateKey},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Book{ID: 123, ISBN: "foo", Name: "bar"}, nil},
	}

	for i := range tests {
		v, err := m.InsertBook("foo", "bar")
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
func TestGetBookByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, isbn, name from book where id = \?`
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}))
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow("foo", "bar", "baz"))
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar").RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Book
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Book{ID: 1, ISBN: "foo", Name: "bar"}, nil},
	}

	for i := range tests {
		v, err := m.GetBookByID(123)
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

func TestGetBookByISBN(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, isbn, name from book where isbn = \?`
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow("foo", "bar", "baz"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar").RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Book
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Book{ID: 1, ISBN: "foo", Name: "bar"}, nil},
	}

	for i := range tests {
		v, err := m.GetBookByISBN("foo")
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

func TestGetBookByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, isbn, name from book where name = \?`
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow("foo", "bar", "baz"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar").RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Book
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Book{ID: 1, ISBN: "foo", Name: "bar"}, nil},
	}

	for i := range tests {
		v, err := m.GetBookByName("foo")
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

func TestGetBookList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, isbn, name FROM book`
	mock.
		ExpectQuery(qry).
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow("foo", "bar", "baz"))
	mock.
		ExpectQuery(qry).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar").RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(1, "foo", "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Book
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Book{model.Book{ID: 1, ISBN: "foo", Name: "bar"}}, nil},
	}

	for i := range tests {
		v, err := m.ListBooks()
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

func TestUpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `UPDATE book set name = \?, update_ts = NOW\(\) where id = \?`
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
		err := m.UpdateBook(123, "foo")
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
func TestDeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `DELETE FROM book where id = \?`
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
		err := m.DeleteBook(123)
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
