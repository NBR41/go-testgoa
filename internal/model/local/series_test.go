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

func TestListSeriesByIDs(t *testing.T) {
	l := New(nil)

	var nf, id int = 999, 1

	tests := []struct {
		desc   string
		params []*int
		exp    int
	}{
		{"not founds", []*int{&nf, &nf, &nf, &nf}, 0},
		{"author not found", []*int{&nf, nil, nil, nil}, 0},
		{"category not found", []*int{nil, &nf, nil, nil}, 0},
		{"class not found", []*int{nil, nil, &nf, nil}, 0},
		{"role not found", []*int{nil, nil, nil, &nf}, 0},
		{"author, with category not found", []*int{&id, &nf, nil, nil}, 0},
		{"author, with class not found", []*int{&id, nil, &nf, nil}, 0},
		{"author, with role not found", []*int{&id, nil, nil, &nf}, 0},
		{"category, with author not found", []*int{&nf, &id, nil, nil}, 0},
		{"category, with class not found", []*int{nil, &id, &nf, nil}, 0},
		{"category, with role not found", []*int{nil, &id, nil, &nf}, 0},
		{"class, with author not found", []*int{&nf, nil, &id, nil}, 0},
		{"class, with category not found", []*int{nil, &nf, &id, nil}, 0},
		{"class, with role not found", []*int{nil, nil, &id, &nf}, 0},
		{"role, with author not found", []*int{&nf, nil, nil, &id}, 0},
		{"role, with category not found", []*int{nil, &nf, nil, &id}, 0},
		{"role, with class not found", []*int{nil, nil, &nf, &id}, 0},
		{"no ids", []*int{nil, nil, nil, nil}, 1},
		{"author", []*int{&id, nil, nil, nil}, 1},
		{"category", []*int{nil, &id, nil, nil}, 1},
		{"class", []*int{nil, nil, &id, nil}, 1},
		{"role", []*int{nil, nil, nil, &id}, 1},
		{"all", []*int{&id, &id, &id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListSeriesByIDs(tests[i].params[0], tests[i].params[1], tests[i].params[2], tests[i].params[3])
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
			if !reflect.DeepEqual(bs[0], l.series[1]) {
				t.Fatalf("unexpected series value for [%s]", tests[i].desc)
			}
		}
	}
}

func TestListSeriesByEditionIDs(t *testing.T) {
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
		{"print not found", []*int{nil, nil, &nf}, 0},
		{"collection id editor not found print not found", []*int{&id, &nf, &nf}, 0},
		{"collection not found editor id print not found", []*int{&nf, &id, &nf}, 0},
		{"collection not found editor not found print id", []*int{&nf, &nf, &id}, 0},
		{"no ids", []*int{nil, nil, nil}, 1},
		{"only collection", []*int{&id, nil, nil}, 1},
		{"only editor", []*int{nil, &id, nil}, 1},
		{"only print", []*int{nil, nil, &id}, 1},
		{"all", []*int{&id, &id, &id}, 1},
	}

	for i := range tests {
		bs, err := l.ListSeriesByEditionIDs(tests[i].params[0], tests[i].params[1], tests[i].params[2])
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
			if !reflect.DeepEqual(bs[0], l.series[1]) {
				t.Fatalf("unexpected series value for [%s]", tests[i].desc)
			}
		}
	}
}
