package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationPrintsControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationPrintsContext{Context: context.Background(), PrintID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, nil, nil, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsControllerListCollections(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCollectionsRelationPrintsContext{Context: context.Background(), PrintID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListCollections(ctx, nil, ctx, nil, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListCollections(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsControllerListEditors(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListEditorsRelationPrintsContext{Context: context.Background(), PrintID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListEditors(ctx, nil, ctx, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListEditors(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsControllerListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationPrintsContext{Context: context.Background(), PrintID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, nil, nil, &ctx.PrintID).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
