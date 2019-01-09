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

func TestAuthenticateAuth(t *testing.T) {
	u := &model.User{ID: 1, Nickname: "foo", Email: "foo@bar.com", IsAdmin: true}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	gomock.InOrder(
		mmock.EXPECT().GetAuthenticatedUser("foobar", "baz").Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetAuthenticatedUser("foobar", "baz").Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetAuthenticatedUser("foobar", "baz").Return(nil, model.ErrInvalidCredentials),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetAuthenticatedUser("foobar", "baz").Return(u, nil),
		tmock.EXPECT().GetAuthToken(u.ID, u.IsAdmin).Return("", errors.New("access token error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetAuthenticatedUser("foobar", "baz").Return(u, nil),
		tmock.EXPECT().GetAuthToken(u.ID, u.IsAdmin).Return("access token", nil),
		tmock.EXPECT().GetRefreshToken(u.ID, u.IsAdmin).Return("", errors.New("refresh token error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetAuthenticatedUser("foobar", "baz").Return(u, nil),
		tmock.EXPECT().GetAuthToken(u.ID, u.IsAdmin).Return("access token", nil),
		tmock.EXPECT().GetRefreshToken(u.ID, u.IsAdmin).Return("refresh token", nil),
		mmock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewAuthenticateController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock)
	test.AuthAuthenticateServiceUnavailable(t, ctx, service, ctrl, &app.AuthenticatePayload{Login: "foobar", Password: "baz"})
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewAuthenticateController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock)

	logbuf.Reset()
	test.AuthAuthenticateInternalServerError(t, ctx, service, ctrl, &app.AuthenticatePayload{Login: "foobar", Password: "baz"})
	exp = "[EROR] failed to auth error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AuthAuthenticateUnprocessableEntity(t, ctx, service, ctrl, &app.AuthenticatePayload{Login: "foobar", Password: "baz"})
	exp = "[EROR] failed to auth error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AuthAuthenticateUnprocessableEntity(t, ctx, service, ctrl, &app.AuthenticatePayload{Login: "foobar", Password: "baz"})
	exp = "[EROR] failed to auth error=invalid credentials\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AuthAuthenticateInternalServerError(t, ctx, service, ctrl, &app.AuthenticatePayload{Login: "foobar", Password: "baz"})
	exp = "[EROR] failed to get access token error=access token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.AuthAuthenticateInternalServerError(t, ctx, service, ctrl, &app.AuthenticatePayload{Login: "foobar", Password: "baz"})
	exp = "[EROR] failed to get refresh token error=refresh token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.AuthAuthenticateOK(t, ctx, service, ctrl, &app.AuthenticatePayload{Login: "foobar", Password: "baz"})
	expres := convert.ToAuthTokenMedia(u, "access token", "refresh token")
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
