package local

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/golang/mock/gomock"
)

func TestListUsers(t *testing.T) {
	l := New(nil)
	us, err := l.ListUsers()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range us {
		if int64(i)+1 != us[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestGetUserByID(t *testing.T) {
	l := New(nil)

	u, err := l.GetUserByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if u.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", u.ID)
	}

	u, err = l.GetUserByID(10)
	expectingError(t, err, model.ErrNotFound)
}

func TestGetUserByEmailOrNickname(t *testing.T) {
	l := New(nil)

	for _, v := range []struct {
		email    string
		nickname string
	}{
		{email: "admin@myinventory.com", nickname: "admin"},
		{email: "admin@myinventory.com", nickname: "foo"},
		{email: "foo", nickname: "admin"},
	} {
		u, err := l.GetUserByEmailOrNickname(v.email, v.nickname)
		if err != nil {
			t.Fatal("unexpected error", err)
		}
		if u.ID != 1 {
			t.Fatalf("expecting ID 1, got %v", u.ID)
		}
	}

	_, err := l.GetUserByEmailOrNickname("foo", "bar")
	expectingError(t, err, model.ErrNotFound)
}

func TestGetUserByEmail(t *testing.T) {
	l := New(nil)

	// found
	u, err := l.GetUserByEmail("admin@myinventory.com")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", u.ID)
	}

	// not found
	_, err = l.GetUserByEmail("foo")
	expectingError(t, err, model.ErrNotFound)
}

func TestGetUserByNickname(t *testing.T) {
	l := New(nil)

	// found
	u, err := l.GetUserByNickname("admin")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", u.ID)
	}

	// not found
	_, err = l.GetUserByNickname("foo")
	expectingError(t, err, model.ErrNotFound)
}

func TestGetAuthenticatedUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockpassworder(ctrl)
	gomock.InOrder(
		mock.EXPECT().ComparePassword("bar", gomock.Any(), gomock.Any()).Return(false, errors.New("compare error")),
		mock.EXPECT().ComparePassword("bar", gomock.Any(), gomock.Any()).Return(false, nil),
		mock.EXPECT().ComparePassword("bar", gomock.Any(), gomock.Any()).Return(false, nil),
		mock.EXPECT().ComparePassword("passwordadmin", gomock.Any(), gomock.Any()).Return(true, nil),
		mock.EXPECT().ComparePassword("passwordadmin", gomock.Any(), gomock.Any()).Return(true, nil),
	)
	l := New(mock)

	_, err := l.GetAuthenticatedUser("foo", "bar")
	expectingError(t, err, model.ErrNotFound)
	_, err = l.GetAuthenticatedUser("admin", "bar")
	expectingError(t, err, errors.New("compare error"))
	_, err = l.GetAuthenticatedUser("admin", "bar")
	expectingError(t, err, model.ErrInvalidCredentials)

	_, err = l.GetAuthenticatedUser(`admin@myinventory.com`, "bar")
	expectingError(t, err, model.ErrInvalidCredentials)

	u, err := l.GetAuthenticatedUser("admin", "passwordadmin")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", u.ID)
	}

	u, err = l.GetAuthenticatedUser(`admin@myinventory.com`, "passwordadmin")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", u.ID)
	}
}

func TestInsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockpassworder(ctrl)
	gomock.InOrder(
		mock.EXPECT().CryptPassword("foobar").Return([]byte("gulp"), []byte("qux"), nil),
		mock.EXPECT().CryptPassword("barfoo").Return([]byte("gulp"), []byte("qux"), nil),
		mock.EXPECT().CryptPassword("barfoo").Return([]byte("gulp"), []byte("qux"), nil),
		mock.EXPECT().CryptPassword("barfoo").Return([]byte("gulp"), []byte("qux"), nil),
	)
	l := New(mock)

	// insert
	u, err := l.InsertUser("foo", "bar", "foobar")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.ID != 5 {
		t.Fatalf("expecting ID 5, got %v", u.ID)
	}
	u2, err := l.GetUserByID(5)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(u, u2) {
		t.Fatal("unexpected user value")
	}

	// duplicate email
	u, err = l.InsertUser("foo", "foobar", "barfoo")
	expectingError(t, err, model.ErrDuplicateEmail)

	// duplicate nickname
	u, err = l.InsertUser("foobar", "bar", "barfoo")
	expectingError(t, err, model.ErrDuplicateNickname)

	// duplicate both
	u, err = l.InsertUser("foo", "bar", "barfoo")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateUserNickname(t *testing.T) {
	l := New(nil)

	// nickname exist
	err := l.UpdateUserNickname(2, "admin")
	expectingError(t, err, model.ErrDuplicateNickname)
	// user doesn't exist
	err = l.UpdateUserNickname(10, "admin2")
	expectingError(t, err, model.ErrNotFound)

	// update with existing nickname but same user
	err = l.UpdateUserNickname(1, "admin")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	// update user 1 nickname
	err = l.UpdateUserNickname(1, "admin2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	u, err := l.GetUserByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Nickname != "admin2" {
		t.Fatalf("expecting admin2, got %s", u.Nickname)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockpassworder(ctrl)
	gomock.InOrder(
		mock.EXPECT().CryptPassword("foo").Return([]byte("gulp"), []byte("qux"), nil),
		mock.EXPECT().CryptPassword("").Return([]byte("gulp"), []byte("qux"), nil),
	)
	l := New(mock)

	err := l.UpdateUserPassword(5, "foo")
	expectingError(t, err, model.ErrNotFound)

	err = l.UpdateUserPassword(1, "")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if bytes.Compare(l.users[1].Password, []byte("qux")) != 0 {
		t.Fatal("unexpected password valud")
	}
	if bytes.Compare(l.users[1].Salt, []byte("gulp")) != 0 {
		t.Fatal("unexpected salt valud")
	}
}

func TestUpdateUserActivation(t *testing.T) {
	l := New(nil)

	// user doesn't exists
	err := l.UpdateUserActivation(10, true)
	expectingError(t, err, model.ErrNotFound)

	// user activation
	err = l.UpdateUserActivation(1, false)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	u, err := l.GetUserByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.IsValidated {
		t.Fatal("expecting false got true")
	}

	err = l.UpdateUserActivation(1, true)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	u, err = l.GetUserByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !u.IsValidated {
		t.Fatal("expecting true got false")
	}
}

func TestDeleteUser(t *testing.T) {
	l := New(nil)

	// user doesn't exists
	err := l.DeleteUser(10)
	expectingError(t, err, model.ErrNotFound)

	// user delete
	err = l.DeleteUser(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetUserByID(1)
	expectingError(t, err, model.ErrNotFound)
}
