package sql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/kylelemons/godebug/pretty"
)

func TestGetUserList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, nickname, email, verified, admin FROM user`
	mock.
		ExpectQuery(qry).
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow("foo", "bar", "baz", "qux", "quux"))
	mock.
		ExpectQuery(qry).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1).RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  []model.User
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", []model.User{{ID: 1, Nickname: "nick", Email: "foo@bar.com", IsValidated: true, IsAdmin: true}}, nil},
	}

	for i := range tests {
		v, err := m.ListUsers()
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

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, nickname, email, activated, admin FROM user WHERE id = ?`
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}))
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow("foo", "bar", "baz", "qux", "quux"))
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1).RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.User
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.User{ID: 1, Nickname: "nick", Email: "foo@bar.com", IsValidated: true, IsAdmin: true}, nil},
	}

	for i := range tests {
		v, err := m.GetUserByID(123)
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

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, nickname, email, activated, admin FROM user WHERE email = ?`
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com").
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}))
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow("foo", "bar", "baz", "qux", "quux"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1).RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.User
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.User{ID: 1, Nickname: "nick", Email: "foo@bar.com", IsValidated: true, IsAdmin: true}, nil},
	}

	for i := range tests {
		v, err := m.GetUserByEmail("foo@bar.com")
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

func TestGetUserByNickname(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, nickname, email, activated, admin FROM user WHERE nickname = ?`
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow("foo", "bar", "baz", "qux", "quux"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1).RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WithArgs("foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.User
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.User{ID: 1, Nickname: "nick", Email: "foo@bar.com", IsValidated: true, IsAdmin: true}, nil},
	}

	for i := range tests {
		v, err := m.GetUserByNickname("foo")
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

func TestGetUserByEmailOrNickname(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `SELECT id, nickname, email, activated, admin FROM user WHERE email = \? OR nickname = \?`
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com", "foo").
		WillReturnError(errors.New("query error"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}))
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow("foo", "bar", "baz", "qux", "quux"))
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1).RowError(0, errors.New("scan error")))
	mock.
		ExpectQuery(qry).
		WithArgs("foo@bar.com", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "nick", "foo@bar.com", 1, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), nil)

	tests := []struct {
		desc string
		exp  *model.User
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"valid", &model.User{ID: 1, Nickname: "nick", Email: "foo@bar.com", IsValidated: true, IsAdmin: true}, nil},
	}

	for i := range tests {
		v, err := m.GetUserByEmailOrNickname("foo@bar.com", "foo")
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

func TestGetAuthenticatedUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockpassworder(ctrl)
	gomock.InOrder(
		mock.EXPECT().ComparePassword("bar", gomock.Any(), gomock.Any()).Return(false, errors.New("compare error")),
		mock.EXPECT().ComparePassword("bar", gomock.Any(), gomock.Any()).Return(false, nil),
		mock.EXPECT().ComparePassword("bar", gomock.Any(), gomock.Any()).Return(true, nil),
	)

	db, dbmock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `
SELECT id, nickname, email, activated, admin, salt, password
FROM user
WHERE email = \? OR nickname =\?`
	dbmock.
		ExpectQuery(qry).
		WithArgs("foo", "foo").
		WillReturnError(errors.New("query error"))
	dbmock.
		ExpectQuery(qry).
		WithArgs("foo", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin", "salt", "password"}))
	dbmock.
		ExpectQuery(qry).
		WithArgs("foo", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin", "salt", "password"}).AddRow("foo", "bar", "baz", "qux", "quux", "corge", "grault"))
	dbmock.
		ExpectQuery(qry).
		WithArgs("foo", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin", "salt", "password"}).AddRow(1, "nick", "foo@bar.com", 1, 1, []byte("corge"), []byte("grault")).RowError(0, errors.New("scan error")))
	dbmock.
		ExpectQuery(qry).
		WithArgs("foo", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin", "salt", "password"}).AddRow(1, "nick", "foo@bar.com", 1, 1, "corge", "grault"))
	dbmock.
		ExpectQuery(qry).
		WithArgs("foo", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin", "salt", "password"}).AddRow(1, "nick", "foo@bar.com", 1, 1, "corge", "grault"))
	dbmock.
		ExpectQuery(qry).
		WithArgs("foo", "foo").
		WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "activated", "admin", "salt", "password"}).AddRow(1, "nick", "foo@bar.com", 1, 1, "corge", "grault"))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), mock)

	tests := []struct {
		desc string
		exp  *model.User
		err  error
	}{
		{"query error", nil, errors.New("query error")},
		{"no rows", nil, model.ErrNotFound},
		{"scan conversion error", nil, errors.New(`sql: Scan error on column index 0, name "id": converting driver.Value type string ("foo") to a int64: invalid syntax`)},
		{"scan error", nil, errors.New("scan error")},
		{"compare error", nil, errors.New("compare error")},
		{"wrong password", nil, model.ErrInvalidCredentials},
		{"valid", &model.User{ID: 1, Nickname: "nick", Email: "foo@bar.com", IsValidated: true, IsAdmin: true}, nil},
	}

	for i := range tests {
		v, err := m.GetAuthenticatedUser("foo", "bar")
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

func TestInsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockpassworder(ctrl)
	gomock.InOrder(
		mock.EXPECT().CryptPassword("baz").Return(nil, nil, errors.New("crypt error")),
		mock.EXPECT().CryptPassword("baz").Return([]byte("qux"), []byte("quux"), nil),
		mock.EXPECT().CryptPassword("baz").Return([]byte("qux"), []byte("quux"), nil),
		mock.EXPECT().CryptPassword("baz").Return([]byte("qux"), []byte("quux"), nil),
	)

	db, dbmock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	userqry := `SELECT id, nickname, email, activated, admin FROM user WHERE email = \? OR nickname = \?`
	qry := `
INSERT INTO user \(id, nickname, email, salt, password, activated, admin, create_ts, update_ts\)
VALUES \(null, \?, \?, \?, \?, 0, 0, NOW\(\), NOW\(\)\)
ON DUPLICATE KEY UPDATE update_ts = VALUES\(update_ts\)`
	dbmock.ExpectQuery(userqry).WithArgs("foo", "bar").WillReturnError(errors.New("get user error"))

	dbmock.ExpectQuery(userqry).WithArgs("foo", "bar").WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}).AddRow(1, "foo", "bar", 1, 1))

	dbmock.ExpectQuery(userqry).WithArgs("foo", "bar").WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}))

	dbmock.ExpectQuery(userqry).WithArgs("foo", "bar").WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}))
	dbmock.ExpectExec(qry).WithArgs("bar", "foo", []byte("qux"), []byte("quux")).WillReturnError(errors.New("query error"))

	dbmock.ExpectQuery(userqry).WithArgs("foo", "bar").WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}))
	dbmock.ExpectExec(qry).WithArgs("bar", "foo", []byte("qux"), []byte("quux")).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))

	dbmock.ExpectQuery(userqry).WithArgs("foo", "bar").WillReturnRows(sqlmock.NewRows([]string{"id", "nickname", "email", "verified", "admin"}))
	dbmock.ExpectExec(qry).WithArgs("bar", "foo", []byte("qux"), []byte("quux")).WillReturnResult(sqlmock.NewResult(123, 1))

	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), mock)

	tests := []struct {
		desc string
		exp  *model.User
		err  error
	}{
		{"get user error", nil, errors.New("get user error")},
		{"duplicate", nil, model.ErrDuplicateKey},
		{"crypt error", nil, errors.New("crypt error")},
		{"query error", nil, errors.New("query error")},
		{"result error", nil, errors.New("result error")},
		{"valid", &model.User{ID: 123, Email: "foo", Nickname: "bar"}, nil},
	}

	for i := range tests {
		v, err := m.InsertUser("foo", "bar", "baz")
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

func TestUpdateUserNickname(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `UPDATE user set nickname = \?, update_ts = NOW\(\) where id = \?`
	mock.ExpectExec(qry).WithArgs("foo", 123).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs("foo", 123).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
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
		{"result error", errors.New("result error")},
		{"no rows", model.ErrNotFound},
		{"valid", nil},
	}
	for i := range tests {
		err := m.UpdateUserNickname(123, "foo")
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

func TestUpdateUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockpassworder(ctrl)
	gomock.InOrder(
		mock.EXPECT().CryptPassword("foo").Return(nil, nil, errors.New("password error")),
		mock.EXPECT().CryptPassword("foo").Return([]byte("gulp"), []byte("qux"), nil),
		mock.EXPECT().CryptPassword("foo").Return([]byte("gulp"), []byte("qux"), nil),
		mock.EXPECT().CryptPassword("foo").Return([]byte("gulp"), []byte("qux"), nil),
		mock.EXPECT().CryptPassword("foo").Return([]byte("gulp"), []byte("qux"), nil),
	)

	db, dbmock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `UPDATE user set salt = \?, password = \?, update_ts = NOW\(\) where id = \?`
	dbmock.ExpectExec(qry).WithArgs([]byte("gulp"), []byte("qux"), 123).WillReturnError(errors.New("query error"))
	dbmock.ExpectExec(qry).WithArgs([]byte("gulp"), []byte("qux"), 123).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	dbmock.ExpectExec(qry).WithArgs([]byte("gulp"), []byte("qux"), 123).WillReturnResult(sqlmock.NewResult(0, 0))
	dbmock.ExpectExec(qry).WithArgs([]byte("gulp"), []byte("qux"), 123).WillReturnResult(sqlmock.NewResult(0, 1))
	m, _ := New(ConnGetter(func() (*sql.DB, error) {
		return db, nil
	}), mock)
	tests := []struct {
		desc string
		err  error
	}{
		{"password error", errors.New("password error")},
		{"query error", errors.New("query error")},
		{"result error", errors.New("result error")},
		{"no rows", model.ErrNotFound},
		{"valid", nil},
	}
	for i := range tests {
		err := m.UpdateUserPassword(123, "foo")
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

func TestUpdateUserActivation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `UPDATE user set activated = \?, update_ts = NOW\(\) where id = \?`
	mock.ExpectExec(qry).WithArgs(true, 123).WillReturnError(errors.New("query error"))
	mock.ExpectExec(qry).WithArgs(true, 123).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	mock.ExpectExec(qry).WithArgs(true, 123).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(qry).WithArgs(true, 123).WillReturnResult(sqlmock.NewResult(0, 1))
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
		err := m.UpdateUserActivation(123, true)
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

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qry := `DELETE FROM user where id = \?`
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
		err := m.DeleteUser(123)
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
