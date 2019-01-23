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

	qry := escapeQuery(qryInsertBook)
	mock.ExpectExec(qry).WithArgs("foo", "bar", 5).WillReturnError(errors.New(fkErr))
	mock.ExpectExec(qry).WithArgs("foo", "bar", 5).WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(qry).WithArgs("foo", "bar", 5).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo", "bar", 5).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs("foo", "bar", 5).WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Book
		err  error
	}{
		{"foreign key not found", nil, model.ErrInvalidID},
		{"duplicate", nil, model.ErrDuplicateKey},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Book{ID: 123, ISBN: "foo", Name: "bar", SeriesID: 5}, nil},
	}

	for i := range tests {
		v, err := m.InsertBook("foo", "bar", 5)
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
	qry := escapeQuery(qryGetBookByID)
	mock.ExpectQuery(qry).WithArgs(123).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow("foo", "bar", "baz", "quux"))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2).RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2))

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
		{"valid", &model.Book{ID: 1, ISBN: "foo", Name: "bar", SeriesID: 2}, nil},
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
	qry := escapeQuery(qryGetBookByISBN)
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2).RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2))

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
		{"valid", &model.Book{ID: 1, ISBN: "foo", Name: "bar", SeriesID: 2}, nil},
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
	qry := escapeQuery(qryGetBookByName)
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow("foo", "bar", "baz", "qux"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2).RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2))

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
		{"valid", &model.Book{ID: 1, ISBN: "foo", Name: "bar", SeriesID: 2}, nil},
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

func TestListBooksByIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	var collectionID, editorID, printID, seriesID int = 1, 2, 3, 4
	qry := escapeQuery(`SELECT DISTINCT book.id, book.isbn, book.name, book.series_id FROM book WHERE 1`)
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2).RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2))
	qry = escapeQuery(`SELECT DISTINCT book.id, book.isbn, book.name, book.series_id FROM book JOIN edition ON (edition.book_id = book.id) JOIN collection ON (edition.collection_id = collection.id) WHERE 1 AND edition.collection_id = ? AND collection.editor_id = ? AND edition.print_id = ? AND book.series_id = ?`)
	mock.ExpectQuery(qry).WithArgs(collectionID, editorID, printID, seriesID).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name", "series_id"}).AddRow(1, "foo", "bar", 2))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc   string
		params []*int
		exp    []model.Book
		err    error
	}{
		{"query error", []*int{nil, nil, nil, nil}, nil, errors.New("query error")},
		{"scan conversion error", []*int{nil, nil, nil, nil}, nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", []*int{nil, nil, nil, nil}, nil, errors.New("scan error")},
		{"valid", []*int{nil, nil, nil, nil}, []model.Book{model.Book{ID: 1, ISBN: "foo", Name: "bar", SeriesID: 2}}, nil},
		{"valid with filter", []*int{&collectionID, &editorID, &printID, &seriesID}, []model.Book{model.Book{ID: 1, ISBN: "foo", Name: "bar", SeriesID: 2}}, nil},
	}

	for i := range tests {
		v, err := m.ListBooksByIDs(tests[i].params[0], tests[i].params[1], tests[i].params[2], tests[i].params[3])
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

func TestUpdateBook0(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)
	err = m.UpdateBook(123, nil, nil)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
}

func TestUpdateBook1(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	n1 := "foo"
	qry := `UPDATE book SET name = \?, update_ts = NOW\(\) WHERE id = \?`
	mock.ExpectExec(qry).WithArgs(n1, 123).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(n1, 123).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs(n1, 123).WillReturnResult(sqlmock.NewResult(0, 1))
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
		err := m.UpdateBook(123, &n1, nil)
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

func TestUpdateBook2(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	series := 456
	qry := `UPDATE book SET series_id = \?, update_ts = NOW\(\) WHERE id = \?`
	mock.ExpectExec(qry).WithArgs(456, 123).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(456, 123).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs(456, 123).WillReturnResult(sqlmock.NewResult(0, 1))
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
		err := m.UpdateBook(123, nil, &series)
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

func TestUpdateBook4(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	name := "foo"
	series := 456
	qry := escapeQuery(`UPDATE book SET name = ?, series_id = ?, update_ts = NOW() WHERE id = ?`)
	mock.ExpectExec(qry).WithArgs("foo", 456, 123).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo", 456, 123).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs("foo", 456, 123).WillReturnResult(sqlmock.NewResult(0, 1))
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
		err := m.UpdateBook(123, &name, &series)
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
	qry := escapeQuery(qryDeleteBook)
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
