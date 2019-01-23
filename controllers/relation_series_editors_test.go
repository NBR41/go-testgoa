package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationSeriesEditorsControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesEditorsControllerListBooksByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesEditorsControllerListBooksByCollectionPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByCollectionPrintRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456, CollectionID: 789, PrintID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByCollectionPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesEditorsControllerListBooksByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456, PrintID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesEditorsControllerListBooksByPrintCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintCollectionRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456, PrintID: 789, CollectionID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByPrintCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesEditorsControllerListCollections(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollections(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesEditorsControllerListCollectionsByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsByPrintRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456, PrintID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollectionsByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesEditorsControllerListPrints(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, nil, &ctx.EditorID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListPrints(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesEditorsControllerListPrintsByCollection(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesEditorsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsByCollectionRelationSeriesEditorsContext{Context: context.Background(), SeriesID: 123, EditorID: 456, CollectionID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListPrintsByCollection(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
