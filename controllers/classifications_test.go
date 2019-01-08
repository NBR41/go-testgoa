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

func TestClassificationsCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertClassification(1, 2).Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertClassification(1, 2).Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertClassification(1, 2).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertClassification(1, 2).Return(&model.Class{ID: 2, Name: "foo"}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewClassificationsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreateClassificationsServiceUnavailable(t, ctx, service, ctrl, 1, &app.CreateClassificationsPayload{ClassID: 2})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewClassificationsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreateClassificationsInternalServerError(t, ctx, service, ctrl, 1, &app.CreateClassificationsPayload{ClassID: 2})
	exp = "[EROR] failed to insert classification error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateClassificationsUnprocessableEntity(t, ctx, service, ctrl, 1, &app.CreateClassificationsPayload{ClassID: 2})
	exp = "[EROR] failed to insert classification error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateClassificationsUnprocessableEntity(t, ctx, service, ctrl, 1, &app.CreateClassificationsPayload{ClassID: 2})
	exp = "[EROR] failed to insert classification error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateClassificationsCreated(t, ctx, service, ctrl, 1, &app.CreateClassificationsPayload{ClassID: 2})
	exp = app.ClassificationsHref(1, 2)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestClassificationsDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeleteClassification(123, 456).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteClassification(123, 456).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteClassification(123, 456).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewClassificationsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeleteClassificationsServiceUnavailable(t, ctx, service, ctrl, 123, 456)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewClassificationsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeleteClassificationsInternalServerError(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to delete classification error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteClassificationsNotFound(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to delete classification error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteClassificationsNoContent(t, ctx, service, ctrl, 123, 456)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestClassificationsList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().ListClassesBySeriesID(2).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().ListClassesBySeriesID(2).Return([]*model.Class{&model.Class{ID: 1, Name: "foo"}, &model.Class{ID: 2, Name: "bar"}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewClassificationsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListClassificationsServiceUnavailable(t, ctx, service, ctrl, 2)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewClassificationsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListClassificationsInternalServerError(t, ctx, service, ctrl, 2)
	exp = "[EROR] failed to get classification list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListClassificationsOK(t, ctx, service, ctrl, 2)
	expres := app.ClassificationCollection{
		convert.ToClassificationMedia(2, &model.Class{ID: 1, Name: "foo"}),
		convert.ToClassificationMedia(2, &model.Class{ID: 2, Name: "bar"}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestClassificationsShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetClassification(123, 456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetClassification(123, 456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetClassification(123, 456).Return(&model.Class{ID: 1, Name: "foo"}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewClassificationsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowClassificationsServiceUnavailable(t, ctx, service, ctrl, 123, 456)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewClassificationsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowClassificationsInternalServerError(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to get classification error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowClassificationsNotFound(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to get classification error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowClassificationsOK(t, ctx, service, ctrl, 123, 456)
	expres := convert.ToClassificationMedia(123, &model.Class{ID: 1, Name: "foo"})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
