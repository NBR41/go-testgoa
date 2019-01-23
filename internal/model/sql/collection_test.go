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
	qry := escapeQuery(qryGetCollectionByID)
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
	qry := escapeQuery(qryGetCollectionByName)
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

	qry := escapeQuery(qryInsertCollection)
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs("foo", 456).WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.Collection
		err  error
	}{
		{"duplicate error", nil, model.ErrDuplicateKey},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.Collection{ID: 123, Name: "foo", EditorID: 456}, nil},
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

func TestListCollectionsByIDs(t *testing.T) {
	var editorID, printID, seriesID int = 1, 2, 3
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := escapeQuery(`SELECT DISTINCT collection.id, collection.name, editor.id, editor.name FROM collection JOIN editor ON (collection.editor_id = editor.id) WHERE 1`)
	mock.ExpectQuery(qry).WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow("foo", "bar", "baz", "qux"))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar"))
	qry = escapeQuery(`SELECT DISTINCT collection.id, collection.name, editor.id, editor.name FROM collection JOIN editor ON (collection.editor_id = editor.id) JOIN edition ON (edition.collection_id = collection.id) JOIN book ON (edition.book_id = book.id) WHERE 1 AND editor.id = ? AND edition.print_id = ? AND book.series_id = ?`)
	mock.ExpectQuery(qry).WithArgs(editorID, printID, seriesID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "eid", "ename"}).AddRow(1, "foo", 2, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc   string
		params []*int
		exp    []model.Collection
		err    error
	}{
		{"query error", []*int{nil, nil, nil}, nil, errors.New("query error")},
		{"scan conversion error", []*int{nil, nil, nil}, nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", []*int{nil, nil, nil}, nil, errors.New("scan error")},
		{"valid", []*int{nil, nil, nil}, []model.Collection{model.Collection{ID: 1, Name: "foo", EditorID: 2, Editor: &model.Editor{ID: 2, Name: "bar"}}}, nil},
		{"valid with filter", []*int{&editorID, &printID, &seriesID}, []model.Collection{model.Collection{ID: 1, Name: "foo", EditorID: 2, Editor: &model.Editor{ID: 2, Name: "bar"}}}, nil},
	}

	for i := range tests {
		v, err := m.ListCollectionsByIDs(tests[i].params[0], tests[i].params[1], tests[i].params[2])
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
	mock.ExpectExec(`UPDATE collection SET name = \?, update_ts = NOW\(\) WHERE id = \?`).WithArgs("foo", 123).WillReturnError(errors.New("update error"))
	mock.ExpectExec(`UPDATE collection SET name = \?, editor_id = \?, update_ts = NOW\(\) WHERE id = \?`).WithArgs("foo", 456, 123).WillReturnError(errors.New("update error"))
	mock.ExpectExec(`UPDATE collection SET name = \?, editor_id = \?, update_ts = NOW\(\) WHERE id = \?`).WithArgs("foo", 456, 123).WillReturnError(errors.New(duplicateErr))
	mock.ExpectExec(`UPDATE collection SET name = \?, editor_id = \?, update_ts = NOW\(\) WHERE id = \?`).WithArgs("foo", 456, 123).WillReturnError(errors.New(fkErr))
	mock.ExpectExec(`UPDATE collection SET name = \?, editor_id = \?, update_ts = NOW\(\) WHERE id = \?`).WithArgs("foo", 456, 123).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(`UPDATE collection SET name = \?, editor_id = \?, update_ts = NOW\(\) WHERE id = \?`).WithArgs("foo", 456, 123).WillReturnResult(sqlmock.NewResult(0, 1))

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
		{"update name error", &name, nil, errors.New("update error")},
		{"name, editor error", &name, &editorID, errors.New("update error")},
		{"name, editor duplicate", &name, &editorID, model.ErrDuplicateKey},
		{"name, editor not found", &name, &editorID, model.ErrInvalidID},
		{"collection not found ", &name, &editorID, model.ErrNotFound},
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
	qry := escapeQuery(qryDeleteCollection)
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
