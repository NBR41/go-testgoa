package local

import (
	"reflect"
	"testing"

	"github.com/NBR41/go-testgoa/internal/model"
)

func TestGetClassification(t *testing.T) {
	l := New(nil)

	b, err := l.GetClassification(999, 999)
	expectingError(t, err, model.ErrNotFound)

	b, err = l.GetClassification(1, 1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 1 {
		t.Fatalf("expecting ID 1, got %v", b.ID)
	}
}

func TestInsertClassification(t *testing.T) {
	l := New(nil)

	//invalid series id
	_, err := l.InsertClassification(999, 999)
	expectingError(t, err, model.ErrInvalidID)
	//invalid class id
	_, err = l.InsertClassification(1, 999)
	expectingError(t, err, model.ErrInvalidID)
	//duplicate key
	_, err = l.InsertClassification(1, 1)
	expectingError(t, err, model.ErrDuplicateKey)

	_, err = l.InsertClass("class2")
	if err != nil {
		t.Fatal("unexpected error")
	}

	b, err := l.InsertClassification(1, 2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if b.ID != 2 {
		t.Fatalf("expecting ID 2, got %v", b.ID)
	}
	b2, err := l.GetClassification(1, 2)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if !reflect.DeepEqual(b, b2) {
		t.Fatal("unexpected user value")
	}
}

func TestDeleteClassification(t *testing.T) {
	l := New(nil)
	// classification doesn't exists
	err := l.DeleteClassification(999, 999)
	expectingError(t, err, model.ErrNotFound)
	err = l.DeleteClassification(1, 999)
	expectingError(t, err, model.ErrNotFound)

	// classification delete
	err = l.DeleteClassification(1, 1)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	_, err = l.GetClassification(1, 1)
	expectingError(t, err, model.ErrNotFound)
}
