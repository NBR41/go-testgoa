package controllers

import (
	"context"
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/app/test"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestPasswordGet(t *testing.T) {
	u := &model.User{ID: 1, Email: "foo@bar.com"}
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	mamock := NewMockMailSender(mctrl)
	gomock.InOrder(
		mmock.EXPECT().GetUserByEmail("foo@bar.com").Return(nil, errors.New("get error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByEmail("foo@bar.com").Return(nil, model.ErrNotFound),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByEmail("foo@bar.com").Return(u, nil),
		tmock.EXPECT().GetPasswordToken(u.ID, u.Email).Return("", errors.New("password token error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByEmail("foo@bar.com").Return(u, nil),
		tmock.EXPECT().GetPasswordToken(u.ID, u.Email).Return("password token", nil),
		mamock.EXPECT().SendResetPasswordMail(u.Email, "password token").Return(errors.New("password email error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByEmail("foo@bar.com").Return(u, nil),
		tmock.EXPECT().GetPasswordToken(u.ID, u.Email).Return("password token", nil),
		mamock.EXPECT().SendResetPasswordMail(u.Email, "password token").Return(nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewPasswordController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)
	test.GetPasswordServiceUnavailable(t, ctx, service, ctrl, "foo@bar.com")
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewPasswordController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.GetPasswordInternalServerError(t, ctx, service, ctrl, "foo@bar.com")
	exp = "[EROR] unable to get user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.GetPasswordUnprocessableEntity(t, ctx, service, ctrl, "foo@bar.com")
	exp = "[EROR] unable to get user error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.GetPasswordInternalServerError(t, ctx, service, ctrl, "foo@bar.com")
	exp = "[EROR] failed to get password token error=password token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.GetPasswordInternalServerError(t, ctx, service, ctrl, "foo@bar.com")
	exp = "[EROR] unable to send password email error=password email error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.GetPasswordNoContent(t, ctx, service, ctrl, "foo@bar.com")
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestPasswordUpdate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	mamock := NewMockMailSender(mctrl)
	gomock.InOrder(
		tmock.EXPECT().ValidatePasswordToken("bar").Return(int64(0), "", errors.New("password token error")),
		tmock.EXPECT().ValidatePasswordToken("bar").Return(int64(1), "foo@bar.com", nil),
		tmock.EXPECT().ValidatePasswordToken("bar").Return(int64(1), "foo@bar.com", nil),
		mmock.EXPECT().UpdateUserPassword(1, "foo").Return(errors.New("update error")),
		mmock.EXPECT().Close(),
		tmock.EXPECT().ValidatePasswordToken("bar").Return(int64(1), "foo@bar.com", nil),
		mmock.EXPECT().UpdateUserPassword(1, "foo").Return(model.ErrNotFound),
		mmock.EXPECT().Close(),
		tmock.EXPECT().ValidatePasswordToken("bar").Return(int64(1), "foo@bar.com", nil),
		mmock.EXPECT().UpdateUserPassword(1, "foo").Return(nil),
		mamock.EXPECT().SendPasswordUpdatedMail("foo@bar.com").Return(errors.New("email error")),
		mmock.EXPECT().Close(),
		tmock.EXPECT().ValidatePasswordToken("bar").Return(int64(1), "foo@bar.com", nil),
		mmock.EXPECT().UpdateUserPassword(1, "foo").Return(nil),
		mamock.EXPECT().SendPasswordUpdatedMail("foo@bar.com").Return(nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewPasswordController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)

	logbuf.Reset()
	test.UpdatePasswordUnprocessableEntity(t, ctx, service, ctrl, &app.PasswordChangePayload{Password: "foo", Token: "bar"})
	exp := "[EROR] invalid password token error=password token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdatePasswordServiceUnavailable(t, ctx, service, ctrl, &app.PasswordChangePayload{Password: "foo", Token: "bar"})
	exp = "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewPasswordController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.UpdatePasswordInternalServerError(t, ctx, service, ctrl, &app.PasswordChangePayload{Password: "foo", Token: "bar"})
	exp = "[EROR] unable to update user password error=update error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdatePasswordUnprocessableEntity(t, ctx, service, ctrl, &app.PasswordChangePayload{Password: "foo", Token: "bar"})
	exp = "[EROR] unable to update user password error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdatePasswordInternalServerError(t, ctx, service, ctrl, &app.PasswordChangePayload{Password: "foo", Token: "bar"})
	exp = "[EROR] unable to send password update email error=email error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.UpdatePasswordNoContent(t, ctx, service, ctrl, &app.PasswordChangePayload{Password: "foo", Token: "bar"})
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
