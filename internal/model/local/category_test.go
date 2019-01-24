package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetCategoryByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetCategoryByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetCategoryByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetCategoryByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetCategoryByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetCategoryByName("category1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestInsertCategory(t *testing.T) {
	l := New(nil)

	// insert
	b, err := l.InsertCategory("category2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetCategoryByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}

	// duplicate category name
	b, err = l.InsertCategory("category2")
	expectingError(t, err, model.ErrDuplicateKey)
}

func TestUpdateCategory(t *testing.T) {
	l := New(nil)

	// category doesn't exist
	err := l.UpdateCategory(10, "test10")
	expectingError(t, err, model.ErrNotFound)

	//duplicate name
	err = l.UpdateCategory(1, "category1")
	switch err {
	case nil:
		t.Fatal("expecting error", err)
	case model.ErrDuplicateKey:
	default:
		t.Fatal("unexpected error", err)
	}

	//update category
	err = l.UpdateCategory(1, "category2")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetCategoryByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != "category2" {
		t.Fatalf("expecting category2, got %s", u.Name)
	}
}

func TestDeleteCategory(t *testing.T) {
	l := New(nil)

	// category doesn't exists
	err := l.DeleteCategory(10)
	expectingError(t, err, model.ErrNotFound)

	// category delete
	err = l.DeleteCategory(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetCategoryByID(1)
	expectingError(t, err, model.ErrNotFound)
}

func TestListCategoriesByIDs(t *testing.T) {
	var nf, id int = 999, 1
	l := New(nil)

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"not founds", []*int{&nf, &nf}, 0},
		{"author not found", []*int{&nf, nil}, 0},
		{"class not found", []*int{nil, &nf}, 0},
		{"author id class not found", []*int{&id, &nf}, 0},
		{"author not found class id", []*int{&nf, &id}, 0},
		{"no ids", []*int{nil, nil}, 1},
		{"only author", []*int{&id, nil}, 1},
		{"only class", []*int{nil, &id}, 1},
		{"author and class", []*int{&id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListCategoriesByIDs(tests[i].params[0], tests[i].params[1])
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
			if !reflect.DeepEqual(bs[0], l.categories[1]) {
				t.Fatalf("unexpected category value for [%s]", tests[i].desc)
			}
		}
	}
}
