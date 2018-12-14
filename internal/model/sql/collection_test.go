package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestGetCollectionByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `
SELECT c.id, c.name, e.id, e.name
FROM collection c
JOIN editors e ON \(e.id = c.editor_id\)
WHERE c.id = \?`
	mock.ExpectQuery(qry).WithArgs(123).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Collection
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Collection{ID: 1, Name: "foo", EditorID: 2, Editor: &model.Editor{ID: 2, Name: "bar"}}, nil},
	}

	for i := range tests {
		v, err := m.GetCollectionByID(123)
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

func TestGetCollectionByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `
SELECT c.id, c.name, e.id, e.name
FROM collection c
JOIN editors e ON \(e.id = c.editor_id\)
WHERE c.name = \?`
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Collection
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.Collection{ID: 1, Name: "foo", EditorID: 2, Editor: &model.Editor{ID: 2, Name: "bar"}}, nil},
	}

	for i := range tests {
		v, err := m.GetCollectionByName("foo")
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

func TestInsertCollection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := `
INSERT INTO collection \(id, name, editor_id, create_ts, update_ts\)
VALUES \(null, \?, \?, NOW\(\), NOW\(\)\)
ON DUPLICATE KEY UPDATE update_ts = VALUES\(update_ts\)`

	idqry := `SELECT id, name FROM editor where id = \?`
	mock.ExpectQuery(idqry).WithArgs(456).WillReturnError(errors.New("duplicate error"))
	mock.ExpectQuery(idqry).WithArgs(456).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(456, "bar"))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(idqry).WithArgs(456).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(456, "bar"))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectQuery(idqry).WithArgs(456).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(456, "bar"))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Collection
		err  error
	}{
		{"duplicate error", nil, errors.New("duplicate error")},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Collection{ID: 123, Name: "foo", EditorID: 456, Editor: &model.Editor{ID: 456, Name: "bar"}}, nil},
	}

	for i := range tests {
		v, err := m.InsertCollection("foo", 456)
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

func TestUpdateCollection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	idqry := `SELECT id, name FROM editor where id = \?`
	mock.ExpectExec(`UPDATE collection SET name = \?, update_ts = NOW\(\) where id = \?`).WithArgs("foo", 123).WillReturnError(errors.New("update error"))
	mock.ExpectQuery(idqry).WithArgs(456).WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery(idqry).WithArgs(456).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(456, "bar"))
	mock.ExpectExec(`UPDATE collection SET name = \?, editor_id = \?, update_ts = NOW\(\) where id = \?`).WithArgs("foo", 456, 123).WillReturnError(errors.New("update error"))
	mock.ExpectQuery(idqry).WithArgs(456).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(456, "bar"))
	mock.ExpectExec(`UPDATE collection SET name = \?, editor_id = \?, update_ts = NOW\(\) where id = \?`).WithArgs("foo", 456, 123).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery(idqry).WithArgs(456).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(456, "bar"))
	mock.ExpectExec(`UPDATE collection SET name = \?, editor_id = \?, update_ts = NOW\(\) where id = \?`).WithArgs("foo", 456, 123).WillReturnResult(sqlmock.NewResult(0, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	name := "foo"
	editorID := 456
	tests := []struct {
		desc     string
		name     *string
		editorID *int
		err      error
	}{
		{"no values", nil, nil, nil},
		{"name error", &name, nil, errors.New("update error")},
		{"name, editor error not found", &name, &editorID, model.ErrNotFound},
		{"name, editor error", &name, &editorID, errors.New("update error")},
		{"name, editor not found", &name, &editorID, model.ErrNotFound},
		{"name, editor valid", &name, &editorID, nil},
	}
	for i := range tests {
		err := m.UpdateCollection(123, tests[i].name, tests[i].editorID)
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

func TestDeleteCollection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `DELETE FROM collection where id = \?`
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
		err := m.DeleteCollection(123)
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

func TestListCollections(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT c.id, c.name, e.id, e.name FROM collection c JOIN editors e ON \(e.id = c.editor_id\)`
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Collection
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Collection{&model.Collection{ID: 1, Name: "foo", EditorID: 2, Editor: &model.Editor{ID: 2, Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListCollections()
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

func TestListCollectionsByEditorID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT c.id, c.name, e.id, e.name FROM collection c JOIN editors e ON \(e.id = c.editor_id\) where e.id = \?`
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []*model.Collection
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []*model.Collection{&model.Collection{ID: 1, Name: "foo", EditorID: 2, Editor: &model.Editor{ID: 2, Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListCollectionsByEditorID(2)
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
