package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetCollectionByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetCollectionByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetCollectionByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetCollectionByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetCollectionByName("test5", 5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetCollectionByName("collection1",1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestInsertCollection(t *testing.T) {
	l := New(nil)

	_, err := l.InsertCollection("collection1", 1)
	expectingError(t, err, model.ErrDuplicateKey)
	// editor not found
	_, err = l.InsertCollection("collection2", 999)
	expectingError(t, err, model.ErrInvalidID)

	// insert
	b, err := l.InsertCollection("collection2", 1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetCollectionByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}
}

func TestUpdateCollection(t *testing.T) {
	l := New(nil)
	n0 := "test10"
	n1 := "collection1"
	n2 := "collection2"
	n3 := "collection3"
	e1 := 5
	e, err := l.InsertEditor("editor2")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	e2 := int(e.ID)

	// collection doesn't exist
	err = l.UpdateCollection(10, &n0, nil)
	expectingError(t, err, model.ErrNotFound)
	//same name
	err = l.UpdateCollection(1, &n1, &e2)
	expectingError(t, err, model.ErrDuplicateKey)
	t.Log(n3)

	// editor doesn't exist
	err = l.UpdateCollection(1, &n2, &e1)
	expectingError(t, err, model.ErrInvalidID)

	//update collection
	err = l.UpdateCollection(1, &n3, &e2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetCollectionByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != n3 {
		t.Fatalf("expecting %s, got %s", n3, u.Name)
	}
	if u.EditorID != e.ID {
		t.Fatalf("expecting editor id %d, got %d", e.ID, u.EditorID)
	}
}

func TestDeleteCollection(t *testing.T) {
	l := New(nil)

	// collection doesn't exists
	err := l.DeleteCollection(10)
	expectingError(t, err, model.ErrNotFound)

	// collection delete
	err = l.DeleteCollection(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetCollectionByID(1)
	expectingError(t, err, model.ErrNotFound)
}

func TestListCollectionsByIDs(t *testing.T) {
	var nf, id int = 999, 1
	l := New(nil)

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"not founds", []*int{&nf, &nf, &nf}, 0},
		{"editor not found", []*int{&nf, nil, nil}, 0},
		{"print not found", []*int{nil, &nf, nil}, 0},
		{"series not found", []*int{nil, nil, &nf}, 0},
		{"editor id print not found series not found", []*int{&id, &nf, &nf}, 0},
		{"editor not found print id series not found", []*int{&nf, &id, &nf}, 0},
		{"editor not found print not found series id", []*int{&nf, &nf, &id}, 0},
		{"no ids", []*int{nil, nil, nil}, 1},
		{"only editor", []*int{&id, nil, nil}, 1},
		{"only print", []*int{nil, &id, nil}, 1},
		{"only series", []*int{nil, nil, &id}, 1},
		{"all", []*int{&id, &id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListCollectionsByIDs(tests[i].params[0], tests[i].params[1], tests[i].params[2])
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
			if !reflect.DeepEqual(bs[0], l.collections[1]) {
				t.Fatalf("unexpected collection value for [%s]", tests[i].desc)
			}
		}
	}
}
