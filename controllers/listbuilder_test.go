package controllers

import (
	"context"
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestBuilderListAuthors(t *testing.T) {
	var categoryID, roleID int = 1, 2
	lret := []*model.Author{&model.Author{ID: 1, Name: "foo"}, &model.Author{ID: 2, Name: "bar"}}
	lparam := app.AuthorCollection{convert.ToAuthorMedia(lret[0]), convert.ToAuthorMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockauthorsResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetCategoryByID(1).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCategoryByID(1).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCategoryByID(1).Return(nil, nil),
		mock.EXPECT().GetRoleByID(2).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCategoryByID(1).Return(nil, nil),
		mock.EXPECT().GetRoleByID(2).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCategoryByID(1).Return(nil, nil),
		mock.EXPECT().GetRoleByID(2).Return(nil, nil),
		mock.EXPECT().ListAuthorsByIDs(&categoryID, &roleID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCategoryByID(1).Return(nil, nil),
		mock.EXPECT().GetRoleByID(2).Return(nil, nil),
		mock.EXPECT().ListAuthorsByIDs(&categoryID, &roleID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"category error", rfm, "[EROR] failed to get category error=get error\n"},
		{"category not found", rfm, "[EROR] failed to get category error=not found\n"},
		{"role error", rfm, "[EROR] failed to get role error=get error\n"},
		{"role not found", rfm, "[EROR] failed to get role error=not found\n"},
		{"list error", rfm, "[EROR] failed to get author list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListAuthors(ctx, tests[i].fm, resp, &categoryID, &roleID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}

func TestBuilderListBooks(t *testing.T) {
	var collectionID, editorID, printID, seriesID int = 1, 2, 3, 4
	lret := []*model.Book{&model.Book{ID: 1, Name: "foo"}, &model.Book{ID: 2, Name: "bar"}}
	lparam := app.BookCollection{convert.ToBookMedia(lret[0]), convert.ToBookMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockbooksResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, &editorID, &printID, &seriesID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, nil),
		mock.EXPECT().ListBooksByIDs(&collectionID, &editorID, &printID, &seriesID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"collection error", rfm, "[EROR] failed to get collection error=get error\n"},
		{"collection not found", rfm, "[EROR] failed to get collection error=not found\n"},
		{"editor error", rfm, "[EROR] failed to get editor error=get error\n"},
		{"editor not found", rfm, "[EROR] failed to get editor error=not found\n"},
		{"print error", rfm, "[EROR] failed to get print error=get error\n"},
		{"print not found", rfm, "[EROR] failed to get print error=not found\n"},
		{"series error", rfm, "[EROR] failed to get series error=get error\n"},
		{"series not found", rfm, "[EROR] failed to get series error=not found\n"},
		{"list error", rfm, "[EROR] failed to get book list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListBooks(ctx, tests[i].fm, resp, &collectionID, &editorID, &printID, &seriesID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}

func TestBuilderListCategories(t *testing.T) {
	var authorID, classID int = 1, 3
	lret := []*model.Category{&model.Category{ID: 1, Name: "foo"}, &model.Category{ID: 2, Name: "bar"}}
	lparam := app.CategoryCollection{convert.ToCategoryMedia(lret[0]), convert.ToCategoryMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockcategoriesResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, nil),
		mock.EXPECT().ListCategoriesByIDs(&authorID, &classID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, nil),
		mock.EXPECT().ListCategoriesByIDs(&authorID, &classID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"author error", rfm, "[EROR] failed to get author error=get error\n"},
		{"author not found", rfm, "[EROR] failed to get author error=not found\n"},
		{"class error", rfm, "[EROR] failed to get class error=get error\n"},
		{"class not found", rfm, "[EROR] failed to get class error=not found\n"},
		{"list error", rfm, "[EROR] failed to get category list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListCategories(ctx, tests[i].fm, resp, &authorID, &classID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}

func TestBuilderListClasses(t *testing.T) {
	var authorID, categoryID, seriesID int = 1, 2, 3
	lret := []*model.Class{&model.Class{ID: 1, Name: "foo"}, &model.Class{ID: 2, Name: "bar"}}
	lparam := app.ClassCollection{convert.ToClassMedia(lret[0]), convert.ToClassMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockclassesResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(3).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(3).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(3).Return(nil, nil),
		mock.EXPECT().ListClassesByIDs(&authorID, &categoryID, &seriesID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(3).Return(nil, nil),
		mock.EXPECT().ListClassesByIDs(&authorID, &categoryID, &seriesID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"author error", rfm, "[EROR] failed to get author error=get error\n"},
		{"author not found", rfm, "[EROR] failed to get author error=not found\n"},
		{"category error", rfm, "[EROR] failed to get category error=get error\n"},
		{"category not found", rfm, "[EROR] failed to get category error=not found\n"},
		{"series error", rfm, "[EROR] failed to get series error=get error\n"},
		{"series not found", rfm, "[EROR] failed to get series error=not found\n"},
		{"list error", rfm, "[EROR] failed to get class list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListClasses(ctx, tests[i].fm, resp, &authorID, &categoryID, &seriesID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}

func TestBuilderListCollections(t *testing.T) {
	var editorID, printID, seriesID int = 2, 3, 4
	lret := []*model.Collection{&model.Collection{ID: 1, Name: "foo"}, &model.Collection{ID: 2, Name: "bar"}}
	lparam := app.CollectionCollection{convert.ToCollectionMedia(lret[0]), convert.ToCollectionMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockcollectionsResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetEditorByID(2).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetEditorByID(2).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, nil),
		mock.EXPECT().ListCollectionsByIDs(&editorID, &printID, &seriesID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, nil),
		mock.EXPECT().ListCollectionsByIDs(&editorID, &printID, &seriesID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"editor error", rfm, "[EROR] failed to get editor error=get error\n"},
		{"editor not found", rfm, "[EROR] failed to get editor error=not found\n"},
		{"print error", rfm, "[EROR] failed to get print error=get error\n"},
		{"print not found", rfm, "[EROR] failed to get print error=not found\n"},
		{"series error", rfm, "[EROR] failed to get series error=get error\n"},
		{"series not found", rfm, "[EROR] failed to get series error=not found\n"},
		{"list error", rfm, "[EROR] failed to get collection list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListCollections(ctx, tests[i].fm, resp, &editorID, &printID, &seriesID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}

func TestBuilderListEditors(t *testing.T) {
	var printID, seriesID int = 3, 4
	lret := []*model.Editor{&model.Editor{ID: 1, Name: "foo"}, &model.Editor{ID: 2, Name: "bar"}}
	lparam := app.EditorCollection{convert.ToEditorMedia(lret[0]), convert.ToEditorMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockeditorsResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetPrintByID(3).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetPrintByID(3).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, nil),
		mock.EXPECT().ListEditorsByIDs(&printID, &seriesID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, nil),
		mock.EXPECT().ListEditorsByIDs(&printID, &seriesID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"print error", rfm, "[EROR] failed to get print error=get error\n"},
		{"print not found", rfm, "[EROR] failed to get print error=not found\n"},
		{"series error", rfm, "[EROR] failed to get series error=get error\n"},
		{"series not found", rfm, "[EROR] failed to get series error=not found\n"},
		{"list error", rfm, "[EROR] failed to get editor list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListEditors(ctx, tests[i].fm, resp, &printID, &seriesID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}

func TestBuilderListPrints(t *testing.T) {
	var collectionID, editorID, seriesID int = 1, 2, 4
	lret := []*model.Print{&model.Print{ID: 1, Name: "foo"}, &model.Print{ID: 2, Name: "bar"}}
	lparam := app.PrintCollection{convert.ToPrintMedia(lret[0]), convert.ToPrintMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockprintsResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, nil),
		mock.EXPECT().ListPrintsByIDs(&collectionID, &editorID, &seriesID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetSeriesByID(4).Return(nil, nil),
		mock.EXPECT().ListPrintsByIDs(&collectionID, &editorID, &seriesID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"collection error", rfm, "[EROR] failed to get collection error=get error\n"},
		{"collection not found", rfm, "[EROR] failed to get collection error=not found\n"},
		{"editor error", rfm, "[EROR] failed to get editor error=get error\n"},
		{"editor not found", rfm, "[EROR] failed to get editor error=not found\n"},
		{"series error", rfm, "[EROR] failed to get series error=get error\n"},
		{"series not found", rfm, "[EROR] failed to get series error=not found\n"},
		{"list error", rfm, "[EROR] failed to get print list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListPrints(ctx, tests[i].fm, resp, &collectionID, &editorID, &seriesID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}

func TestBuilderListSeries(t *testing.T) {
	var authorID, categoryID, classID, roleID int = 1, 2, 3, 4
	lret := []*model.Series{&model.Series{ID: 1, Name: "foo"}, &model.Series{ID: 2, Name: "bar"}}
	lparam := app.SeriesCollection{convert.ToSeriesMedia(lret[0]), convert.ToSeriesMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockseriesResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, nil),
		mock.EXPECT().GetRoleByID(4).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, nil),
		mock.EXPECT().GetRoleByID(4).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, nil),
		mock.EXPECT().GetRoleByID(4).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, &categoryID, &classID, &roleID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetAuthorByID(1).Return(nil, nil),
		mock.EXPECT().GetCategoryByID(2).Return(nil, nil),
		mock.EXPECT().GetClassByID(3).Return(nil, nil),
		mock.EXPECT().GetRoleByID(4).Return(nil, nil),
		mock.EXPECT().ListSeriesByIDs(&authorID, &categoryID, &classID, &roleID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"author error", rfm, "[EROR] failed to get author error=get error\n"},
		{"author not found", rfm, "[EROR] failed to get author error=not found\n"},
		{"category error", rfm, "[EROR] failed to get category error=get error\n"},
		{"category not found", rfm, "[EROR] failed to get category error=not found\n"},
		{"class error", rfm, "[EROR] failed to get class error=get error\n"},
		{"class not found", rfm, "[EROR] failed to get class error=not found\n"},
		{"role error", rfm, "[EROR] failed to get role error=get error\n"},
		{"role not found", rfm, "[EROR] failed to get role error=not found\n"},
		{"list error", rfm, "[EROR] failed to get series list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListSeries(ctx, tests[i].fm, resp, &authorID, &categoryID, &classID, &roleID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}

func TestBuilderListSeriesByEditionIDs(t *testing.T) {
	var collectionID, editorID, printID int = 1, 2, 3
	lret := []*model.Series{&model.Series{ID: 1, Name: "foo"}, &model.Series{ID: 2, Name: "bar"}}
	lparam := app.SeriesCollection{convert.ToSeriesMedia(lret[0]), convert.ToSeriesMedia(lret[1])}
	mctrl := gomock.NewController(t)

	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	resp := NewMockseriesResponse(mctrl)
	gomock.InOrder(
		resp.EXPECT().ServiceUnavailable(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, errors.New("get error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, model.ErrNotFound),
		resp.EXPECT().NotFound(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().ListSeriesByEditionIDs(&collectionID, &editorID, &printID).Return(nil, errors.New("list error")),
		resp.EXPECT().InternalServerError(),
		mock.EXPECT().Close(),

		mock.EXPECT().GetCollectionByID(1).Return(nil, nil),
		mock.EXPECT().GetEditorByID(2).Return(nil, nil),
		mock.EXPECT().GetPrintByID(3).Return(nil, nil),
		mock.EXPECT().ListSeriesByEditionIDs(&collectionID, &editorID, &printID).Return(lret, nil),
		resp.EXPECT().OK(lparam),
		mock.EXPECT().Close(),
	)
	wfm := Fmodeler(func() (Modeler, error) { return nil, errors.New("model error") })
	rfm := Fmodeler(func() (Modeler, error) { return mock, nil })
	tests := []struct {
		desc string
		fm   Fmodeler
		log  string
	}{
		{"model error", wfm, "[EROR] unable to get model error=model error\n"},
		{"collection error", rfm, "[EROR] failed to get collection error=get error\n"},
		{"collection not found", rfm, "[EROR] failed to get collection error=not found\n"},
		{"editor error", rfm, "[EROR] failed to get editor error=get error\n"},
		{"editor not found", rfm, "[EROR] failed to get editor error=not found\n"},
		{"print error", rfm, "[EROR] failed to get print error=get error\n"},
		{"print not found", rfm, "[EROR] failed to get print error=not found\n"},
		{"list error", rfm, "[EROR] failed to get series list error=list error\n"},
		{"list ok", rfm, ""},
	}
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	lb := ListBuilder{}
	for i := range tests {
		logbuf.Reset()
		lb.ListSeriesByEditionIDs(ctx, tests[i].fm, resp, &collectionID, &editorID, &printID)
		if tests[i].log != logbuf.String() {
			t.Errorf("unexpected log for [%s]\n exp [%s]\ngot [%s]", tests[i].desc, tests[i].log, logbuf.String())
		}
	}
}
