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

	//duplicate name
	err = l.UpdateRole(1, "role1")
	switch err {
	case nil:
		t.Fatal("expecting error", err)
	case model.ErrDuplicateKey:
	default:
		t.Fatal("unexpected error", err)
	}

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

func TestListRolesByIDs(t *testing.T) {
	var nf, id int = 999, 1
	l := New(nil)

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"author not found", []*int{&nf}, 0},
		{"no ids", []*int{nil}, 1},
		{"only author", []*int{&id}, 1},
	}

	for i := range tests {
		bs, err := l.ListRolesByIDs(tests[i].params[0])
		if err != nil {
			t.Fatalf("unexpected error for [%s], [%v]", tests[i].desc, err)
		}

		if tests[i].exp == 0 {
			if len(bs) != 0 {
				t.Errorf("unexpected length for [%s], exp 0 got %d", tests[i].desc, len(bs))
			}
		} else {
			if len(bs) != 1 {
				t.Fatalf("unexpected length for [%s], exp 1 got %d", tests[i].desc, len(bs))
			}
			if !reflect.DeepEqual(bs[0], l.roles[1]) {
				t.Fatalf("unexpected role value for [%s]", tests[i].desc)
			}
		}
	}
}
