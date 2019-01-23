package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationSeriesCollectionsControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationSeriesCollectionsContext{Context: context.Background(), SeriesID: 123, CollectionID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesCollectionsControllerListBooksByPrint(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksByPrintRelationSeriesCollectionsContext{Context: context.Background(), SeriesID: 123, CollectionID: 456, PrintID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksByPrint(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesCollectionsControllerListPrints(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsRelationSeriesCollectionsContext{Context: context.Background(), SeriesID: 123, CollectionID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListPrints(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
