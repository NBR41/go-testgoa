package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestInsertBook(t *testing.T) {
	l := New(nil)

	//duplicate ISBN
	_, err := l.InsertBook("isbn-123", "foo", 1)
	expectingError(t, err, model.ErrDuplicateKey)
	//duplicate ISBN
	_, err = l.InsertBook("isbn-foo", "test1", 1)
	expectingError(t, err, model.ErrDuplicateKey)
	// invalid series id
	_, err = l.InsertBook("isbn-foo", "foo", 999)
	expectingError(t, err, model.ErrInvalidID)

	// insert
	b, err := l.InsertBook("isbn-foo", "foo", 1)
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

func TestListBooks(t *testing.T) {
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
	n1 := "test10"
	n2 := "test20"
	s1 := 999
	s2 := int(l.books[1].SeriesID)
	// book doesn't exist
	err := l.UpdateBook(10, &n1, nil)
	expectingError(t, err, model.ErrNotFound)

	//duplicate book name
	err = l.UpdateBook(1, &l.books[1].Name, nil)
	expectingError(t, err, model.ErrDuplicateKey)
	//invalid id
	err = l.UpdateBook(1, &n1, &s1)
	expectingError(t, err, model.ErrInvalidID)

	err = l.UpdateBook(1, &n2, &s2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	u, err := l.GetBookByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "test20" {
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

func TestListBooksByIDs(t *testing.T) {
	l := New(nil)
	collection1, print1, series1 := 999, 999, 999
	collection2, print2, series2 := 1, 1, 1
	//empty list
	li, err := l.ListBooksByIDs(&collection1, &print1, &series1)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 0 {
			t.Fatal("unexpected value")
		}
	}

	//valid list
	li, err = l.ListBooksByIDs(&collection2, &print2, &series2)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 1 {
			t.Fatal("unexpected value")
		}
		if li[0] != l.books[1] {
			t.Fatal("unexpected value")
		}
	}

	//valid list
	li, err = l.ListBooksByIDs(nil, nil, &series2)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 1 {
			t.Fatal("unexpected value")
		}
		if li[0] != l.books[1] {
			t.Fatal("unexpected value")
		}
	}
}
