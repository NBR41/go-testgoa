package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetClassByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetClassByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetClassByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetClassByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetClassByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetClassByName("class1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestInsertClass(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertClass("genre2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetClassByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate genre name
	b, err = l.InsertClass("genre2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateClass(t *testing.T) {
	l := New(nil)

	// genre doesn't exist
	err := l.UpdateClass(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//duplicate name
	err = l.UpdateClass(1, "class1")
	switch err {
	case nil:
		t.Fatal("expecting error", err)
	case model.ErrDuplicateKey:
	default:
		t.Fatal("unexpected error", err)
	}

	//update genre
	err = l.UpdateClass(1, "genre2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetClassByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "genre2" {
		t.Fatalf("expecting genre2, got %s", u.Name)
	}
}

func TestDeleteClass(t *testing.T) {
	l := New(nil)

	// genre doesn't exists
	err := l.DeleteClass(10)
	expectingError(t, err, model.ErrNotFound)

	// genre delete
	err = l.DeleteClass(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetClassByID(1)
	expectingError(t, err, model.ErrNotFound)
}

func TestListClassesByIDs(t *testing.T) {
	var nf, id int = 999, 1
	l := New(nil)

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"not founds", []*int{&nf, &nf, &nf}, 0},
		{"author not found", []*int{&nf, nil, nil}, 0},
		{"category not found", []*int{nil, &nf, nil}, 0},
		{"series not found", []*int{nil, nil, &nf}, 0},
		{"author id category not found series not found", []*int{&id, &nf, &nf}, 0},
		{"author not found category id series not found", []*int{&nf, &id, &nf}, 0},
		{"author not found category not found series id", []*int{&nf, &nf, &id}, 0},
		{"no ids", []*int{nil, nil, nil}, 1},
		{"only author", []*int{&id, nil, nil}, 1},
		{"only class", []*int{nil, &id, nil}, 1},
		{"only class", []*int{nil, nil, &id}, 1},
		{"author and class", []*int{&id, &id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListClassesByIDs(tests[i].params[0], tests[i].params[1], tests[i].params[2])
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
			if !reflect.DeepEqual(bs[0], l.classes[1]) {
				t.Fatalf("unexpected class value for [%s]", tests[i].desc)
			}
		}
	}
}
