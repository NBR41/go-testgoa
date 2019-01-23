package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationPrintsSeriesControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationPrintsSeriesContext{Context: context.Background(), PrintID: 123, SeriesID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsSeriesControllerListBooksByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionRelationPrintsSeriesContext{Context: context.Background(), PrintID: 123, SeriesID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsSeriesControllerListBooksByCollectionEditor(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionEditorRelationPrintsSeriesContext{Context: context.Background(), PrintID: 123, SeriesID: 456, CollectionID: 789, EditorID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByCollectionEditor(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsSeriesControllerListBooksByEditor(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByEditorRelationPrintsSeriesContext{Context: context.Background(), PrintID: 123, SeriesID: 456, EditorID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByEditor(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsSeriesControllerListBooksByEditorCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByEditorCollectionRelationPrintsSeriesContext{Context: context.Background(), PrintID: 123, SeriesID: 456, EditorID: 789, CollectionID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByEditorCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsSeriesControllerListCollections(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsRelationPrintsSeriesContext{Context: context.Background(), PrintID: 123, SeriesID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollections(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsSeriesControllerListCollectionsByEditor(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsByEditorRelationPrintsSeriesContext{Context: context.Background(), PrintID: 123, SeriesID: 456, EditorID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollectionsByEditor(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsSeriesControllerListEditors(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListEditorsRelationPrintsSeriesContext{Context: context.Background(), PrintID: 123, SeriesID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListEditors(ctx, nil, ctx, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListEditors(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsSeriesControllerListEditorsByCollection(t *testing.T) {
}
