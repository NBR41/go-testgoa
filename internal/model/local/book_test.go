package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestInsertBook(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertBook("isbn-foo", "foo")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 5 {
		t.Fatalf("expecting ID 5, got %v", b.ID)
	}
	b2, err := l.GetBookByID(5)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate book name
	b, err = l.InsertBook("isbn-foo", "foo")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestGetBookByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetBookByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetBookByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetBookByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetBookByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetBookByName("test1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetBookByISBN(t *testing.T) {
	l := New(nil)

	b, err := l.GetBookByISBN("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetBookByISBN("isbn-123")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetBookList(t *testing.T) {
	l := New(nil)

	bs, err := l.ListBooks()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestUpdateBook(t *testing.T) {
	l := New(nil)

	// user doesn't exist
	err := l.UpdateBook(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	err = l.UpdateBook(1, "test10")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	u, err := l.GetBookByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "test10" {
		t.Fatalf("expecting test10, got %s", u.Name)
	}
}

func TestDeleteBook(t *testing.T) {
	l := New(nil)

	// book doesn't exists
	err := l.DeleteBook(10)
	expectingError(t, err, model.ErrNotFound)

	// book delete
	err = l.DeleteBook(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetBookByID(1)
	expectingError(t, err, model.ErrNotFound)
}
