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

func TestRelationRoleListAuthors(t *testing.T) {
	m1, m2 := &model.Author{ID: 1, Name: "foo"}, &model.Author{ID: 2, Name: "bar"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetRoleByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetRoleByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetRoleByID(123).Return(nil, nil),
		mock.EXPECT().ListAuthorsByRoleID(123).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetRoleByID(123).Return(nil, nil),
		mock.EXPECT().ListAuthorsByRoleID(123).Return([]*model.Author{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationRoleController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListAuthorsRelationRoleServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationRoleController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListAuthorsRelationRoleNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get role error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListAuthorsRelationRoleInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get role error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListAuthorsRelationRoleInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListAuthorsRelationRoleOK(t, ctx, service, ctrl, 123)
	expres := app.AuthorCollection{convert.ToAuthorMedia(m1), convert.ToAuthorMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationRoleListSeriesByAuthors(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	authorID := 123
	roleID := 456
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetRoleByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetRoleByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetRoleByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, &roleID, nil, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetRoleByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, &roleID, nil, nil).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationRoleController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesByAuthorRelationRoleServiceUnavailable(t, ctx, service, ctrl, roleID, authorID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationRoleController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesByAuthorRelationRoleNotFound(t, ctx, service, ctrl, roleID, authorID)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByAuthorRelationRoleInternalServerError(t, ctx, service, ctrl, roleID, authorID)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByAuthorRelationRoleNotFound(t, ctx, service, ctrl, roleID, authorID)
	exp = "[EROR] failed to get role error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByAuthorRelationRoleInternalServerError(t, ctx, service, ctrl, roleID, authorID)
	exp = "[EROR] failed to get role error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByAuthorRelationRoleInternalServerError(t, ctx, service, ctrl, roleID, authorID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesByAuthorRelationRoleOK(t, ctx, service, ctrl, roleID, authorID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
