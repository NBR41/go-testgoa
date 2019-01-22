package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationCollecionListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationCollectionContext{Context: context.Background(), CollectionID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, nil, nil).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCollecionListBooksByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintRelationCollectionContext{Context: context.Background(), CollectionID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListBooksByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCollecionListBooksByPrintSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintSeriesRelationCollectionContext{Context: context.Background(), CollectionID: 123, PrintID: 456, SeriesID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByPrintSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCollecionListBooksBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesRelationCollectionContext{Context: context.Background(), CollectionID: 123, SeriesID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCollecionListBooksBySeriesPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesPrintRelationCollectionContext{Context: context.Background(), CollectionID: 123, PrintID: 456, SeriesID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeriesPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCollecionListPrints(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsRelationCollectionContext{Context: context.Background(), CollectionID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, &ctx.CollectionID, nil, nil).Return(errTest),
	)
	err := ctrl.ListPrints(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCollecionListPrintsBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsBySeriesRelationCollectionContext{Context: context.Background(), CollectionID: 123, SeriesID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListPrintsBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCollecionListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationCollectionContext{Context: context.Background(), CollectionID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, &ctx.CollectionID, nil, nil).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCollecionListSeriesByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCollectionController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByPrintRelationCollectionContext{Context: context.Background(), CollectionID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID).Return(errTest),
	)
	err := ctrl.ListSeriesByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
