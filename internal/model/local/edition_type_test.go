package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetEditionTypeByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetEditionTypeByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetEditionTypeByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetEditionTypeByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetEditionTypeByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetEditionTypeByName("edition_type1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetEditionTypeList(t *testing.T) {
	l := New(nil)

	bs, err := l.GetEditionTypeList()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertEditionType(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertEditionType("edition_type2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetEditionTypeByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate edition_type name
	b, err = l.InsertEditionType("edition_type2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateEditionType(t *testing.T) {
	l := New(nil)

	// edition_type doesn't exist
	err := l.UpdateEditionType(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//update edition_type
	err = l.UpdateEditionType(1, "edition_type2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetEditionTypeByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "edition_type2" {
		t.Fatalf("expecting edition_type2, got %s", u.Name)
	}
}

func TestDeleteEditionType(t *testing.T) {
	l := New(nil)

	// edition_type doesn't exists
	err := l.DeleteEditionType(10)
	expectingError(t, err, model.ErrNotFound)

	// edition_type delete
	err = l.DeleteEditionType(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetEditionTypeByID(1)
	expectingError(t, err, model.ErrNotFound)
}
