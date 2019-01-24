package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetEditorByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetEditorByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetEditorByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetEditorByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetEditorByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetEditorByName("editor1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestListEditors(t *testing.T) {
	l := New(nil)

	bs, err := l.ListEditors()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestInsertEditor(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertEditor("editor2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetEditorByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate editor name
	b, err = l.InsertEditor("editor2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateEditor(t *testing.T) {
	l := New(nil)

	// editor doesn't exist
	err := l.UpdateEditor(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//duplicate name
	err = l.UpdateEditor(1, "editor1")
	switch err {
	case nil:
		t.Fatal("expecting error", err)
	case model.ErrDuplicateKey:
	default:
		t.Fatal("unexpected error", err)
	}

	//update editor
	err = l.UpdateEditor(1, "editor2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetEditorByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "editor2" {
		t.Fatalf("expecting editor2, got %s", u.Name)
	}
}

func TestDeleteEditor(t *testing.T) {
	l := New(nil)

	// editor doesn't exists
	err := l.DeleteEditor(10)
	expectingError(t, err, model.ErrNotFound)

	// editor delete
	err = l.DeleteEditor(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetEditorByID(1)
	expectingError(t, err, model.ErrNotFound)
}

func TestListEditorsByIDs(t *testing.T) {
	var nf, id int = 999, 1
	l := New(nil)

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"not founds", []*int{&nf, &nf}, 0},
		{"print not found", []*int{&nf, nil}, 0},
		{"series not found", []*int{nil, &nf}, 0},
		{"print id series not found", []*int{&id, &nf}, 0},
		{"print not found series id", []*int{&nf, &id}, 0},
		{"no ids", []*int{nil, nil}, 1},
		{"only print", []*int{&id, nil}, 1},
		{"only series", []*int{nil, &id}, 1},
		{"print and series", []*int{&id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListEditorsByIDs(tests[i].params[0], tests[i].params[1])
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
			if !reflect.DeepEqual(bs[0], l.editors[1]) {
				t.Fatalf("unexpected editor value for [%s]", tests[i].desc)
			}
		}
	}
}
