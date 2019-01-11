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

func TestRelationClassListCategories(t *testing.T) {
	m1, m2 := &model.Category{ID: 1, Name: "foo"}, &model.Category{ID: 2, Name: "bar"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetClassByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetClassByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetClassByID(123).Return(nil, nil),
		mock.EXPECT().ListCategoriesByClassID(123).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetClassByID(123).Return(nil, nil),
		mock.EXPECT().ListCategoriesByClassID(123).Return([]*model.Category{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationClassController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListCategoriesRelationClassServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationClassController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListCategoriesRelationClassNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get class error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListCategoriesRelationClassInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get class error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListCategoriesRelationClassInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get category list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListCategoriesRelationClassOK(t, ctx, service, ctrl, 123)
	expres := app.CategoryCollection{convert.ToCategoryMedia(m1), convert.ToCategoryMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationClassListSeries(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	classID := 456
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetClassByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetClassByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetClassByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(nil, nil, nil, &classID).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetClassByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(nil, nil, nil, &classID).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationClassController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesRelationClassServiceUnavailable(t, ctx, service, ctrl, classID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationClassController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesRelationClassNotFound(t, ctx, service, ctrl, classID)
	exp = "[EROR] failed to get class error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesRelationClassInternalServerError(t, ctx, service, ctrl, classID)
	exp = "[EROR] failed to get class error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesRelationClassInternalServerError(t, ctx, service, ctrl, classID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesRelationClassOK(t, ctx, service, ctrl, classID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationClassListSeriesByCategory(t *testing.T) {
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

	ctrl := NewRelationClassController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesByCategoryRelationClassServiceUnavailable(t, ctx, service, ctrl, classID, categoryID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationClassController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesByCategoryRelationClassNotFound(t, ctx, service, ctrl, classID, categoryID)
	exp = "[EROR] failed to get category error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByCategoryRelationClassInternalServerError(t, ctx, service, ctrl, classID, categoryID)
	exp = "[EROR] failed to get category error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByCategoryRelationClassNotFound(t, ctx, service, ctrl, classID, categoryID)
	exp = "[EROR] failed to get class error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByCategoryRelationClassInternalServerError(t, ctx, service, ctrl, classID, categoryID)
	exp = "[EROR] failed to get class error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByCategoryRelationClassInternalServerError(t, ctx, service, ctrl, classID, categoryID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesByCategoryRelationClassOK(t, ctx, service, ctrl, classID, categoryID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
