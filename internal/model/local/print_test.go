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
	var nf, id int = 999, 1
	l := New(nil)

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"not founds", []*int{&nf, &nf, &nf}, 0},
		{"collection not found", []*int{&nf, nil, nil}, 0},
		{"editor not found", []*int{nil, &nf, nil}, 0},
		{"series not found", []*int{nil, nil, &nf}, 0},
		{"collection id editor not found series not found", []*int{&id, &nf, &nf}, 0},
		{"collection not found editor id series not found", []*int{&nf, &id, &nf}, 0},
		{"collection not found editor not found series id", []*int{&nf, &nf, &id}, 0},
		{"no ids", []*int{nil, nil, nil}, 1},
		{"only collection", []*int{&id, nil, nil}, 1},
		{"only editor", []*int{nil, &id, nil}, 1},
		{"only series", []*int{nil, nil, &id}, 1},
		{"all", []*int{&id, &id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListPrintsByIDs(tests[i].params[0], tests[i].params[1], tests[i].params[2])
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
			if !reflect.DeepEqual(bs[0], l.prints[1]) {
				t.Fatalf("unexpected print value for [%s]", tests[i].desc)
			}
		}
	}
}
