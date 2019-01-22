package controllers

import (
	"context"
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/app/test"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
	"github.com/kylelemons/godebug/pretty"
)

func TestRelationAuthorListCategories(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationAuthorController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListCategoriesRelationAuthorContext{Context: context.Background(), AuthorID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListCategories(ctx, nil, ctx, &ctx.AuthorID, nil).Return(errTest),
	)
	err := ctrl.ListCategories(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationAuthorListClasses(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationAuthorController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListClassesRelationAuthorContext{Context: context.Background(), AuthorID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListClasses(ctx, nil, ctx, &ctx.AuthorID, nil, nil).Return(errTest),
	)
	err := ctrl.ListClasses(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationAuthorListRoles(t *testing.T) {
	m1, m2 := &model.Role{ID: 1, Name: "foo"}, &model.Role{ID: 2, Name: "bar"}
	var authorID int = 123
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	lmock := NewMockLister(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(authorID).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(authorID).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(authorID).Return(nil, nil),
		mock.EXPECT().ListRolesByIDs(&authorID).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(authorID).Return(nil, nil),
		mock.EXPECT().ListRolesByIDs(&authorID).Return([]*model.Role{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), lmock)

	logbuf.Reset()
	test.ListRolesRelationAuthorServiceUnavailable(t, ctx, service, ctrl, authorID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}), lmock)

	logbuf.Reset()
	test.ListRolesRelationAuthorNotFound(t, ctx, service, ctrl, authorID)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListRolesRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListRolesRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID)
	exp = "[EROR] failed to get role list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListRolesRelationAuthorOK(t, ctx, service, ctrl, authorID)
	expres := app.RoleCollection{convert.ToRoleMedia(m1), convert.ToRoleMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationAuthorListSeries(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationAuthorController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesRelationAuthorContext{Context: context.Background(), AuthorID: 123}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, &ctx.AuthorID, nil, nil, nil).Return(errTest),
	)
	err := ctrl.ListSeries(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationAuthorListSeriesByCategory(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationAuthorController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByCategoryRelationAuthorContext{Context: context.Background(), AuthorID: 123, CategoryID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, &ctx.AuthorID, &ctx.CategoryID, nil, nil).Return(errTest),
	)
	err := ctrl.ListSeriesByCategory(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationAuthorListSeriesByClass(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationAuthorController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByClassRelationAuthorContext{Context: context.Background(), AuthorID: 123, ClassID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, &ctx.AuthorID, nil, &ctx.ClassID, nil).Return(errTest),
	)
	err := ctrl.ListSeriesByClass(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}

func TestRelationAuthorListSeriesByRole(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	lmock := NewMockLister(mctrl)
	errTest := errors.New("model error")
	ctrl := NewRelationAuthorController(goa.New("my-inventory-test"), nil, lmock)
	ctx := &app.ListSeriesByRoleRelationAuthorContext{Context: context.Background(), AuthorID: 123, RoleID: 456}
	gomock.InOrder(
		lmock.EXPECT().ListSeries(ctx, nil, ctx, &ctx.AuthorID, nil, nil, &ctx.RoleID).Return(errTest),
	)
	err := ctrl.ListSeriesByRole(ctx)
	if err != errTest {
		t.Errorf("unexpected error, exp [%v] got [%v]", errTest, err)
	}
}
