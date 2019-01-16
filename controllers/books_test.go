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

func TestBooksCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertBook("foo", "bar", 1).Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertBook("foo", "bar", 1).Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertBook("foo", "bar", 1).Return(nil, model.ErrInvalidID),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertBook("foo", "bar", 1).Return(&model.Book{ID: 2, ISBN: "foo", Name: "bar", SeriesID: 1}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreateBooksServiceUnavailable(t, ctx, service, ctrl, &app.CreateBooksPayload{BookIsbn: "foo", BookName: "bar", SeriesID: 1})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreateBooksInternalServerError(t, ctx, service, ctrl, &app.CreateBooksPayload{BookIsbn: "foo", BookName: "bar", SeriesID: 1})
	exp = "[EROR] failed to insert book error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateBooksUnprocessableEntity(t, ctx, service, ctrl, &app.CreateBooksPayload{BookIsbn: "foo", BookName: "bar", SeriesID: 1})
	exp = "[EROR] failed to insert book error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateBooksUnprocessableEntity(t, ctx, service, ctrl, &app.CreateBooksPayload{BookIsbn: "foo", BookName: "bar", SeriesID: 1})
	exp = "[EROR] failed to insert book error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateBooksCreated(t, ctx, service, ctrl, &app.CreateBooksPayload{BookIsbn: "foo", BookName: "bar", SeriesID: 1})
	exp = app.BooksHref(2)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestBooksDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeleteBook(123).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteBook(123).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteBook(123).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeleteBooksServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeleteBooksInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete book error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteBooksNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete book error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteBooksNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestBooksList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().ListBooks().Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().ListBooks().Return([]*model.Book{&model.Book{ID: 1, ISBN: "foo1", Name: "bar1", SeriesID: 2}, &model.Book{ID: 3, ISBN: "foo2", Name: "bar2", SeriesID: 4}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListBooksServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListBooksInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get book list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListBooksOK(t, ctx, service, ctrl)
	expres := app.BookCollection{
		convert.ToBookMedia(&model.Book{ID: 1, ISBN: "foo1", Name: "bar1", SeriesID: 2}),
		convert.ToBookMedia(&model.Book{ID: 3, ISBN: "foo2", Name: "bar2", SeriesID: 4}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestBooksShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetBookByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetBookByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetBookByID(123).Return(&model.Book{ID: 1, ISBN: "foo1", Name: "bar1", SeriesID: 2}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowBooksServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowBooksInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get book error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowBooksNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get book error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowBooksOK(t, ctx, service, ctrl, 123)
	expres := convert.ToBookMedia(&model.Book{ID: 1, ISBN: "foo1", Name: "bar1", SeriesID: 2})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestBooksUpdate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	bookName := "foo"
	seriesID := 2
	gomock.InOrder(
		mock.EXPECT().UpdateBook(123, &bookName, &seriesID).Return(errors.New("update error")),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateBook(123, &bookName, &seriesID).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateBook(123, &bookName, &seriesID).Return(model.ErrInvalidID),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateBook(123, &bookName, &seriesID).Return(model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateBook(123, &bookName, &seriesID).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.UpdateBooksServiceUnavailable(t, ctx, service, ctrl, 123, &app.UpdateBooksPayload{BookName: &bookName, SeriesID: &seriesID})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewBooksController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.UpdateBooksInternalServerError(t, ctx, service, ctrl, 123, &app.UpdateBooksPayload{BookName: &bookName, SeriesID: &seriesID})
	exp = "[EROR] failed to update book error=update error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateBooksNotFound(t, ctx, service, ctrl, 123, &app.UpdateBooksPayload{BookName: &bookName, SeriesID: &seriesID})
	exp = "[EROR] failed to update book error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateBooksUnprocessableEntity(t, ctx, service, ctrl, 123, &app.UpdateBooksPayload{BookName: &bookName, SeriesID: &seriesID})
	exp = "[EROR] failed to update book error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateBooksUnprocessableEntity(t, ctx, service, ctrl, 123, &app.UpdateBooksPayload{BookName: &bookName, SeriesID: &seriesID})
	exp = "[EROR] failed to update book error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateBooksNoContent(t, ctx, service, ctrl, 123, &app.UpdateBooksPayload{BookName: &bookName, SeriesID: &seriesID})
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
