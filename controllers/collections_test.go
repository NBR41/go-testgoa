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

func TestCollectionsCreate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().InsertCollection("foo", 1).Return(nil, errors.New("insert error")),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertCollection("foo", 1).Return(nil, model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertCollection("foo", 1).Return(nil, model.ErrInvalidID),
		mock.EXPECT().Close(),
		mock.EXPECT().InsertCollection("foo", 1).Return(&model.Collection{ID: 2, Name: "foo", EditorID: 1}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.CreateCollectionsServiceUnavailable(t, ctx, service, ctrl, &app.CreateCollectionsPayload{CollectionName: "foo", EditorID: 1})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.CreateCollectionsInternalServerError(t, ctx, service, ctrl, &app.CreateCollectionsPayload{CollectionName: "foo", EditorID: 1})
	exp = "[EROR] failed to insert collection error=insert error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateCollectionsUnprocessableEntity(t, ctx, service, ctrl, &app.CreateCollectionsPayload{CollectionName: "foo", EditorID: 1})
	exp = "[EROR] failed to insert collection error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateCollectionsUnprocessableEntity(t, ctx, service, ctrl, &app.CreateCollectionsPayload{CollectionName: "foo", EditorID: 1})
	exp = "[EROR] failed to insert collection error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateCollectionsCreated(t, ctx, service, ctrl, &app.CreateCollectionsPayload{CollectionName: "foo", EditorID: 1})
	exp = app.CollectionsHref(2)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestCollectionsDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().DeleteCollection(123).Return(errors.New("delete error")),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteCollection(123).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().DeleteCollection(123).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.DeleteCollectionsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.DeleteCollectionsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete collection error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteCollectionsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to delete collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteCollectionsNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestCollectionsList(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().ListCollections().Return(nil, errors.New("list error")),
		mock.EXPECT().Close(),
		mock.EXPECT().ListCollections().Return([]*model.Collection{&model.Collection{ID: 1, Name: "foo", EditorID: 2}, &model.Collection{ID: 3, Name: "bar", EditorID: 2}}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ListCollectionsServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ListCollectionsInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get collection list error=list error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	_, res := test.ListCollectionsOK(t, ctx, service, ctrl)
	expres := app.CollectionCollection{
		convert.ToCollectionMedia(&model.Collection{ID: 1, Name: "foo", EditorID: 2}),
		convert.ToCollectionMedia(&model.Collection{ID: 3, Name: "bar", EditorID: 2}),
	}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
}

func TestCollectionsShow(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	gomock.InOrder(
		mock.EXPECT().GetCollectionByID(123).Return(nil, errors.New("get error")),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(nil, model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().GetCollectionByID(123).Return(&model.Collection{ID: 123, Name: "foo", EditorID: 2}, nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.ShowCollectionsServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.ShowCollectionsInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get collection error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowCollectionsNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowCollectionsOK(t, ctx, service, ctrl, 123)
	expres := convert.ToCollectionMedia(&model.Collection{ID: 123, Name: "foo", EditorID: 2})
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestCollectionsUpdate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mock := NewMockModeler(mctrl)
	name := "foo"
	editorID := 2
	gomock.InOrder(
		mock.EXPECT().UpdateCollection(123, &name, &editorID).Return(errors.New("update error")),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateCollection(123, &name, &editorID).Return(model.ErrNotFound),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateCollection(123, &name, &editorID).Return(model.ErrInvalidID),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateCollection(123, &name, &editorID).Return(model.ErrDuplicateKey),
		mock.EXPECT().Close(),
		mock.EXPECT().UpdateCollection(123, &name, &editorID).Return(nil),
		mock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	test.UpdateCollectionsServiceUnavailable(t, ctx, service, ctrl, 123, &app.UpdateCollectionsPayload{CollectionName: &name, EditorID: &editorID})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewCollectionsController(service, Fmodeler(func() (Modeler, error) {
		return mock, nil
	}))

	logbuf.Reset()
	test.UpdateCollectionsInternalServerError(t, ctx, service, ctrl, 123, &app.UpdateCollectionsPayload{CollectionName: &name, EditorID: &editorID})
	exp = "[EROR] failed to update collection error=update error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateCollectionsNotFound(t, ctx, service, ctrl, 123, &app.UpdateCollectionsPayload{CollectionName: &name, EditorID: &editorID})
	exp = "[EROR] failed to update collection error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateCollectionsUnprocessableEntity(t, ctx, service, ctrl, 123, &app.UpdateCollectionsPayload{CollectionName: &name, EditorID: &editorID})
	exp = "[EROR] failed to update collection error=invalid id\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateCollectionsUnprocessableEntity(t, ctx, service, ctrl, 123, &app.UpdateCollectionsPayload{CollectionName: &name, EditorID: &editorID})
	exp = "[EROR] failed to update collection error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateCollectionsNoContent(t, ctx, service, ctrl, 123, &app.UpdateCollectionsPayload{CollectionName: &name, EditorID: &editorID})
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
