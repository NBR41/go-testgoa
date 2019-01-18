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

func TestListAuthorsByCategoryID(t *testing.T) {
	l := New(nil)

	//empty list
	li, err := l.ListAuthorsByCategoryID(999)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 0 {
			t.Fatal("unexpected value")
		}
	}

	//valid list
	li, err = l.ListAuthorsByCategoryID(1)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 1 {
			t.Fatal("unexpected value")
		}
		if li[0] != l.authors[1] {
			t.Fatal("unexpected value")
		}
	}
}

func TestListAuthorsByRoleID(t *testing.T) {
	l := New(nil)

	//empty list
	li, err := l.ListAuthorsByRoleID(999)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 0 {
			t.Fatal("unexpected value")
		}
	}

	//valid list
	li, err = l.ListAuthorsByRoleID(1)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 1 {
			t.Fatal("unexpected value")
		}
		if li[0] != l.authors[1] {
			t.Fatal("unexpected value")
		}
	}
}
