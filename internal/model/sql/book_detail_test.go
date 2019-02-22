package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/kylelemons/godebug/pretty"
)

func TestGetBookDetail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	columns := []string{"book.id", "book.isbn", "book.name", "series.id", "series.name", "category.id", "category.name", "edition.id", "print.id", "print.name", "collection.id", "collection.name", "editor.id", "editor.name"}
	qry := escapeQuery(qryGetBookDetail)
	classQry := escapeQuery(`SELECT DISTINCT class.id, class.name from class JOIN classification ON (classification.class_id = class.id) WHERE 1 AND classification.series_id = ?`)
	authorQry := escapeQuery(qryListAuthorships + ` WHERE authorship.book_id = ?`)
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnError(errors.New("query error"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows(columns))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "foo", "bar", 2, "baz", 3, "qux", 12, 4, "quux", 5, "corge", 6, "grault").RowError(0, errors.New("scan error")))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "foo", "bar", 2, "baz", 3, "qux", 12, 4, "quux", 5, "corge", 6, "grault"))
	mock.ExpectQuery(classQry).WithArgs(2).WillReturnError(errors.New("class error"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "foo", "bar", 2, "baz", 3, "qux", 12, 4, "quux", 5, "corge", 6, "grault"))
	mock.ExpectQuery(classQry).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"class.id", "class.name"}).AddRow(7, "garply").AddRow(8, "waldo"))
	mock.ExpectQuery(authorQry).WithArgs(1).WillReturnError(errors.New("authorship error"))
	mock.ExpectQuery(qry).WithArgs("foo").WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "foo", "bar", 2, "baz", 3, "qux", 12, 4, "quux", 5, "corge", 6, "grault"))
	mock.ExpectQuery(classQry).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"class.id", "class.name"}).AddRow(7, "garply").AddRow(8, "waldo"))
	mock.ExpectQuery(authorQry).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"authorship.id", "author.id", "author.name", "role.id", "role.name", "book.id", "book.name"}).AddRow(9, 10, "fred", 11, "plugh", 1, "bar"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.BookDetail
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan error", nil, errors.New("scan error")},
		{"list class error", nil, errors.New("class error")},
		{"list authorship error", nil, errors.New("authorship error")},
		{
			"valid",
			&model.BookDetail{
				Book: &model.Book{
					ID:       1,
					ISBN:     "foo",
					Name:     "bar",
					URL:      "",
					SeriesID: 2,
					Series: &model.Series{
						ID:         2,
						Name:       "baz",
						CategoryID: 3,
						Category: &model.Category{
							ID:   3,
							Name: "qux",
						},
					},
				},
				Classes: []*model.Class{
					&model.Class{ID: 7, Name: "garply"},
					&model.Class{ID: 8, Name: "waldo"},
				},
				Edition: &model.Edition{
					ID:           12,
					BookID:       1,
					CollectionID: 5,
					Collection: &model.Collection{
						ID:       5,
						Name:     "corge",
						EditorID: 6,
						Editor:   &model.Editor{ID: 6, Name: "grault"},
					},
					PrintID: 4,
					Print:   &model.Print{ID: 4, Name: "quux"},
				},
				Authors: []*model.Authorship{
					&model.Authorship{
						ID:       9,
						AuthorID: 10,
						Author:   &model.Author{ID: 10, Name: "fred"},
						BookID:   1,
						Book:     &model.Book{ID: 1, Name: "bar"},
						RoleID:   11,
						Role:     &model.Role{ID: 11, Name: "plugh"},
					},
				},
			},
			nil,
		},
	}

	for i := range tests {
		v, err := m.GetBookDetail("foo")
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
