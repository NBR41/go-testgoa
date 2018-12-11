package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetRoleByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetRoleByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetRoleByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetRoleByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetRoleByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetRoleByName("role1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetRoleList(t *testing.T) {
	l := New(nil)

	bs, err := l.ListRoles()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertRole(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertRole("role2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetRoleByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate role name
	b, err = l.InsertRole("role2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateRole(t *testing.T) {
	l := New(nil)

	// role doesn't exist
	err := l.UpdateRole(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//update role
	err = l.UpdateRole(1, "role2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetRoleByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "role2" {
		t.Fatalf("expecting role2, got %s", u.Name)
	}
}

func TestDeleteRole(t *testing.T) {
	l := New(nil)

	// role doesn't exists
	err := l.DeleteRole(10)
	expectingError(t, err, model.ErrNotFound)

	// role delete
	err = l.DeleteRole(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetRoleByID(1)
	expectingError(t, err, model.ErrNotFound)
}
