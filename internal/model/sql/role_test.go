package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestInsertRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := escapeQuery(qryInsertRole)
	mock.ExpectExec(qry).WithArgs("foo").WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Role
		err  error
	}{
		{"duplicate", nil, model.ErrDuplicateKey},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Role{ID: 123, Name: "foo"}, nil},
	}

	for i := range tests {
		v, err := m.InsertRole("foo")
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

func TestGetRoleByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetRoleByID)
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
		exp  *model.Role
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Role{ID: 1, Name: "foo"}, nil},
	}

	for i := range tests {
		v, err := m.GetRoleByID(123)
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

func TestGetRoleByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetRoleByName)
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
		exp  *model.Role
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Role{ID: 1, Name: "foo"}, nil},
	}

	for i := range tests {
		v, err := m.GetRoleByName("foo")
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

func TestListRolesByIDs(t *testing.T) {
	var authorID int = 1
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery("SELECT DISTINCT role.id, role.name FROM role WHERE 1")
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))
	qry = escapeQuery("SELECT DISTINCT role.id, role.name FROM role JOIN authorship ON (authorship.role_id = role.id) WHERE 1 AND authorship.author_id = ?")
	mock.ExpectQuery(qry).WithArgs(authorID).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc   string
		params *int
		exp    []model.Role
		err    error
	}{
		{"query error", nil, nil, errors.New("query error")},
		{"scan conversion error", nil, nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, nil, errors.New("scan error")},
		{"valid", nil, []model.Role{model.Role{ID: 1, Name: "foo"}}, nil},
		{"valid with filter", &authorID, []model.Role{model.Role{ID: 1, Name: "foo"}}, nil},
	}

	for i := range tests {
		v, err := m.ListRolesByIDs(tests[i].params)
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

func TestUpdateRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryUpdateRole)
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
		err := m.UpdateRole(123, "foo")
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

func TestDeleteRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryDeleteRole)
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
		err := m.DeleteRole(123)
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
