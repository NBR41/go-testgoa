package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationEditorsPrintsControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsPrintsControllerListBooksByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListBooksByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsPrintsControllerListBooksByCollectionSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionSeriesRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456, CollectionID: 789, SeriesID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByCollectionSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsPrintsControllerListBooksBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456, SeriesID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsPrintsControllerListBooksBySeriesCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesCollectionRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456, CollectionID: 789, SeriesID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeriesCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsPrintsControllerListCollections(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListCollections(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsPrintsControllerListCollectionsBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsBySeriesRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456, SeriesID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollectionsBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsPrintsControllerListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsPrintsControllerListSeriesByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByCollectionRelationEditorsPrintsContext{Context: context.Background(), EditorID: 123, PrintID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID).Return(errTest),
	)
	err := ctrl.ListSeriesByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
