package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestInsertClass(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := `
INSERT INTO class \(id, name, create_ts, update_ts\)
VALUES \(null, \?, NOW\(\), NOW\(\)\)
ON DUPLICATE KEY UPDATE update_ts = VALUES\(update_ts\)`

	mock.ExpectExec(qry).WithArgs("foo").WillReturnError(errors.New("ERROR 1062"))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs("foo").WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Class
		err  error
	}{
		{"duplicate", nil, model.ErrDuplicateKey},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Class{ID: 123, Name: "foo"}, nil},
	}

	for i := range tests {
		v, err := m.InsertClass("foo")
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

func TestGetClassByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, name FROM class where id = \?`
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
		v, err := m.GetClassByID(123)
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

func TestGetClassByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, name FROM class where name = \?`
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
		v, err := m.GetClassByName("foo")
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

func TestListClasses(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, name FROM class`
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Class
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Class{model.Class{ID: 1, Name: "foo"}}, nil},
	}

	for i := range tests {
		v, err := m.ListClasses()
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

func TestListClassesBySeriesID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `
SELECT distinct cl.id, cl.name
FROM class cl
JOIN classification c ON \(c.class_id = cl.id\)
WHERE c.series_id = \?`
	secQry := `SELECT s.id, s.name, c.id, c.name FROM series s JOIN category c ON \(c.id = s.category_id\) WHERE s.id = \?`
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"sid", "sname", "cid", "cname"}))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"sid", "sname", "cid", "cname"}).AddRow(1, "baz", 2, "quux"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"sid", "sname", "cid", "cname"}).AddRow(1, "baz", 2, "quux"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"sid", "sname", "cid", "cname"}).AddRow(1, "baz", 2, "quux"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"sid", "sname", "cid", "cname"}).AddRow(1, "baz", 2, "quux"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Class
		err  error
	}{
		{"series query error", nil, errors.New("query error")},
		{"series not found", nil, model.ErrNotFound},
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Class{model.Class{ID: 1, Name: "foo"}}, nil},
	}

	for i := range tests {
		v, err := m.ListClassesBySeriesID(1)
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

func TestListClassesByAuthorID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `
SELECT distinct cl.id, cl.name
FROM class cl
JOIN classification c ON \(c.class_id = cl.id\)
JOIN series s ON \(s.series_id = s.id\)
JOIN book b ON \(b.series_id = s.id\)
JOIN authorship a \(a.book_id = b.id\)
WHERE a.author_id = \?`
	secQry := `SELECT id, name FROM author where id = \?`
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "baz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "baz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "baz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "baz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Class
		err  error
	}{
		{"author query error", nil, errors.New("query error")},
		{"author not found", nil, model.ErrNotFound},
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Class{model.Class{ID: 1, Name: "foo"}}, nil},
	}

	for i := range tests {
		v, err := m.ListClassesByAuthorID(1)
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

func TestListClassesByCategoryID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `
SELECT distinct cl.id, cl.name
FROM class cl
JOIN classification c ON \(c.class_id = cl.id\)
JOIN series s ON \(s.series_id = s.id\)
WHERE s.category_id = \?`
	secQry := `SELECT id, name FROM category where id = \?`
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "baz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "baz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("foo", "bar"))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "baz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(secQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "baz"))
	mock.ExpectQuery(qry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Class
		err  error
	}{
		{"category query error", nil, errors.New("query error")},
		{"category not found", nil, model.ErrNotFound},
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Class{model.Class{ID: 1, Name: "foo"}}, nil},
	}

	for i := range tests {
		v, err := m.ListClassesByCategoryID(1)
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

func TestUpdateClass(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `UPDATE class SET name = \?, update_ts = NOW\(\) WHERE id = \?`
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
		err := m.UpdateClass(123, "foo")
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

func TestDeleteClass(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `DELETE FROM class where id = \?`
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
		err := m.DeleteClass(123)
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
