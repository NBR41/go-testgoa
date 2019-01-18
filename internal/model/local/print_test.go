package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetPrintByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetPrintByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetPrintByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetPrintByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetPrintByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetPrintByName("print1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestListPrints(t *testing.T) {
	l := New(nil)

	bs, err := l.ListPrints()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertPrint(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertPrint("print2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetPrintByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate print name
	b, err = l.InsertPrint("print2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdatePrint(t *testing.T) {
	l := New(nil)

	// print doesn't exist
	err := l.UpdatePrint(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//duplicate name
	err = l.UpdatePrint(1, "print1")
	switch err {
	case nil:
		t.Fatal("expecting error", err)
	case model.ErrDuplicateKey:
	default:
		t.Fatal("unexpected error", err)
	}

	//update print
	err = l.UpdatePrint(1, "print2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetPrintByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "print2" {
		t.Fatalf("expecting print2, got %s", u.Name)
	}
}

func TestDeletePrint(t *testing.T) {
	l := New(nil)

	// print doesn't exists
	err := l.DeletePrint(10)
	expectingError(t, err, model.ErrNotFound)

	// print delete
	err = l.DeletePrint(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetPrintByID(1)
	expectingError(t, err, model.ErrNotFound)
}

func TestListPrintsByIDs(t *testing.T) {
	l := New(nil)
	collection1, collection2, series1, series2 := 999, 1, 999, 1
	li, err := l.ListPrintsByIDs(&collection1, &series1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if len(li) != 0 {
		t.Errorf("unexpected length, exp 0 got %d", len(li))
	}

	li, err = l.ListPrintsByIDs(&collection2, &series2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if len(li) != 1 {
		t.Errorf("unexpected length, exp 1 got %d", len(li))
	} else {
		if li[0] != l.prints[1] {
			t.Fatal("unexpected value")
		}
	}

	li, err = l.ListPrintsByIDs(&collection2, nil)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if len(li) != 1 {
		t.Errorf("unexpected length, exp 1 got %d", len(li))
	} else {
		if li[0] != l.prints[1] {
			t.Fatal("unexpected value")
		}
	}

	li, err = l.ListPrintsByIDs(nil, &series2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if len(li) != 1 {
		t.Errorf("unexpected length, exp 1 got %d", len(li))
	} else {
		if li[0] != l.prints[1] {
			t.Fatal("unexpected value")
		}
	}
}
