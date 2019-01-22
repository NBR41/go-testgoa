package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestRelationRoleListAuthors(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationRoleController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListAuthorsRelationRoleContext{Context: context.Background(), RoleID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListAuthors(ctx, nil, ctx, nil, &ctx.RoleID).Return(errTest),
	)
	err := ctrl.ListAuthors(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationRoleListSeriesByAuthors(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationRoleController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByAuthorRelationRoleContext{Context: context.Background(), AuthorID: 123, RoleID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, &ctx.AuthorID, nil, nil, &ctx.RoleID).Return(errTest),
	)
	err := ctrl.ListSeriesByAuthor(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
