package controllers

import (
	"context"
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/app/test"
	"github.com/NBR41/go-testgoa/internal/api"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
	"github.com/kylelemons/godebug/pretty"
)

func TestOwnershipsCreate(t *testing.T) {
	payload := &app.CreateOwnershipsPayload{BookID: 456}
	o := &model.Ownership{UserID: 123, BookID: 456}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	amock := NewMockAPIHelper(mctrl)
	gomock.InOrder(
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, errors.New("insert error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, model.ErrDuplicateKey),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, model.ErrInvalidID),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertOwnership(123, 456).Return(o, nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), amock)

	logbuf.Reset()
	test.CreateOwnershipsServiceUnavailable(t, ctx, service, ctrl, 123, payload)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), amock)

	logbuf.Reset()
	test.CreateOwnershipsInternalServerError(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert ownership error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateOwnershipsNotFound(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert ownership error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateOwnershipsUnprocessableEntity(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert ownership error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateOwnershipsUnprocessableEntity(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert ownership error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateOwnershipsCreated(t, ctx, service, ctrl, 123, payload)
	exp = app.OwnershipsHref(123, 456)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestOwnershipsAdd(t *testing.T) {
	payload := &app.AddOwnershipsPayload{BookIsbn: "baz"}
	book := &model.Book{ID: 456, ISBN: "baz", Name: "bar"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	amock := NewMockAPIHelper(mctrl)
	gomock.InOrder(
		mmock.EXPECT().GetBookByISBN("baz").Return(nil, errors.New("isbn error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(nil, model.ErrNotFound),
		amock.EXPECT().GetBookName("baz").Return("", errors.New("api error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(nil, model.ErrNotFound),
		amock.EXPECT().GetBookName("baz").Return("", api.ErrNoResult),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(nil, model.ErrNotFound),
		amock.EXPECT().GetBookName("baz").Return("bar", nil),
		mmock.EXPECT().InsertBook("baz", "bar", 0).Return(nil, errors.New("insert error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(nil, model.ErrNotFound),
		amock.EXPECT().GetBookName("baz").Return("bar", nil),
		mmock.EXPECT().InsertBook("baz", "bar", 0).Return(nil, model.ErrDuplicateKey),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(nil, model.ErrNotFound),
		amock.EXPECT().GetBookName("baz").Return("bar", nil),
		mmock.EXPECT().InsertBook("baz", "bar", 0).Return(nil, model.ErrInvalidID),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(nil, model.ErrNotFound),
		amock.EXPECT().GetBookName("baz").Return("bar", nil),
		mmock.EXPECT().InsertBook("baz", "bar", 0).Return(book, nil),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, nil),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(book, nil),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, errors.New("insert error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(book, nil),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(book, nil),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, model.ErrDuplicateKey),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(book, nil),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, model.ErrInvalidID),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetBookByISBN("baz").Return(book, nil),
		mmock.EXPECT().InsertOwnership(123, 456).Return(nil, nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), amock)

	logbuf.Reset()
	test.AddOwnershipsServiceUnavailable(t, ctx, service, ctrl, 123, payload)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), amock)

	logbuf.Reset()
	test.AddOwnershipsInternalServerError(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to get book by isbn error=isbn error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsInternalServerError(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to get book name error=api error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsUnprocessableEntity(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to get book name error=no volume found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsInternalServerError(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert book error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsUnprocessableEntity(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert book error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsUnprocessableEntity(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert book error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.AddOwnershipsCreated(t, ctx, service, ctrl, 123, payload)
	exp = app.OwnershipsHref(123, 456)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsInternalServerError(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert ownership error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsNotFound(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert ownership error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsUnprocessableEntity(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert ownership error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AddOwnershipsUnprocessableEntity(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to insert ownership error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw = test.AddOwnershipsCreated(t, ctx, service, ctrl, 123, payload)
	exp = app.OwnershipsHref(123, 456)
	v = rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\nexp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestOwnershipsDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	amock := NewMockAPIHelper(mctrl)
	gomock.InOrder(
		mmock.EXPECT().DeleteOwnership(123, 456).Return(errors.New("delete error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().DeleteOwnership(123, 456).Return(model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().DeleteOwnership(123, 456).Return(nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), amock)

	logbuf.Reset()
	test.DeleteOwnershipsServiceUnavailable(t, ctx, service, ctrl, 123, 456)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), amock)

	logbuf.Reset()
	test.DeleteOwnershipsInternalServerError(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] unable to delete ownership error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteOwnershipsNotFound(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] unable to delete ownership error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteOwnershipsNoContent(t, ctx, service, ctrl, 123, 456)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestOwnershipsList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	amock := NewMockAPIHelper(mctrl)
	gomock.InOrder(
		mmock.EXPECT().ListOwnershipsByUserID(123).Return(nil, errors.New("list error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().ListOwnershipsByUserID(123).Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().ListOwnershipsByUserID(123).Return([]*model.Ownership{&model.Ownership{UserID: 123, BookID: 456}, &model.Ownership{UserID: 123, BookID: 789}}, nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), amock)

	logbuf.Reset()
	test.ListOwnershipsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), amock)

	logbuf.Reset()
	test.ListOwnershipsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to get ownership list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListOwnershipsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to get ownership list error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListOwnershipsOK(t, ctx, service, ctrl, 123)
	expres := app.OwnershipCollection{
		convert.ToOwnershipMedia(&model.Ownership{UserID: 123, BookID: 456}),
		convert.ToOwnershipMedia(&model.Ownership{UserID: 123, BookID: 789}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestOwnershipsShow(t *testing.T) {
	o := &model.Ownership{UserID: 123, BookID: 456}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	amock := NewMockAPIHelper(mctrl)
	gomock.InOrder(
		mmock.EXPECT().GetOwnership(123, 456).Return(nil, errors.New("show error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetOwnership(123, 456).Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetOwnership(123, 456).Return(o, nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), amock)

	logbuf.Reset()
	test.ShowOwnershipsServiceUnavailable(t, ctx, service, ctrl, 123, 456)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewOwnershipsController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), amock)

	logbuf.Reset()
	test.ShowOwnershipsInternalServerError(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] unable to get ownership error=show error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowOwnershipsNotFound(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] unable to get ownership error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowOwnershipsOK(t, ctx, service, ctrl, 123, 456)
	expres := convert.ToOwnershipMedia(o)
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
