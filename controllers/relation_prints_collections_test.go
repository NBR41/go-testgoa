package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationPrintsCollectionsControllerListBooks(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksRelationPrintsCollectionsContext{Context: context.Background(), CollectionID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID, nil).Return(errTest),
	)
	err := ctrl.ListBooks(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsCollectionsControllerListBooksBySeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListBooksBySeriesRelationPrintsCollectionsContext{Context: context.Background(), CollectionID: 123, PrintID: 456, SeriesID: 789}
	gomock.InOrder(
		lmock.EXPECT().ListBooks(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID, &ctx.SeriesID).Return(errTest),
	)
	err := ctrl.ListBooksBySeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationPrintsCollectionsControllerListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationPrintsCollectionsController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationPrintsCollectionsContext{Context: context.Background(), CollectionID: 123, PrintID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeriesByEditionIDs(ctx, nil, ctx, &ctx.CollectionID, nil, &ctx.PrintID).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
