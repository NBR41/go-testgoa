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

func TestRelationCollecionListBooks(t *testing.T) {
	m1, m2 := &model.Book{ID: 1, Name: "foo", ISBN: "qux"}, &model.Book{ID: 2, Name: "bar", ISBN: "quux"}
	collectionID := 123
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, nil, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, nil, nil).Return([]*model.Book{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListBooksRelationCollectionServiceUnavailable(t, ctx, service, ctrl, collectionID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListBooksRelationCollectionNotFound(t, ctx, service, ctrl, collectionID)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID)
	exp = "[EROR] failed to get book list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListBooksRelationCollectionOK(t, ctx, service, ctrl, collectionID)
	expres := app.BookCollection{convert.ToBookMedia(m1), convert.ToBookMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCollecionListBooksByPrint(t *testing.T) {
	m1, m2 := &model.Book{ID: 1, Name: "foo", ISBN: "qux"}, &model.Book{ID: 2, Name: "bar", ISBN: "quux"}
	collectionID := 123
	printID := 456
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, &printID, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, &printID, nil).Return([]*model.Book{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListBooksByPrintRelationCollectionServiceUnavailable(t, ctx, service, ctrl, collectionID, printID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListBooksByPrintRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, printID)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, printID)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, printID)
	exp = "[EROR] failed to get print error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, printID)
	exp = "[EROR] failed to get print error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, printID)
	exp = "[EROR] failed to get book list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListBooksByPrintRelationCollectionOK(t, ctx, service, ctrl, collectionID, printID)
	expres := app.BookCollection{convert.ToBookMedia(m1), convert.ToBookMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCollecionListBooksByPrintSeries(t *testing.T) {
	m1, m2 := &model.Book{ID: 1, Name: "foo", ISBN: "qux"}, &model.Book{ID: 2, Name: "bar", ISBN: "quux"}
	collectionID := 123
	printID := 456
	seriesID := 789
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(789).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(789).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(789).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, &printID, &seriesID).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(789).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, &printID, &seriesID).Return([]*model.Book{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListBooksByPrintSeriesRelationCollectionServiceUnavailable(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListBooksByPrintSeriesRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintSeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintSeriesRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp = "[EROR] failed to get print error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintSeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp = "[EROR] failed to get print error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintSeriesRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp = "[EROR] failed to get series error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintSeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp = "[EROR] failed to get series error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksByPrintSeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp = "[EROR] failed to get book list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListBooksByPrintSeriesRelationCollectionOK(t, ctx, service, ctrl, collectionID, printID, seriesID)
	expres := app.BookCollection{convert.ToBookMedia(m1), convert.ToBookMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCollecionListBooksBySeries(t *testing.T) {
	m1, m2 := &model.Book{ID: 1, Name: "foo", ISBN: "qux"}, &model.Book{ID: 2, Name: "bar", ISBN: "quux"}
	collectionID := 123
	seriesID := 456
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(456).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, nil, &seriesID).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(456).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, nil, &seriesID).Return([]*model.Book{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListBooksBySeriesRelationCollectionServiceUnavailable(t, ctx, service, ctrl, collectionID, seriesID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListBooksBySeriesRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get series error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get series error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get book list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListBooksBySeriesRelationCollectionOK(t, ctx, service, ctrl, collectionID, seriesID)
	expres := app.BookCollection{convert.ToBookMedia(m1), convert.ToBookMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCollecionListBooksBySeriesPrint(t *testing.T) {
	m1, m2 := &model.Book{ID: 1, Name: "foo", ISBN: "qux"}, &model.Book{ID: 2, Name: "bar", ISBN: "quux"}
	collectionID := 123
	printID := 456
	seriesID := 789
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(789).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(789).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(789).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, &printID, &seriesID).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(789).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, &printID, &seriesID).Return([]*model.Book{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListBooksBySeriesPrintRelationCollectionServiceUnavailable(t, ctx, service, ctrl, collectionID, printID, seriesID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListBooksBySeriesPrintRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, seriesID, printID)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID, printID)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesPrintRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, seriesID, printID)
	exp = "[EROR] failed to get print error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID, printID)
	exp = "[EROR] failed to get print error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesPrintRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, seriesID, printID)
	exp = "[EROR] failed to get series error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID, printID)
	exp = "[EROR] failed to get series error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListBooksBySeriesPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID, printID)
	exp = "[EROR] failed to get book list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListBooksBySeriesPrintRelationCollectionOK(t, ctx, service, ctrl, collectionID, seriesID, printID)
	expres := app.BookCollection{convert.ToBookMedia(m1), convert.ToBookMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCollecionListPrints(t *testing.T) {
	m1, m2 := &model.Print{ID: 1, Name: "foo"}, &model.Print{ID: 2, Name: "bar"}
	collectionID := 123
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().ListPrintsByIDs(&collectionID, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().ListPrintsByIDs(&collectionID, nil).Return([]*model.Print{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListPrintsRelationCollectionServiceUnavailable(t, ctx, service, ctrl, collectionID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListPrintsRelationCollectionNotFound(t, ctx, service, ctrl, collectionID)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListPrintsRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListPrintsRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID)
	exp = "[EROR] failed to get print list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListPrintsRelationCollectionOK(t, ctx, service, ctrl, collectionID)
	expres := app.PrintCollection{convert.ToPrintMedia(m1), convert.ToPrintMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCollecionListPrintsBySeries(t *testing.T) {
	m1, m2 := &model.Print{ID: 1, Name: "foo"}, &model.Print{ID: 2, Name: "bar"}
	collectionID := 123
	seriesID := 456
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(456).Return(nil, nil),
		mock.EXPECT().ListPrintsByIDs(&collectionID, &seriesID).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(456).Return(nil, nil),
		mock.EXPECT().ListPrintsByIDs(&collectionID, &seriesID).Return([]*model.Print{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListPrintsBySeriesRelationCollectionServiceUnavailable(t, ctx, service, ctrl, collectionID, seriesID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListPrintsBySeriesRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListPrintsBySeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListPrintsBySeriesRelationCollectionNotFound(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get series error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListPrintsBySeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get series error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListPrintsBySeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, collectionID, seriesID)
	exp = "[EROR] failed to get print list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListPrintsBySeriesRelationCollectionOK(t, ctx, service, ctrl, collectionID, seriesID)
	expres := app.PrintCollection{convert.ToPrintMedia(m1), convert.ToPrintMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCollecionListSeries(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().ListSeriesByCollectionID(123).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().ListSeriesByCollectionID(123).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesRelationCollectionServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesRelationCollectionNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesRelationCollectionInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesRelationCollectionOK(t, ctx, service, ctrl, 123)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCollecionListSeriesByPrint(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByCollectionIDPrintID(123, 456).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, nil),
		mock.EXPECT().GetPrintByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByCollectionIDPrintID(123, 456).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesByPrintRelationCollectionServiceUnavailable(t, ctx, service, ctrl, 123, 456)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCollectionController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesByPrintRelationCollectionNotFound(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByPrintRelationCollectionNotFound(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to get print error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to get print error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByPrintRelationCollectionInternalServerError(t, ctx, service, ctrl, 123, 456)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesByPrintRelationCollectionOK(t, ctx, service, ctrl, 123, 456)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
