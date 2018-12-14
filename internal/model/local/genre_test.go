package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetGenreByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetGenreByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetGenreByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetGenreByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetGenreByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetGenreByName("genre1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestListGenres(t *testing.T) {
	l := New(nil)

	bs, err := l.ListGenres()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertGenre(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertGenre("genre2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetGenreByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate genre name
	b, err = l.InsertGenre("genre2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateGenre(t *testing.T) {
	l := New(nil)

	// genre doesn't exist
	err := l.UpdateGenre(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//update genre
	err = l.UpdateGenre(1, "genre2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetGenreByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "genre2" {
		t.Fatalf("expecting genre2, got %s", u.Name)
	}
}

func TestDeleteGenre(t *testing.T) {
	l := New(nil)

	// genre doesn't exists
	err := l.DeleteGenre(10)
	expectingError(t, err, model.ErrNotFound)

	// genre delete
	err = l.DeleteGenre(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetGenreByID(1)
	expectingError(t, err, model.ErrNotFound)
}
