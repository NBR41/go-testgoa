package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetAuthorByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetAuthorByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetAuthorByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetAuthorByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetAuthorByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetAuthorByName("author1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestListAuthors(t *testing.T) {
	l := New(nil)

	bs, err := l.ListAuthors()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertAuthor(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertAuthor("author2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetAuthorByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate author name
	b, err = l.InsertAuthor("author2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateAuthor(t *testing.T) {
	l := New(nil)

	// author doesn't exist
	err := l.UpdateAuthor(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//duplicate name
	err = l.UpdateAuthor(1, "author1")
	switch err {
	case nil:
		t.Fatal("expecting error", err)
	case model.ErrDuplicateKey:
	default:
		t.Fatal("unexpected error", err)
	}

	//update author
	err = l.UpdateAuthor(1, "author2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetAuthorByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "author2" {
		t.Fatalf("expecting author2, got %s", u.Name)
	}
}

func TestDeleteAuthor(t *testing.T) {
	l := New(nil)

	// author doesn't exists
	err := l.DeleteAuthor(10)
	expectingError(t, err, model.ErrNotFound)

	// author delete
	err = l.DeleteAuthor(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetAuthorByID(1)
	expectingError(t, err, model.ErrNotFound)
}

func TestListAuthorsByIDs(t *testing.T) {
	l := New(nil)
	var nf, id int = 999, 1

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"not founds", []*int{&nf, &nf}, 0},
		{"category not found", []*int{&nf, nil}, 0},
		{"role not found", []*int{nil, &nf}, 0},
		{"category id role not found", []*int{&id, &nf}, 0},
		{"category not found role id", []*int{&nf, &id}, 0},
		{"no ids", []*int{nil, nil}, 1},
		{"only category", []*int{&id, nil}, 1},
		{"only role", []*int{nil, &id}, 1},
		{"category and role", []*int{&id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListAuthorsByIDs(tests[i].params[0], tests[i].params[1])
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
			if !reflect.DeepEqual(bs[0], l.authors[1]) {
				t.Fatalf("unexpected author value for [%s]", tests[i].desc)
			}
		}
	}
}
