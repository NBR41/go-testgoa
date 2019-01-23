package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationSeriesPrintsControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationSeriesPrintsContext{Context: context.Background(), SeriesID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesPrintsControllerListBooksByEditor(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByEditorRelationSeriesPrintsContext{Context: context.Background(), SeriesID: 123, PrintID: 456, EditorID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByEditor(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesPrintsControllerListBooksByEditorCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByEditorCollectionRelationSeriesPrintsContext{Context: context.Background(), SeriesID: 123, PrintID: 456, EditorID: 789, CollectionID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByEditorCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesPrintsControllerListCollections(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsRelationSeriesPrintsContext{Context: context.Background(), SeriesID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollections(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesPrintsControllerListCollectionsByEditor(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsByEditorRelationSeriesPrintsContext{Context: context.Background(), SeriesID: 123, PrintID: 456, EditorID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollectionsByEditor(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesPrintsControllerListEditors(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListEditorsRelationSeriesPrintsContext{Context: context.Background(), SeriesID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListEditors(ctx, nil, ctx, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListEditors(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
