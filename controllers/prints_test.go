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

func TestPrintsCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertPrint("foo").Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertPrint("foo").Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertPrint("foo").Return(&model.Print{ID: 123, Name: "foo"}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreatePrintsServiceUnavailable(t, ctx, service, ctrl, &app.CreatePrintsPayload{PrintName: "foo"})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreatePrintsInternalServerError(t, ctx, service, ctrl, &app.CreatePrintsPayload{PrintName: "foo"})
	exp = "[EROR] failed to insert print error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreatePrintsUnprocessableEntity(t, ctx, service, ctrl, &app.CreatePrintsPayload{PrintName: "foo"})
	exp = "[EROR] failed to insert print error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreatePrintsCreated(t, ctx, service, ctrl, &app.CreatePrintsPayload{PrintName: "foo"})
	exp = app.PrintsHref(123)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestPrintDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeletePrint(123).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeletePrint(123).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeletePrint(123).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeletePrintsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeletePrintsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete print error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeletePrintsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete print error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeletePrintsNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestPrintList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().ListPrints().Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().ListPrints().Return([]*model.Print{&model.Print{ID: 123, Name: "foo"}, &model.Print{ID: 456, Name: "bar"}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListPrintsServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListPrintsInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get print list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListPrintsOK(t, ctx, service, ctrl)
	expres := app.PrintCollection{
		convert.ToPrintMedia(&model.Print{ID: 123, Name: "foo"}),
		convert.ToPrintMedia(&model.Print{ID: 456, Name: "bar"}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestPrintShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetPrintByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetPrintByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetPrintByID(123).Return(&model.Print{ID: 123, Name: "foo"}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowPrintsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowPrintsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get print error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowPrintsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get print error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowPrintsOK(t, ctx, service, ctrl, 123)
	expres := convert.ToPrintMedia(&model.Print{ID: 123, Name: "foo"})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestPrintUpdate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().UpdatePrint(123, "foo").Return(errors.New("update error")),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdatePrint(123, "foo").Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdatePrint(123, "foo").Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.UpdatePrintsServiceUnavailable(t, ctx, service, ctrl, 123, &app.UpdatePrintsPayload{PrintName: "foo"})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewPrintsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.UpdatePrintsInternalServerError(t, ctx, service, ctrl, 123, &app.UpdatePrintsPayload{PrintName: "foo"})
	exp = "[EROR] failed to update print error=update error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdatePrintsNotFound(t, ctx, service, ctrl, 123, &app.UpdatePrintsPayload{PrintName: "foo"})
	exp = "[EROR] failed to update print error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdatePrintsNoContent(t, ctx, service, ctrl, 123, &app.UpdatePrintsPayload{PrintName: "foo"})
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
