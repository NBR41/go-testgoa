package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestListOwnershipsByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListOwnershipsByUserID)
	uQry := escapeQuery(qryGetUserByID)
	mock.ExpectQuery(uQry).WithArgs(123).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(uQry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin"}))
	mock.ExpectQuery(uQry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin"}).AddRow(2, "bar", "qux", 1, 1))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(uQry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin"}).AddRow(2, "bar", "qux", 1, 1))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}))
	mock.ExpectQuery(uQry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin"}).AddRow(2, "bar", "qux", 1, 1))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow("foo", "bar", "baz"))
	mock.ExpectQuery(uQry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin"}).AddRow(2, "bar", "qux", 1, 1))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(456, "foo", "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(uQry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin"}).AddRow(2, "bar", "qux", 1, 1))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "isbn", "name"}).AddRow(456, "foo", "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Ownership
		err  error
	}{
		{"query error on user", nil, errors.New("query error")},
		{"user not found", nil, model.ErrNotFound},
		{"query error", nil, errors.New("query error")},
		{"no rows", []*model.Ownership{}, nil},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Ownership{&model.Ownership{UserID: 123, BookID: 456, Book: &model.Book{ID: 456, ISBN: "foo", Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListOwnershipsByUserID(123)
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
func TestGetOwnership(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetOwnership)
	mock.ExpectQuery(qry).WithArgs(123, 456).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(123, 456).WillReturnRows(sqlmock.NewRows([]string{"book_id", "isbn", "name", "series_id"}))
	mock.ExpectQuery(qry).WithArgs(123, 456).WillReturnRows(sqlmock.NewRows([]string{"book_id", "isbn", "name", "series_id"}).AddRow(456, "foo", "bar", 789))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)
	tests := []struct {
		desc string
		exp  *model.Ownership
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"not found", nil, model.ErrNotFound},
		{"valid", &model.Ownership{UserID: 123, BookID: 456, Book: &model.Book{ID: 456, ISBN: "foo", Name: "bar", SeriesID: 789}}, nil},
	}

	for i := range tests {
		v, err := m.GetOwnership(123, 456)
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

func TestInsertOwnership(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryInsertOwnership)
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnResult(sqlmock.NewResult(0, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)
	tests := []struct {
		desc string
		exp  *model.Ownership
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"valid", &model.Ownership{UserID: 123, BookID: 456}, nil},
	}

	for i := range tests {
		v, err := m.InsertOwnership(123, 456)
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
func TestUpdateOwnership(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryUpdateOwnership)
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnResult(sqlmock.NewResult(0, 1))
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
		err := m.UpdateOwnership(123, 456)
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

func TestDeleteOwnership(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryDeleteOwnership)
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs(123, 456).WillReturnResult(sqlmock.NewResult(0, 1))
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
		err := m.DeleteOwnership(123, 456)
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
