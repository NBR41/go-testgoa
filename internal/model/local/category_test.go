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

func TestGetCategoryList(t *testing.T) {
	l := New(nil)

	bs, err := l.ListCategories()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
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
