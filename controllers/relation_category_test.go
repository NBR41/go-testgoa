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

func TestRelationCategoryListAuthors(t *testing.T) {
	m1, m2 := &model.Author{ID: 1, Name: "foo"}, &model.Author{ID: 2, Name: "bar"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCategoryByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().ListAuthorsByCategoryID(123).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().ListAuthorsByCategoryID(123).Return([]*model.Author{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCategoryController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListAuthorsRelationCategoryServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCategoryController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListAuthorsRelationCategoryNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get category error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListAuthorsRelationCategoryInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get category error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListAuthorsRelationCategoryInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListAuthorsRelationCategoryOK(t, ctx, service, ctrl, 123)
	expres := app.AuthorCollection{convert.ToAuthorMedia(m1), convert.ToAuthorMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCategoryListClasses(t *testing.T) {
	m1, m2 := &model.Class{ID: 1, Name: "foo"}, &model.Class{ID: 2, Name: "bar"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCategoryByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().ListClassesByCategoryID(123).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().ListClassesByCategoryID(123).Return([]*model.Class{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCategoryController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListClassesRelationCategoryServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCategoryController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListClassesRelationCategoryNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get category error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListClassesRelationCategoryInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get category error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListClassesRelationCategoryInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get class list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListClassesRelationCategoryOK(t, ctx, service, ctrl, 123)
	expres := app.ClassCollection{convert.ToClassMedia(m1), convert.ToClassMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCategoryListSeries(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	categoryID := 123
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCategoryByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(nil, nil, &categoryID, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(nil, nil, &categoryID, nil).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCategoryController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesRelationCategoryServiceUnavailable(t, ctx, service, ctrl, categoryID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCategoryController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesRelationCategoryNotFound(t, ctx, service, ctrl, categoryID)
	exp = "[EROR] failed to get category error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesRelationCategoryInternalServerError(t, ctx, service, ctrl, categoryID)
	exp = "[EROR] failed to get category error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesRelationCategoryInternalServerError(t, ctx, service, ctrl, categoryID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesRelationCategoryOK(t, ctx, service, ctrl, categoryID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationCategoryListSeriesByClass(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	categoryID := 123
	classID := 456
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCategoryByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().GetClassByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().GetClassByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().GetClassByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(nil, nil, &categoryID, &classID).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCategoryByID(123).Return(nil, nil),
		mock.EXPECT().GetClassByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(nil, nil, &categoryID, &classID).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationCategoryController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesByClassRelationCategoryServiceUnavailable(t, ctx, service, ctrl, categoryID, classID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationCategoryController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesByClassRelationCategoryNotFound(t, ctx, service, ctrl, categoryID, classID)
	exp = "[EROR] failed to get category error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByClassRelationCategoryInternalServerError(t, ctx, service, ctrl, categoryID, classID)
	exp = "[EROR] failed to get category error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByClassRelationCategoryNotFound(t, ctx, service, ctrl, categoryID, classID)
	exp = "[EROR] failed to get class error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByClassRelationCategoryInternalServerError(t, ctx, service, ctrl, categoryID, classID)
	exp = "[EROR] failed to get class error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByClassRelationCategoryInternalServerError(t, ctx, service, ctrl, categoryID, classID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesByClassRelationCategoryOK(t, ctx, service, ctrl, categoryID, classID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
