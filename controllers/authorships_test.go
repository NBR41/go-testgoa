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

func TestAuthorshipsCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertAuthorship(1, 2, 3).Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertAuthorship(1, 2, 3).Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertAuthorship(1, 2, 3).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertAuthorship(1, 2, 3).Return(&model.Authorship{ID: 4, AuthorID: 1, BookID: 2, RoleID: 3}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorshipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreateAuthorshipsServiceUnavailable(t, ctx, service, ctrl, &app.CreateAuthorshipsPayload{AuthorID: 1, BookID: 2, RoleID: 3})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorshipsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreateAuthorshipsInternalServerError(t, ctx, service, ctrl, &app.CreateAuthorshipsPayload{AuthorID: 1, BookID: 2, RoleID: 3})
	exp = "[EROR] failed to insert authorship error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateAuthorshipsUnprocessableEntity(t, ctx, service, ctrl, &app.CreateAuthorshipsPayload{AuthorID: 1, BookID: 2, RoleID: 3})
	exp = "[EROR] failed to insert authorship error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateAuthorshipsUnprocessableEntity(t, ctx, service, ctrl, &app.CreateAuthorshipsPayload{AuthorID: 1, BookID: 2, RoleID: 3})
	exp = "[EROR] failed to insert authorship error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateAuthorshipsCreated(t, ctx, service, ctrl, &app.CreateAuthorshipsPayload{AuthorID: 1, BookID: 2, RoleID: 3})
	exp = app.AuthorshipsHref(4)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestAuthorshipsDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeleteAuthorship(123).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteAuthorship(123).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteAuthorship(123).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorshipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeleteAuthorshipsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorshipsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeleteAuthorshipsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete authorship error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteAuthorshipsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete authorship error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteAuthorshipsNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestAuthorshipsList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().ListAuthorships().Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().ListAuthorships().Return([]*model.Authorship{&model.Authorship{ID: 1, AuthorID: 2, BookID: 3, RoleID: 4}, &model.Authorship{ID: 5, AuthorID: 6, BookID: 7, RoleID: 8}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorshipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListAuthorshipsServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorshipsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListAuthorshipsInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get authorship list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListAuthorshipsOK(t, ctx, service, ctrl)
	expres := app.AuthorshipCollection{
		convert.ToAuthorshipMedia(&model.Authorship{ID: 1, AuthorID: 2, BookID: 3, RoleID: 4}),
		convert.ToAuthorshipMedia(&model.Authorship{ID: 5, AuthorID: 6, BookID: 7, RoleID: 8}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestAuthorshipsShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorshipByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorshipByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorshipByID(123).Return(&model.Authorship{ID: 1, AuthorID: 2, BookID: 3, RoleID: 4}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorshipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowAuthorshipsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorshipsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowAuthorshipsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get authorship error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowAuthorshipsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get authorship error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowAuthorshipsOK(t, ctx, service, ctrl, 123)
	expres := convert.ToAuthorshipMedia(&model.Authorship{ID: 1, AuthorID: 2, BookID: 3, RoleID: 4})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
