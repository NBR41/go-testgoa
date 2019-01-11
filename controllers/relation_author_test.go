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

func TestRelationAuthorListCategories(t *testing.T) {
	m1, m2 := &model.Category{ID: 1, Name: "foo"}, &model.Category{ID: 2, Name: "bar"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().ListCategoriesByAuthorID(123).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().ListCategoriesByAuthorID(123).Return([]*model.Category{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListCategoriesRelationAuthorServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListCategoriesRelationAuthorNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListCategoriesRelationAuthorInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListCategoriesRelationAuthorInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get category list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListCategoriesRelationAuthorOK(t, ctx, service, ctrl, 123)
	expres := app.CategoryCollection{convert.ToCategoryMedia(m1), convert.ToCategoryMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationAuthorListClasses(t *testing.T) {
	m1, m2 := &model.Class{ID: 1, Name: "foo"}, &model.Class{ID: 2, Name: "bar"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().ListClassesByAuthorID(123).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().ListClassesByAuthorID(123).Return([]*model.Class{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListClassesRelationAuthorServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListClassesRelationAuthorNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListClassesRelationAuthorInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListClassesRelationAuthorInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get class list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListClassesRelationAuthorOK(t, ctx, service, ctrl, 123)
	expres := app.ClassCollection{convert.ToClassMedia(m1), convert.ToClassMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationAuthorListRoles(t *testing.T) {
	m1, m2 := &model.Role{ID: 1, Name: "foo"}, &model.Role{ID: 2, Name: "bar"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().ListRolesByAuthorID(123).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().ListRolesByAuthorID(123).Return([]*model.Role{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListRolesRelationAuthorServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListRolesRelationAuthorNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListRolesRelationAuthorInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListRolesRelationAuthorInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get role list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListRolesRelationAuthorOK(t, ctx, service, ctrl, 123)
	expres := app.RoleCollection{convert.ToRoleMedia(m1), convert.ToRoleMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationAuthorListSeries(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	authorID := 123
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, nil, nil, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, nil, nil, nil).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesRelationAuthorServiceUnavailable(t, ctx, service, ctrl, authorID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesRelationAuthorNotFound(t, ctx, service, ctrl, authorID)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesRelationAuthorOK(t, ctx, service, ctrl, authorID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationAuthorListSeriesByCategory(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	authorID := 123
	categoryID := 456
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, nil, &categoryID, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, nil, &categoryID, nil).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesByCategoryRelationAuthorServiceUnavailable(t, ctx, service, ctrl, authorID, categoryID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesByCategoryRelationAuthorNotFound(t, ctx, service, ctrl, authorID, categoryID)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByCategoryRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, categoryID)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByCategoryRelationAuthorNotFound(t, ctx, service, ctrl, authorID, categoryID)
	exp = "[EROR] failed to get category error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByCategoryRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, categoryID)
	exp = "[EROR] failed to get category error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByCategoryRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, categoryID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesByCategoryRelationAuthorOK(t, ctx, service, ctrl, authorID, categoryID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationAuthorListSeriesByClass(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	authorID := 123
	classID := 456
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetClassByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetClassByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetClassByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, nil, nil, &classID).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetClassByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, nil, nil, &classID).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesByClassRelationAuthorServiceUnavailable(t, ctx, service, ctrl, authorID, classID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesByClassRelationAuthorNotFound(t, ctx, service, ctrl, authorID, classID)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByClassRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, classID)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByClassRelationAuthorNotFound(t, ctx, service, ctrl, authorID, classID)
	exp = "[EROR] failed to get class error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByClassRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, classID)
	exp = "[EROR] failed to get class error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByClassRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, classID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesByClassRelationAuthorOK(t, ctx, service, ctrl, authorID, classID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestRelationAuthorListSeriesByRole(t *testing.T) {
	m1, m2 := &model.Series{ID: 1, Name: "foo", CategoryID: 3}, &model.Series{ID: 2, Name: "bar", CategoryID: 3}
	mctrl := gomock.NewController(t)
	authorID := 123
	roleID := 456
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetAuthorByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetRoleByID(456).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetRoleByID(456).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetRoleByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, &roleID, nil, nil).Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetAuthorByID(123).Return(nil, nil),
		mock.EXPECT().GetRoleByID(456).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, &roleID, nil, nil).Return([]*model.Series{m1, m2}, nil),
		mock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))

	logbuf.Reset()
	test.ListSeriesByRoleRelationAuthorServiceUnavailable(t, ctx, service, ctrl, authorID, roleID)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewRelationAuthorController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListSeriesByRoleRelationAuthorNotFound(t, ctx, service, ctrl, authorID, roleID)
	exp = "[EROR] failed to get author error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByRoleRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, roleID)
	exp = "[EROR] failed to get author error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByRoleRelationAuthorNotFound(t, ctx, service, ctrl, authorID, roleID)
	exp = "[EROR] failed to get role error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByRoleRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, roleID)
	exp = "[EROR] failed to get role error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListSeriesByRoleRelationAuthorInternalServerError(t, ctx, service, ctrl, authorID, roleID)
	exp = "[EROR] failed to get series list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListSeriesByRoleRelationAuthorOK(t, ctx, service, ctrl, authorID, roleID)
	expres := app.SeriesCollection{convert.ToSeriesMedia(m1), convert.ToSeriesMedia(m2)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
