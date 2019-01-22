package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationClassListCategories(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationClassController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCategoriesRelationClassContext{Context: context.Background(), ClassID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListCategories(ctx, nil, ctx, nil, &ctx.ClassID).Return(errTest),
	)
	err := ctrl.ListCategories(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationClassListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationClassController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationClassContext{Context: context.Background(), ClassID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, nil, nil, &ctx.ClassID, nil).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationClassListSeriesByCategory(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationClassController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByCategoryRelationClassContext{Context: context.Background(), CategoryID: 456, ClassID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, nil, &ctx.CategoryID, &ctx.ClassID, nil).Return(errTest),
	)
	err := ctrl.ListSeriesByCategory(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
