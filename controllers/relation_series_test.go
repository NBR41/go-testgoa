package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationSeriesControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationSeriesContext{Context: context.Background(), SeriesID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, nil, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesControllerListCollections(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsRelationSeriesContext{Context: context.Background(), SeriesID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, nil, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListCollections(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesControllerListEditors(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListEditorsRelationSeriesContext{Context: context.Background(), SeriesID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListEditors(ctx, nil, ctx, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListEditors(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationSeriesControllerListPrints(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationSeriesController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListPrintsRelationSeriesContext{Context: context.Background(), SeriesID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListPrints(ctx, nil, ctx, nil, nil, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListPrints(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
