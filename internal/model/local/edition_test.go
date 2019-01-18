package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetEditionByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetEditionByID(999)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetEditionByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestListEditions(t *testing.T) {
	l := New(nil)

	bs, err := l.ListEditions()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertEdition(t *testing.T) {
	l := New(nil)

	// invalid book id
	_, err := l.InsertEdition(999, 999, 999)
	expectingError(t, err, model.ErrInvalidID)
	// invalid collection id
	_, err = l.InsertEdition(1, 999, 999)
	expectingError(t, err, model.ErrInvalidID)
	//invalid print id
	_, err = l.InsertEdition(1, 1, 999)
	expectingError(t, err, model.ErrInvalidID)
	// duplicate key
	_, err = l.InsertEdition(1, 1, 1)
	expectingError(t, err, model.ErrDuplicateKey)

	_, err = l.InsertPrint("print2")
	if err != nil {
		t.Fatal("unexpected error")
	}

	// insert
	b, err := l.InsertEdition(1, 1, 2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetEditionByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}
}

func TestDeleteEdition(t *testing.T) {
	l := New(nil)

	// edition doesn't exists
	err := l.DeleteEdition(999)
	expectingError(t, err, model.ErrNotFound)

	// edition delete
	err = l.DeleteEdition(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetEditionByID(1)
	expectingError(t, err, model.ErrNotFound)
}
