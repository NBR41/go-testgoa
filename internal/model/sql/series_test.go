package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestGetSeriesByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetSeriesByID)
	mock.ExpectQuery(qry).WithArgs(123).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Series
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Series{ID: 1, Name: "foo", CategoryID: 2, Category: &model.Category{ID: 2, Name: "bar"}}, nil},
	}

	for i := range tests {
		v, err := m.GetSeriesByID(123)
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

func TestGetSeriesByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryGetSeriesByName)
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Series
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Series{ID: 1, Name: "foo", CategoryID: 2, Category: &model.Category{ID: 2, Name: "bar"}}, nil},
	}

	for i := range tests {
		v, err := m.GetSeriesByName("foo")
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
func TestListSeries(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListSeries)
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Series
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Series{&model.Series{ID: 1, Name: "foo", CategoryID: 2, Category: &model.Category{ID: 2, Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListSeries()
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

func TestListSeriesByIDs1(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListSeries)
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Series
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Series{&model.Series{ID: 1, Name: "foo", CategoryID: 2, Category: &model.Category{ID: 2, Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListSeriesByIDs(nil, nil, nil, nil)
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

func TestListSeriesByIDs2(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	var authorID, roleID, categoryID, classID int = 1, 2, 3, 4
	qry := escapeQuery(`SELECT DISTINCT s.id, s.name, c.id, c.name FROM series s JOIN category c ON (c.id = s.series_id) JOIN book b ON (b.series_id = s.id) JOIN authorship a ON (a.book_id = b.id) JOIN classification cl ON (cl.series_id = s.id) WHERE a.author_id = ? AND a.role_id = ? AND cl.class_id = ? AND c.id = ?`)
	mock.ExpectQuery(qry).WithArgs(authorID, roleID, classID, categoryID).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(authorID, roleID, classID, categoryID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "sid", "sname"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WithArgs(authorID, roleID, classID, categoryID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "sid", "sname"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(authorID, roleID, classID, categoryID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "sid", "sname"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.Series
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.Series{model.Series{ID: 1, Name: "foo", CategoryID: 2, Category: &model.Category{ID: 2, Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListSeriesByIDs(&authorID, &roleID, &categoryID, &classID)
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

func TestListSeriesByCollectionID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListSeriesByCollectionID)
	mock.ExpectQuery(qry).WithArgs(2).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Series
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Series{&model.Series{ID: 1, Name: "foo", CategoryID: 2, Category: &model.Category{ID: 2, Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListSeriesByCollectionID(2)
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

func TestListSeriesByCollectionIDPrintID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryListSeriesByCollectionID)
	mock.ExpectQuery(qry).WithArgs(2, 3).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(2, 3).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WithArgs(2, 3).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(2, 3).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cid", "cname"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Series
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Series{&model.Series{ID: 1, Name: "foo", CategoryID: 2, Category: &model.Category{ID: 2, Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListSeriesByCollectionIDPrintID(2, 3)
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

func TestInsertSeries(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := escapeQuery(qryInsertSeries)
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Series
		err  error
	}{
		{"duplicate error", nil, model.ErrDuplicateKey},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Series{ID: 123, Name: "foo", CategoryID: 456}, nil},
	}

	for i := range tests {
		v, err := m.InsertSeries("foo", 456)
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

func TestUpdateSeries0(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)
	err = m.UpdateSeries(123, nil, nil)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
}

func TestUpdateSeries1(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	n1 := "foo"
	qry := `UPDATE series SET name = \?, update_ts = NOW\(\) WHERE id = \?`
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
		err := m.UpdateSeries(123, &n1, nil)
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

func TestUpdateSeries2(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	category := 456
	qry := `UPDATE series SET category_id = \?, update_ts = NOW\(\) WHERE id = \?`
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
		err := m.UpdateSeries(123, nil, &category)
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

func TestUpdateSeries4(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	name := "foo"
	category := 456
	qry := escapeQuery(`UPDATE series SET name = ?, category_id = ?, update_ts = NOW() WHERE id = ?`)
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
		err := m.UpdateSeries(123, &name, &category)
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

func TestDeleteSeries(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := escapeQuery(qryDeleteSeries)
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
		err := m.DeleteSeries(123)
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
