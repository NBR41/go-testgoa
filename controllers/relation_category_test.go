package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationCategoryListAuthors(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCategoryController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListAuthorsRelationCategoryContext{Context: context.Background(), CategoryID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListAuthors(ctx, nil, ctx, &ctx.CategoryID, nil).Return(errTest),
	)
	err := ctrl.ListAuthors(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCategoryListClasses(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCategoryController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListClassesRelationCategoryContext{Context: context.Background(), CategoryID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListClasses(ctx, nil, ctx, nil, &ctx.CategoryID, nil).Return(errTest),
	)
	err := ctrl.ListClasses(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCategoryListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCategoryController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationCategoryContext{Context: context.Background(), CategoryID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, nil, &ctx.CategoryID, nil, nil).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationCategoryListSeriesByClass(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationCategoryController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByClassRelationCategoryContext{Context: context.Background(), CategoryID: 123, ClassID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, nil, &ctx.CategoryID, &ctx.ClassID, nil).Return(errTest),
	)
	err := ctrl.ListSeriesByClass(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
