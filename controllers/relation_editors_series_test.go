package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationEditorsSeriesControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsSeriesControllerListBooksByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsSeriesControllerListBooksByCollectionPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionPrintRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456, CollectionID: 789, PrintID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByCollectionPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsSeriesControllerListBooksByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456, PrintID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsSeriesControllerListBooksByPrintCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintCollectionRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456, PrintID: 789, CollectionID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByPrintCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsSeriesControllerListCollections(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollections(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsSeriesControllerListCollectionsByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsByPrintRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456, PrintID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollectionsByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsSeriesControllerListPrints(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListPrints(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsSeriesControllerListPrintsByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsByCollectionRelationEditorsSeriesContext{Context: context.Background(), EditorID: 123, SeriesID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListPrintsByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
