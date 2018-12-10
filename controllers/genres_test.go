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

func TestGenresCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertGenre("foo").Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertGenre("foo").Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertGenre("foo").Return(&model.Genre{ID: 123, Name: "foo"}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreateGenresServiceUnavailable(t, ctx, service, ctrl, &app.CreateGenresPayload{Name: "foo"})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreateGenresInternalServerError(t, ctx, service, ctrl, &app.CreateGenresPayload{Name: "foo"})
	exp = "[EROR] failed to insert genre error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateGenresUnprocessableEntity(t, ctx, service, ctrl, &app.CreateGenresPayload{Name: "foo"})
	exp = "[EROR] failed to insert genre error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateGenresCreated(t, ctx, service, ctrl, &app.CreateGenresPayload{Name: "foo"})
	exp = app.GenresHref(123)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestGenreDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeleteGenre(123).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteGenre(123).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteGenre(123).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeleteGenresServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeleteGenresInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete genre error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteGenresNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete genre error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteGenresNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestGenreList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetGenreList().Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetGenreList().Return([]*model.Genre{&model.Genre{ID: 123, Name: "foo"}, &model.Genre{ID: 456, Name: "bar"}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListGenresServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListGenresInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get genre list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListGenresOK(t, ctx, service, ctrl)
	expres := app.GenreCollection{
		convert.ToGenreMedia(&model.Genre{ID: 123, Name: "foo"}),
		convert.ToGenreMedia(&model.Genre{ID: 456, Name: "bar"}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestGenreShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetGenreByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetGenreByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetGenreByID(123).Return(&model.Genre{ID: 123, Name: "foo"}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowGenresServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowGenresInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get genre error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowGenresNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get genre error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowGenresOK(t, ctx, service, ctrl, 123)
	expres := convert.ToGenreMedia(&model.Genre{ID: 123, Name: "foo"})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestGenreUpdate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().UpdateGenre(123, "foo").Return(errors.New("update error")),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateGenre(123, "foo").Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateGenre(123, "foo").Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.UpdateGenresServiceUnavailable(t, ctx, service, ctrl, 123, &app.UpdateGenresPayload{Name: "foo"})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewGenresController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.UpdateGenresInternalServerError(t, ctx, service, ctrl, 123, &app.UpdateGenresPayload{Name: "foo"})
	exp = "[EROR] failed to update genre error=update error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateGenresNotFound(t, ctx, service, ctrl, 123, &app.UpdateGenresPayload{Name: "foo"})
	exp = "[EROR] failed to update genre error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateGenresNoContent(t, ctx, service, ctrl, 123, &app.UpdateGenresPayload{Name: "foo"})
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
