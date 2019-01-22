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

func TestSeriesCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertSeries("foo", 1).Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertSeries("foo", 1).Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertSeries("foo", 1).Return(nil, model.ErrInvalidID),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertSeries("foo", 1).Return(&model.Series{ID: 2, Name: "foo", CategoryID: 1}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreateSeriesServiceUnavailable(t, ctx, service, ctrl, &app.CreateSeriesPayload{SeriesName: "foo", CategoryID: 1})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreateSeriesInternalServerError(t, ctx, service, ctrl, &app.CreateSeriesPayload{SeriesName: "foo", CategoryID: 1})
	exp = "[EROR] failed to insert series error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateSeriesUnprocessableEntity(t, ctx, service, ctrl, &app.CreateSeriesPayload{SeriesName: "foo", CategoryID: 1})
	exp = "[EROR] failed to insert series error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateSeriesUnprocessableEntity(t, ctx, service, ctrl, &app.CreateSeriesPayload{SeriesName: "foo", CategoryID: 1})
	exp = "[EROR] failed to insert series error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateSeriesCreated(t, ctx, service, ctrl, &app.CreateSeriesPayload{SeriesName: "foo", CategoryID: 1})
	exp = app.SeriesHref(2)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestSeriesDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeleteSeries(123).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteSeries(123).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteSeries(123).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeleteSeriesServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeleteSeriesInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete series error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteSeriesNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete series error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteSeriesNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestSeriesList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().ListSeriesByIDs(nil, nil, nil, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().ListSeriesByIDs(nil, nil, nil, nil).Return([]*model.Series{&model.Series{ID: 1, Name: "foo", CategoryID: 2}, &model.Series{ID: 3, Name: "bar", CategoryID: 2}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListSeriesServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListSeriesOK(t, ctx, service, ctrl)
	expres := app.SeriesCollection{
		convert.ToSeriesMedia(&model.Series{ID: 1, Name: "foo", CategoryID: 2}),
		convert.ToSeriesMedia(&model.Series{ID: 3, Name: "bar", CategoryID: 2}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestSeriesShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetSeriesByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetSeriesByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetSeriesByID(123).Return(&model.Series{ID: 123, Name: "foo", CategoryID: 2}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowSeriesServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowSeriesInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get series error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowSeriesNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get series error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowSeriesOK(t, ctx, service, ctrl, 123)
	expres := convert.ToSeriesMedia(&model.Series{ID: 123, Name: "foo", CategoryID: 2})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestSeriesUpdate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	name := "foo"
	editorID := 2
	gomock.InOrder(
		mock.EXPECT().UpdateSeries(123, &name, &editorID).Return(errors.New("update error")),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateSeries(123, &name, &editorID).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateSeries(123, &name, &editorID).Return(model.ErrInvalidID),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateSeries(123, &name, &editorID).Return(model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateSeries(123, &name, &editorID).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.UpdateSeriesServiceUnavailable(t, ctx, service, ctrl, 123, &app.UpdateSeriesPayload{SeriesName: &name, CategoryID: &editorID})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewSeriesController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.UpdateSeriesInternalServerError(t, ctx, service, ctrl, 123, &app.UpdateSeriesPayload{SeriesName: &name, CategoryID: &editorID})
	exp = "[EROR] failed to update series error=update error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateSeriesNotFound(t, ctx, service, ctrl, 123, &app.UpdateSeriesPayload{SeriesName: &name, CategoryID: &editorID})
	exp = "[EROR] failed to update series error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateSeriesUnprocessableEntity(t, ctx, service, ctrl, 123, &app.UpdateSeriesPayload{SeriesName: &name, CategoryID: &editorID})
	exp = "[EROR] failed to update series error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateSeriesUnprocessableEntity(t, ctx, service, ctrl, 123, &app.UpdateSeriesPayload{SeriesName: &name, CategoryID: &editorID})
	exp = "[EROR] failed to update series error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateSeriesNoContent(t, ctx, service, ctrl, 123, &app.UpdateSeriesPayload{SeriesName: &name, CategoryID: &editorID})
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
