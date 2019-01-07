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

func TestListClasses(t *testing.T) {
	l := New(nil)

	bs, err := l.ListClasses()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
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
