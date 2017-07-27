package appmodel

import (
	"reflect"
	"testing"
)

func expectingError(t *testing.T, err, exp error) {
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != exp {
			t.Fatal("unexpected error", err)
		}
	}
}

func TestGetUserList(t *testing.T) {
	l := NewLocal()
	us, err := l.GetUserList()
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
	l := NewLocal()

	u, err := l.GetUserByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if u.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", u.ID)
	}

	u, err = l.GetUserByID(10)
	expectingError(t, err, ErrNotFound)
}

func TestGetUserByEmailOrNickname(t *testing.T) {
	l := NewLocal()

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
	expectingError(t, err, ErrNotFound)
}

func TestGetUserByEmail(t *testing.T) {
	l := NewLocal()

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
	expectingError(t, err, ErrNotFound)
}

func TestGetUserByNickname(t *testing.T) {
	l := NewLocal()

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
	expectingError(t, err, ErrNotFound)
}

func TestGetAuthenticatedUser(t *testing.T) {
	l := NewLocal()

	u, err := l.GetAuthenticatedUser("foo", "bar")
	expectingError(t, err, ErrNotFound)

	u, err = l.GetAuthenticatedUser("admin", "bar")
	expectingError(t, err, ErrInvalidCredentials)

	u, err = l.GetAuthenticatedUser(`admin@myinventory.com`, "bar")
	expectingError(t, err, ErrInvalidCredentials)

	u, err = l.GetAuthenticatedUser("admin", "passwordadmin")
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
	l := NewLocal()

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
	expectingError(t, err, ErrDuplicateKey)

	// duplicate nickname
	u, err = l.InsertUser("foobar", "bar", "barfoo")
	expectingError(t, err, ErrDuplicateKey)
}

func TestUpdateUserNickname(t *testing.T) {
	l := NewLocal()

	// nickname exist
	err := l.UpdateUserNickname(2, "admin")
	expectingError(t, err, ErrDuplicateKey)
	// user doesn't exist
	err = l.UpdateUserNickname(10, "admin2")
	expectingError(t, err, ErrNotFound)

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
	l := NewLocal()

	err := l.UpdateUserPassword(5, "foo")
	expectingError(t, err, ErrNotFound)

	err = l.UpdateUserPassword(1, "")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	_, err = l.GetAuthenticatedUser("admin", "")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
}

func TestUpdateUserActivation(t *testing.T) {
	l := NewLocal()

	// user doesn't exists
	err := l.UpdateUserActivation(10, true)
	expectingError(t, err, ErrNotFound)

	// user activation
	err = l.UpdateUserActivation(1, false)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	u, err := l.GetUserByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.IsVerified {
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
	if !u.IsVerified {
		t.Fatal("expecting true got false")
	}
}

func TestDeleteUser(t *testing.T) {
	l := NewLocal()

	// user doesn't exists
	err := l.DeleteUser(10)
	expectingError(t, err, ErrNotFound)

	// user delete
	err = l.DeleteUser(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetUserByID(1)
	expectingError(t, err, ErrNotFound)
}

func TestInsertBook(t *testing.T) {
	l := NewLocal()

	// insert
	b, err := l.InsertBook("foo")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 5 {
		t.Fatalf("expecting ID 5, got %v", b.ID)
	}
	b2, err := l.GetBookByID(5)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate book name
	b, err = l.InsertBook("foo")
	expectingError(t, err, ErrDuplicateKey)
}

func TestGetBookByID(t *testing.T) {
	l := NewLocal()

	b, err := l.GetBookByID(5)
	expectingError(t, err, ErrNotFound)

	b, err = l.GetBookByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetBookByName(t *testing.T) {
	l := NewLocal()

	b, err := l.GetBookByName("test5")
	expectingError(t, err, ErrNotFound)

	b, err = l.GetBookByName("test1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetBookList(t *testing.T) {
	l := NewLocal()
	bs, err := l.GetBookList()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestUpdateBook(t *testing.T) {
	l := NewLocal()

	// bookname exist
	err := l.UpdateBook(2, "test1")
	expectingError(t, err, ErrDuplicateKey)
	// user doesn't exist
	err = l.UpdateBook(10, "test10")
	expectingError(t, err, ErrNotFound)

	// update with existing nickname but same user
	err = l.UpdateBook(1, "test1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	// update user 1 nickname
	err = l.UpdateBook(1, "test10")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	u, err := l.GetBookByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "test10" {
		t.Fatalf("expecting test10, got %s", u.Name)
	}

}

func TestDeleteBook(t *testing.T) {
	l := NewLocal()

	// book doesn't exists
	err := l.DeleteBook(10)
	expectingError(t, err, ErrNotFound)

	// book delete
	err = l.DeleteBook(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetBookByID(1)
	expectingError(t, err, ErrNotFound)
}

func TestGetOwnershipList(t *testing.T) {
	l := NewLocal()

	_, err := l.GetOwnershipList(5)
	expectingError(t, err, ErrNotFound)

	ows, err := l.GetOwnershipList(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	for i := range ows {
		if ows[i].UserID != 1 {
			t.Fatalf("unexpected user ID, expecting 1, got %d", ows[i].UserID)
		}

		if ows[i].BookID != l.ownerships[1][i].ID {
			t.Fatalf("unexpected user ID, expecting %d, got %d", l.ownerships[1][i].ID, ows[i].BookID)
		}
	}
}

func TestGetOwnership(t *testing.T) {
	l := NewLocal()

	// not found
	o, err := l.GetOwnership(5, 5)
	expectingError(t, err, ErrNotFound)
	// user not found
	o, err = l.GetOwnership(5, 1)
	expectingError(t, err, ErrNotFound)
	// book not found
	o, err = l.GetOwnership(1, 5)
	expectingError(t, err, ErrNotFound)

	//  found
	o, err = l.GetOwnership(1, 1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if o.UserID != 1 || o.BookID != 1 || o.Book.ID != 1 {
		t.Fatal("unexpected value", o)
	}
}

func TestInsertOwnership(t *testing.T) {
	l := NewLocal()

	// not found
	o, err := l.InsertOwnership(5, 5)
	expectingError(t, err, ErrNotFound)
	// user not found
	o, err = l.InsertOwnership(5, 1)
	expectingError(t, err, ErrNotFound)
	// book not found
	o, err = l.InsertOwnership(1, 5)
	expectingError(t, err, ErrNotFound)
	//  duplicate key
	o, err = l.InsertOwnership(1, 1)
	expectingError(t, err, ErrDuplicateKey)

	//  found
	o, err = l.InsertOwnership(1, 2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if o.UserID != 1 || o.BookID != 2 || o.Book.ID != 2 {
		t.Fatal("unexpected value", o)
	}
}

func TestUpdateOwnership(t *testing.T) {

}

func TestDeleteOwnership(t *testing.T) {
	l := NewLocal()

	// not found
	err := l.DeleteOwnership(5, 5)
	expectingError(t, err, ErrNotFound)
	// user not found
	err = l.DeleteOwnership(5, 1)
	expectingError(t, err, ErrNotFound)
	// book not found
	err = l.DeleteOwnership(1, 5)
	expectingError(t, err, ErrNotFound)

	//  found
	err = l.DeleteOwnership(1, 1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	_, err = l.GetOwnership(1, 1)
	expectingError(t, err, ErrNotFound)
}
