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

	var nf, id int = 999, 1

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"not founds", []*int{&nf, &nf, &nf, &nf}, 0},
		{"collection not found", []*int{&nf, nil, nil, nil}, 0},
		{"editor not found", []*int{nil, &nf, nil, nil}, 0},
		{"print not found", []*int{nil, nil, &nf, nil}, 0},
		{"series not found", []*int{nil, nil, nil, &nf}, 0},
		{"collection, with editor not found", []*int{&id, &nf, nil, nil}, 0},
		{"collection, with print not found", []*int{&id, nil, &nf, nil}, 0},
		{"collection, with series not found", []*int{&id, nil, nil, &nf}, 0},
		{"editor, with collection not found", []*int{&nf, &id, nil, nil}, 0},
		{"editor, with print not found", []*int{nil, &id, &nf, nil}, 0},
		{"editor, with series not found", []*int{nil, &id, nil, &nf}, 0},
		{"print, with collection not found", []*int{&nf, nil, &id, nil}, 0},
		{"print, with editor not found", []*int{nil, &nf, &id, nil}, 0},
		{"print, with series not found", []*int{nil, nil, &id, &nf}, 0},
		{"series, with collection not found", []*int{&nf, nil, nil, &id}, 0},
		{"series, with editor not found", []*int{nil, &nf, nil, &id}, 0},
		{"series, with print not found", []*int{nil, nil, &nf, &id}, 0},
		{"no ids", []*int{nil, nil, nil, nil}, 1},
		{"collection", []*int{&id, nil, nil, nil}, 1},
		{"editor", []*int{nil, &id, nil, nil}, 1},
		{"print", []*int{nil, nil, &id, nil}, 1},
		{"series", []*int{nil, nil, nil, &id}, 1},
		{"all", []*int{&id, &id, &id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListBooksByIDs(tests[i].params[0], tests[i].params[1], tests[i].params[2], tests[i].params[3])
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
			if !reflect.DeepEqual(bs[0], l.books[1]) {
				t.Fatalf("unexpected book value for [%s]", tests[i].desc)
			}
		}
	}
}
