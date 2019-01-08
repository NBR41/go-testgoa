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

func TestEditionsCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertEdition(1, 2, 3).Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertEdition(1, 2, 3).Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertEdition(1, 2, 3).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertEdition(1, 2, 3).Return(&model.Edition{ID: 4, BookID: 1, CollectionID: 2, PrintID: 3}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewEditionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreateEditionsServiceUnavailable(t, ctx, service, ctrl, &app.CreateEditionsPayload{BookID: 1, CollectionID: 2, PrintID: 3})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewEditionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreateEditionsInternalServerError(t, ctx, service, ctrl, &app.CreateEditionsPayload{BookID: 1, CollectionID: 2, PrintID: 3})
	exp = "[EROR] failed to insert edition error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateEditionsUnprocessableEntity(t, ctx, service, ctrl, &app.CreateEditionsPayload{BookID: 1, CollectionID: 2, PrintID: 3})
	exp = "[EROR] failed to insert edition error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateEditionsUnprocessableEntity(t, ctx, service, ctrl, &app.CreateEditionsPayload{BookID: 1, CollectionID: 2, PrintID: 3})
	exp = "[EROR] failed to insert edition error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateEditionsCreated(t, ctx, service, ctrl, &app.CreateEditionsPayload{BookID: 1, CollectionID: 2, PrintID: 3})
	exp = app.EditionsHref(4)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestEditionsDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeleteEdition(123).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteEdition(123).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteEdition(123).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewEditionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeleteEditionsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewEditionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeleteEditionsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete edition error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteEditionsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete edition error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteEditionsNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestEditionsList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().ListEditions().Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().ListEditions().Return([]*model.Edition{&model.Edition{ID: 1, BookID: 2, CollectionID: 3, PrintID: 4}, &model.Edition{ID: 5, BookID: 6, CollectionID: 7, PrintID: 8}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewEditionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListEditionsServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewEditionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListEditionsInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get edition list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListEditionsOK(t, ctx, service, ctrl)
	expres := app.EditionCollection{
		convert.ToEditionMedia(&model.Edition{ID: 1, BookID: 2, CollectionID: 3, PrintID: 4}),
		convert.ToEditionMedia(&model.Edition{ID: 5, BookID: 6, CollectionID: 7, PrintID: 8}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestEditionsShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetEditionByID(1).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetEditionByID(1).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetEditionByID(1).Return(&model.Edition{ID: 1, BookID: 2, CollectionID: 3, PrintID: 4}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewEditionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowEditionsServiceUnavailable(t, ctx, service, ctrl, 1)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewEditionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowEditionsInternalServerError(t, ctx, service, ctrl, 1)
	exp = "[EROR] failed to get edition error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowEditionsNotFound(t, ctx, service, ctrl, 1)
	exp = "[EROR] failed to get edition error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowEditionsOK(t, ctx, service, ctrl, 1)
	expres := convert.ToEditionMedia(&model.Edition{ID: 1, BookID: 2, CollectionID: 3, PrintID: 4})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
