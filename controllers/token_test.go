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

func TestTokenAccess(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)

	gomock.InOrder(
		tmock.EXPECT().GetAccessToken(int64(123), true).Return("", errors.New("token error")),
		tmock.EXPECT().GetAccessToken(int64(123), true).Return("auth token", nil),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	mctx := context.Background()
	mctx = context.WithValue(mctx, CtxKey("user_id"), int64(123))
	mctx = context.WithValue(mctx, CtxKey("is_admin"), true)
	ctx := goa.WithLogger(mctx, goa.NewLogger(log.New(logbuf, "", 0)))
	ctrl := NewTokenController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock)

	logbuf.Reset()
	test.AccessTokenInternalServerError(t, ctx, service, ctrl)
	exp := "[EROR] failed to get access token error=token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.AccessTokenOK(t, ctx, service, ctrl)
	expres := &app.Token{Token: "auth token"}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestTokenAuth(t *testing.T) {
	u := &model.User{ID: 123, Nickname: "foo", Email: "foo@bar.com", IsAdmin: true}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)

	gomock.InOrder(
		mmock.EXPECT().GetUserByID(123).Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByID(123).Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByID(123).Return(u, nil),
		tmock.EXPECT().GetAccessToken(u.ID, u.IsAdmin).Return("", errors.New("auth token error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByID(123).Return(u, nil),
		tmock.EXPECT().GetAccessToken(u.ID, u.IsAdmin).Return("auth token", nil),
		tmock.EXPECT().GetRefreshToken(u.ID, u.IsAdmin).Return("", errors.New("refresh token error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByID(123).Return(u, nil),
		tmock.EXPECT().GetAccessToken(u.ID, u.IsAdmin).Return("auth token", nil),
		tmock.EXPECT().GetRefreshToken(u.ID, u.IsAdmin).Return("refresh token", nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	mctx := context.Background()
	mctx = context.WithValue(mctx, CtxKey("user_id"), int64(123))
	ctx := goa.WithLogger(mctx, goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewTokenController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock)
	test.AuthTokenServiceUnavailable(t, ctx, service, ctrl)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewTokenController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock)

	logbuf.Reset()
	test.AuthTokenInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] unable to get user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AuthTokenUnprocessableEntity(t, ctx, service, ctrl)
	exp = "[EROR] unable to get user error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AuthTokenInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get access token error=auth token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AuthTokenInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] failed to get refresh token error=refresh token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.AuthTokenOK(t, ctx, service, ctrl)
	expres := convert.ToAuthTokenMedia(u, "auth token", "refresh token")
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
