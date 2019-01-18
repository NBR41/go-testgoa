package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetEditorByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetEditorByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetEditorByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetEditorByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetEditorByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetEditorByName("editor1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestListEditors(t *testing.T) {
	l := New(nil)

	bs, err := l.ListEditors()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertEditor(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertEditor("editor2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetEditorByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate editor name
	b, err = l.InsertEditor("editor2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateEditor(t *testing.T) {
	l := New(nil)

	// editor doesn't exist
	err := l.UpdateEditor(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//duplicate name
	err = l.UpdateEditor(1, "editor1")
	switch err {
	case nil:
		t.Fatal("expecting error", err)
	case model.ErrDuplicateKey:
	default:
		t.Fatal("unexpected error", err)
	}

	//update editor
	err = l.UpdateEditor(1, "editor2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetEditorByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "editor2" {
		t.Fatalf("expecting editor2, got %s", u.Name)
	}
}

func TestDeleteEditor(t *testing.T) {
	l := New(nil)

	// editor doesn't exists
	err := l.DeleteEditor(10)
	expectingError(t, err, model.ErrNotFound)

	// editor delete
	err = l.DeleteEditor(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetEditorByID(1)
	expectingError(t, err, model.ErrNotFound)
}
