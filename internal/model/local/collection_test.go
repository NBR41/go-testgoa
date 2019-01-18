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

	b, err := l.GetCollectionByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetCollectionByName("collection1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestInsertCollection(t *testing.T) {
	l := New(nil)

	_, err := l.InsertCollection("collection1", 3)
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

func TestListCollections(t *testing.T) {
	l := New(nil)

	bs, err := l.ListCollections()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestListCollectionsByEditorID(t *testing.T) {
	l := New(nil)
	bs, err := l.ListCollectionsByEditorID(5)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if len(bs) != 0 {
		t.Errorf("unexpected length, exp 0 got %d", len(bs))
	}

	bs, err = l.ListCollectionsByEditorID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if len(bs) != 1 {
		t.Errorf("unexpected length, exp 1 got %d", len(bs))
	}
}
