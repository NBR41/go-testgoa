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

func TestAuthorsCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertAuthor("foo").Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertAuthor("foo").Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertAuthor("foo").Return(&model.Author{ID: 123, Name: "foo"}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreateAuthorsServiceUnavailable(t, ctx, service, ctrl, &app.CreateAuthorsPayload{Name: "foo"})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreateAuthorsInternalServerError(t, ctx, service, ctrl, &app.CreateAuthorsPayload{Name: "foo"})
	exp = "[EROR] failed to insert author error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateAuthorsUnprocessableEntity(t, ctx, service, ctrl, &app.CreateAuthorsPayload{Name: "foo"})
	exp = "[EROR] failed to insert author error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateAuthorsCreated(t, ctx, service, ctrl, &app.CreateAuthorsPayload{Name: "foo"})
	exp = app.AuthorsHref(123)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestAuthorDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeleteAuthor(123).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteAuthor(123).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteAuthor(123).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeleteAuthorsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeleteAuthorsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete author error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteAuthorsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteAuthorsNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestAuthorList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorList().Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorList().Return([]*model.Author{&model.Author{ID: 123, Name: "foo"}, &model.Author{ID: 456, Name: "bar"}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListAuthorsServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListAuthorsInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get author list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListAuthorsOK(t, ctx, service, ctrl)
	expres := app.AuthorCollection{
		convert.ToAuthorMedia(&model.Author{ID: 123, Name: "foo"}),
		convert.ToAuthorMedia(&model.Author{ID: 456, Name: "bar"}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestAuthorShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(&model.Author{ID: 123, Name: "foo"}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowAuthorsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowAuthorsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowAuthorsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowAuthorsOK(t, ctx, service, ctrl, 123)
	expres := convert.ToAuthorMedia(&model.Author{ID: 123, Name: "foo"})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestAuthorUpdate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().UpdateAuthor(123, "foo").Return(errors.New("update error")),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateAuthor(123, "foo").Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateAuthor(123, "foo").Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.UpdateAuthorsServiceUnavailable(t, ctx, service, ctrl, 123, &app.UpdateAuthorsPayload{Name: "foo"})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthorsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.UpdateAuthorsInternalServerError(t, ctx, service, ctrl, 123, &app.UpdateAuthorsPayload{Name: "foo"})
	exp = "[EROR] failed to update author error=update error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateAuthorsNotFound(t, ctx, service, ctrl, 123, &app.UpdateAuthorsPayload{Name: "foo"})
	exp = "[EROR] failed to update author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateAuthorsNoContent(t, ctx, service, ctrl, 123, &app.UpdateAuthorsPayload{Name: "foo"})
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
