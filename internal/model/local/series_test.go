package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetSeriesByID(t *testing.T) {
	l := New(nil)

	b, err := l.GetSeriesByID(5)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetSeriesByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestGetSeriesByName(t *testing.T) {
	l := New(nil)

	b, err := l.GetSeriesByName("test5")
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetSeriesByName("series1")
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestListSeries(t *testing.T) {
	l := New(nil)

	bs, err := l.ListSeries()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for i := range bs {
		if int64(i)+1 != bs[i].ID {
			t.Fatal("unexpected ID , list must be sorted")
		}
	}
}

func TestListSeriesByIDs(t *testing.T) {
	l := New(nil)
	author1, role1, category1, class1 := 999, 999, 999, 999
	author2, role2, category2, class2 := 1, 1, 1, 1
	//empty list
	li, err := l.ListSeriesByIDs(&author1, &role1, &category1, &class1)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 0 {
			t.Fatal("unexpected value")
		}
	}

	//valid list
	li, err = l.ListSeriesByIDs(&author2, &role2, &category2, nil)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 1 {
			t.Fatal("unexpected value")
		}
		if li[0] != l.series[1] {
			t.Fatal("unexpected value")
		}
	}

	//valid list
	li, err = l.ListSeriesByIDs(&author2, &role2, &category2, &class2)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 1 {
			t.Fatal("unexpected value")
		}
		if li[0] != l.series[1] {
			t.Fatal("unexpected value")
		}
	}
}

func TestListSeriesByCollectionID(t *testing.T) {
	l := New(nil)

	//empty list
	li, err := l.ListSeriesByCollectionID(999)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 0 {
			t.Fatal("unexpected value")
		}
	}

	//valid list
	li, err = l.ListSeriesByCollectionID(1)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 1 {
			t.Fatal("unexpected value")
		}
		if li[0] != l.series[1] {
			t.Fatal("unexpected value")
		}
	}
}

func TestListSeriesByCollectionIDPrintID(t *testing.T) {
	l := New(nil)

	//empty list
	li, err := l.ListSeriesByCollectionIDPrintID(999, 999)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 0 {
			t.Fatal("unexpected value")
		}
	}

	//valid list
	li, err = l.ListSeriesByCollectionIDPrintID(1, 1)
	if err != nil {
		t.Fatalf("unexpected error [%v]", err)
	} else {
		if len(li) != 1 {
			t.Fatal("unexpected value")
		}
		if li[0] != l.series[1] {
			t.Fatal("unexpected value")
		}
	}
}

func TestInsertSeries(t *testing.T) {
	l := New(nil)

	_, err := l.InsertSeries("series1", 3)
	expectingError(t, err, model.ErrDuplicateKey)
	// editor not found
	_, err = l.InsertSeries("series2", 999)
	expectingError(t, err, model.ErrInvalidID)

	// insert
	b, err := l.InsertSeries("series2", 1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetSeriesByID(2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}
}

func TestUpdateSeries(t *testing.T) {
	l := New(nil)
	n0 := "test10"
	n1 := "series1"
	n2 := "series2"
	n3 := "series3"
	e1 := 5
	e, err := l.InsertCategory("category2")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	e2 := int(e.ID)

	// series doesn't exist
	err = l.UpdateSeries(10, &n0, nil)
	expectingError(t, err, model.ErrNotFound)
	//same name
	err = l.UpdateSeries(1, &n1, &e2)
	expectingError(t, err, model.ErrDuplicateKey)
	t.Log(n3)

	// category doesn't exist
	err = l.UpdateSeries(1, &n2, &e1)
	expectingError(t, err, model.ErrInvalidID)

	//update series
	err = l.UpdateSeries(1, &n3, &e2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	u, err := l.GetSeriesByID(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if u.Name != n3 {
		t.Fatalf("expecting %s, got %s", n3, u.Name)
	}
	if u.CategoryID != e.ID {
		t.Fatalf("expecting category id %d, got %d", e.ID, u.CategoryID)
	}
}

func TestDeleteSeries(t *testing.T) {
	l := New(nil)

	// series doesn't exists
	err := l.DeleteSeries(10)
	expectingError(t, err, model.ErrNotFound)

	// series delete
	err = l.DeleteSeries(1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetSeriesByID(1)
	expectingError(t, err, model.ErrNotFound)
}
