package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationEditorsCollectionsControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, nil, nil).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsCollectionsControllerListBooksByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456, PrintID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListBooksByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsCollectionsControllerListBooksByPrintSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintSeriesRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456, PrintID: 789, SeriesID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByPrintSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsCollectionsControllerListBooksBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456, SeriesID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsCollectionsControllerListBooksBySeriesPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesPrintRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456, PrintID: 789, SeriesID: 321}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeriesPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsCollectionsControllerListPrints(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, nil).Return(errTest),
	)
	err := ctrl.ListPrints(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsCollectionsControllerListPrintsBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsBySeriesRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456, SeriesID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListPrintsBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsCollectionsControllerListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, nil).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationEditorsCollectionsControllerListSeriesByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationEditorsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByPrintRelationEditorsCollectionsContext{Context: context.Background(), CollectionID: 123, EditorID: 456, PrintID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID).Return(errTest),
	)
	err := ctrl.ListSeriesByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
