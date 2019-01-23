package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationPrintsEditorsControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsEditorsControllerListBooksByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListBooksByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsEditorsControllerListBooksByCollectionSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionSeriesRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456, CollectionID: 789, SeriesID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByCollectionSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsEditorsControllerListBooksBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456, SeriesID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsEditorsControllerListBooksBySeriesCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesCollectionRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456, SeriesID: 789, CollectionID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeriesCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}

}

func TestRelationPrintsEditorsControllerListCollections(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListCollections(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsEditorsControllerListCollectionsBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsBySeriesRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456, SeriesID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollectionsBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsEditorsControllerListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsEditorsControllerListSeriesByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByCollectionRelationPrintsEditorsContext{Context: context.Background(), EditorID: 123, PrintID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID).Return(errTest),
	)
	err := ctrl.ListSeriesByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
