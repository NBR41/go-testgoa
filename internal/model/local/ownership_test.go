package local

import (
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetOwnershipList(t *testing.T) {
	l := New(nil)

	_, err := l.ListOwnershipsByUserID(5)
	expectingError(t, err, model.ErrNotFound)

	ows, err := l.ListOwnershipsByUserID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	for i := range ows {
		if ows[i].UserID != 1 {
			t.Fatalf("unexpected user ID, expecting 1, got %d", ows[i].UserID)
		}

		if ows[i].BookID != l.ownerships[1][i].ID {
			t.Fatalf("unexpected user ID, expecting %d, got %d", l.ownerships[1][i].ID, ows[i].BookID)
		}
	}
}

func TestGetOwnership(t *testing.T) {
	l := New(nil)

	// not found
	o, err := l.GetOwnership(5, 5)
	expectingError(t, err, model.ErrNotFound)
	// user not found
	o, err = l.GetOwnership(5, 1)
	expectingError(t, err, model.ErrNotFound)
	// book not found
	o, err = l.GetOwnership(1, 5)
	expectingError(t, err, model.ErrNotFound)

	//  found
	o, err = l.GetOwnership(1, 1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if o.UserID != 1 || o.BookID != 1 || o.Book.ID != 1 {
		t.Fatal("unexpected value", o)
	}
}

func TestInsertOwnership(t *testing.T) {
	l := New(nil)

	// not found
	o, err := l.InsertOwnership(5, 5)
	expectingError(t, err, model.ErrNotFound)
	// user not found
	o, err = l.InsertOwnership(5, 1)
	expectingError(t, err, model.ErrNotFound)
	// book not found
	o, err = l.InsertOwnership(1, 5)
	expectingError(t, err, model.ErrNotFound)
	//  duplicate key
	o, err = l.InsertOwnership(1, 1)
	expectingError(t, err, model.ErrDuplicateKey)

	//  found
	o, err = l.InsertOwnership(1, 2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if o.UserID != 1 || o.BookID != 2 || o.Book.ID != 2 {
		t.Fatal("unexpected value", o)
	}
}

func TestUpdateOwnership(t *testing.T) {

}

func TestDeleteOwnership(t *testing.T) {
	l := New(nil)

	// not found
	err := l.DeleteOwnership(5, 5)
	expectingError(t, err, model.ErrNotFound)
	// user not found
	err = l.DeleteOwnership(5, 1)
	expectingError(t, err, model.ErrNotFound)
	// book not found
	err = l.DeleteOwnership(1, 5)
	expectingError(t, err, model.ErrNotFound)

	//  found
	err = l.DeleteOwnership(1, 1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	_, err = l.GetOwnership(1, 1)
	expectingError(t, err, model.ErrNotFound)
}
