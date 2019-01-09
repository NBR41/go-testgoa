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

func TestUsersCreate(t *testing.T) {
	payload := &app.UserCreatePayload{Email: "foo@bar.com", Nickname: "foo", Password: "baz"}

	u := &model.User{
		ID:          123,
		Email:       "foo@bar.com",
		Nickname:    "foo",
		IsValidated: false,
		IsAdmin:     true,
	}

	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	mamock := NewMockMailSender(mctrl)
	gomock.InOrder(
		mmock.EXPECT().InsertUser("foo@bar.com", "foo", "baz").Return(nil, model.ErrDuplicateKey),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertUser("foo@bar.com", "foo", "baz").Return(nil, model.ErrDuplicateEmail),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertUser("foo@bar.com", "foo", "baz").Return(nil, model.ErrDuplicateNickname),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertUser("foo@bar.com", "foo", "baz").Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertUser("foo@bar.com", "foo", "baz").Return(u, nil),
		tmock.EXPECT().GetValidationToken(u.ID, u.Email).Return("", errors.New("token error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertUser("foo@bar.com", "foo", "baz").Return(u, nil),
		tmock.EXPECT().GetValidationToken(u.ID, u.Email).Return("valid token", nil),
		mamock.EXPECT().SendNewUserMail(u, "valid token").Return(errors.New("mail error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().InsertUser("foo@bar.com", "foo", "baz").Return(u, nil),
		tmock.EXPECT().GetValidationToken(u.ID, u.Email).Return("valid token", nil),
		mamock.EXPECT().SendNewUserMail(u, "valid token").Return(nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)

	logbuf.Reset()
	test.CreateUsersServiceUnavailable(t, ctx, service, ctrl, payload)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.CreateUsersUnprocessableEntity(t, ctx, service, ctrl, payload)
	exp = "[EROR] unable to insert user error=duplicate key\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateUsersUnprocessableEntity(t, ctx, service, ctrl, payload)
	exp = "[EROR] unable to insert user error=duplicate email\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateUsersUnprocessableEntity(t, ctx, service, ctrl, payload)
	exp = "[EROR] unable to insert user error=duplicate nickname\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.CreateUsersInternalServerError(t, ctx, service, ctrl, payload)
	exp = "[EROR] unable to insert user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	rw := test.CreateUsersCreated(t, ctx, service, ctrl, payload)
	exp = "[EROR] unable to get validation token error=token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
	exp = app.UsersHref(u.ID)
	v := rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}

	logbuf.Reset()
	rw = test.CreateUsersCreated(t, ctx, service, ctrl, payload)
	exp = "[EROR] unable to send user creation email error=mail error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
	exp = app.UsersHref(u.ID)
	v = rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}

	logbuf.Reset()
	rw = test.CreateUsersCreated(t, ctx, service, ctrl, payload)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
	exp = app.UsersHref(u.ID)
	v = rw.Header().Get("Location")
	if exp != v {
		t.Errorf("unexpected value, exp [%s] got [%s]", exp, v)
	}
}

func TestUsersDelete(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	mamock := NewMockMailSender(mctrl)
	gomock.InOrder(
		mmock.EXPECT().DeleteUser(123).Return(model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().DeleteUser(123).Return(errors.New("delete error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().DeleteUser(123).Return(nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)

	logbuf.Reset()
	test.DeleteUsersServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.DeleteUsersNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to delete user error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteUsersInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to delete user error=delete error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.DeleteUsersNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestUsersList(t *testing.T) {
	u := &model.User{
		ID:          123,
		Email:       "foo@bar.com",
		Nickname:    "baz",
		IsValidated: false,
		IsAdmin:     true,
	}
	email := "foo@bar.com"
	nickname := "baz"
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	mamock := NewMockMailSender(mctrl)
	gomock.InOrder(
		mmock.EXPECT().GetUserByEmail(email).Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByNickname(nickname).Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByEmailOrNickname(email, nickname).Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByEmailOrNickname(email, nickname).Return(u, nil),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByEmailOrNickname(email, nickname).Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().ListUsers().Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().ListUsers().Return([]*model.User{u}, nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)

	logbuf.Reset()
	test.ListUsersServiceUnavailable(t, ctx, service, ctrl, &email, &nickname)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.ListUsersInternalServerError(t, ctx, service, ctrl, &email, nil)
	exp = "[EROR] unable to get user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListUsersInternalServerError(t, ctx, service, ctrl, nil, &nickname)
	exp = "[EROR] unable to get user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListUsersInternalServerError(t, ctx, service, ctrl, &email, &nickname)
	exp = "[EROR] unable to get user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ListUsersOK(t, ctx, service, ctrl, &email, &nickname)
	expres := app.UserCollection{convert.ToUserMedia(u)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res = test.ListUsersOK(t, ctx, service, ctrl, &email, &nickname)
	expres = app.UserCollection{}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ListUsersInternalServerError(t, ctx, service, ctrl, nil, nil)
	exp = "[EROR] unable to get user list error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res = test.ListUsersOK(t, ctx, service, ctrl, nil, nil)
	expres = app.UserCollection{convert.ToUserMedia(u)}
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestUsersShow(t *testing.T) {
	u := &model.User{
		ID:          123,
		Email:       "foo@bar.com",
		Nickname:    "baz",
		IsValidated: false,
		IsAdmin:     true,
	}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	mamock := NewMockMailSender(mctrl)
	gomock.InOrder(
		mmock.EXPECT().GetUserByID(123).Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByID(123).Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByID(123).Return(u, nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)

	logbuf.Reset()
	test.ShowUsersServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.ShowUsersInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to get user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ShowUsersNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to get user error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	_, res := test.ShowUsersOK(t, ctx, service, ctrl, 123)
	expres := convert.ToUserMedia(u)
	if diff := pretty.Compare(res, expres); diff != "" {
		t.Errorf("diff: (-got +want)\n%s", diff)
	}
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestUsersUpdate(t *testing.T) {
	payload := &app.UpdateUsersPayload{Nickname: "foo"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	mamock := NewMockMailSender(mctrl)
	gomock.InOrder(
		mmock.EXPECT().UpdateUserNickname(123, "foo").Return(errors.New("update error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().UpdateUserNickname(123, "foo").Return(model.ErrDuplicateNickname),
		mmock.EXPECT().Close(),
		mmock.EXPECT().UpdateUserNickname(123, "foo").Return(model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().UpdateUserNickname(123, "foo").Return(nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)

	logbuf.Reset()
	test.UpdateUsersServiceUnavailable(t, ctx, service, ctrl, 123, payload)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewUsersController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.UpdateUsersInternalServerError(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to update user error=update error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
	logbuf.Reset()
	test.UpdateUsersUnprocessableEntity(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to update user error=duplicate nickname\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateUsersNotFound(t, ctx, service, ctrl, 123, payload)
	exp = "[EROR] unable to update user error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdateUsersNoContent(t, ctx, service, ctrl, 123, payload)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
