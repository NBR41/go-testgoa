package appmodel

import (
	"reflect"
	"testing"
)

func TestGetUserList(t *testing.T) {

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
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}

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
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}
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
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}
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
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}
}

func TestGetAuthenticatedUser(t *testing.T) {
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
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrDuplicateKey {
			t.Fatal("unexpected error", err)
		}
	}

	// duplicate nickname
	u, err = l.InsertUser("foobar", "bar", "barfoo")
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrDuplicateKey {
			t.Fatal("unexpected error", err)
		}
	}
}

func TestUpdateUserNickname(t *testing.T) {
	l := NewLocal()

	// nickname exist
	err := l.UpdateUserNickname(2, "admin")
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrDuplicateKey {
			t.Fatal("unexpected error", err)
		}
	}

	// user doesn't exist
	err = l.UpdateUserNickname(10, "admin2")
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}

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
}

func TestUpdateUserActivation(t *testing.T) {
	l := NewLocal()

	// user doesn't exists
	err := l.UpdateUserActivation(10, true)
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}

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
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}

	// user delete
	err = l.DeleteUser(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetUserByID(1)
	if err != ErrNotFound {
		t.Fatal("expecting ErrNotFound, got %s", err)
	}
}

func TestInsertBook(t *testing.T) {
}

func TestGetBookByID(t *testing.T) {
}

func TestGetBookByName(t *testing.T) {
}

func TestGetBookList(t *testing.T) {
}

func TestUpdateBook(t *testing.T) {
	l := NewLocal()

	// bookname exist
	err := l.UpdateBook(2, "test1")
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrDuplicateKey {
			t.Fatal("unexpected error", err)
		}
	}

	// user doesn't exist
	err = l.UpdateBook(10, "test10")
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}

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
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err != ErrNotFound {
			t.Fatal("unexpected error", err)
		}
	}

	// book delete
	err = l.DeleteBook(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetBookByID(1)
	if err != ErrNotFound {
		t.Fatal("expecting ErrNotFound, got %s", err)
	}
}

func TestGetOwnershipList(t *testing.T) {
}

func TestGetOwnership(t *testing.T) {
}

func TestInsertOwnership(t *testing.T) {
}

func TestUpdateOwnership(t *testing.T) {

}
func TestDeleteOwnership(t *testing.T) {
}
