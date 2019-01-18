package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetAuthorshipByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetAuthorshipByID(999)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetAuthorshipByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestListAuthorships(t *testing.T) {
	l := New(nil)

	bs, err := l.ListAuthorships()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertAuthorship(t *testing.T) {
	l := New(nil)

	// invalid author id
	_, err := l.InsertAuthorship(999, 999, 999)
	expectingError(t, err, model.ErrInvalidID)
	// invalid book id
	_, err = l.InsertAuthorship(1, 999, 999)
	expectingError(t, err, model.ErrInvalidID)
	//invalid role id
	_, err = l.InsertAuthorship(1, 1, 999)
	expectingError(t, err, model.ErrInvalidID)
	// duplicate key
	_, err = l.InsertAuthorship(1, 1, 1)
	expectingError(t, err, model.ErrDuplicateKey)

	_, err = l.InsertRole("role2")
	if err != nil {
		t.Fatal("unexpected error")
	}

	// insert
	b, err := l.InsertAuthorship(1, 1, 2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetAuthorshipByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}
}

func TestDeleteAuthorship(t *testing.T) {
	l := New(nil)

	// authorship doesn't exists
	err := l.DeleteAuthorship(999)
	expectingError(t, err, model.ErrNotFound)

	// authorship delete
	err = l.DeleteAuthorship(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetAuthorshipByID(1)
	expectingError(t, err, model.ErrNotFound)
}
