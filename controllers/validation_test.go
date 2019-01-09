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

func TestValidationGet(t *testing.T) {
	u := &model.User{ID: 123, Email: "foo@bar.com"}
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
		tmock.EXPECT().GetValidationToken(u.ID, u.Email).Return("", errors.New("valid token error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByID(123).Return(u, nil),
		tmock.EXPECT().GetValidationToken(u.ID, u.Email).Return("valid token", nil),
		mamock.EXPECT().SendActivationMail(u, "valid token").Return(errors.New("mail error")),
		mmock.EXPECT().Close(),
		mmock.EXPECT().GetUserByID(123).Return(u, nil),
		tmock.EXPECT().GetValidationToken(u.ID, u.Email).Return("valid token", nil),
		mamock.EXPECT().SendActivationMail(u, "valid token").Return(nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewValidationController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)
	test.GetValidationServiceUnavailable(t, ctx, service, ctrl, 123)
	exp := "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewValidationController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.GetValidationInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to get user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.GetValidationNotFound(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to get user error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.GetValidationInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] failed to get validation token error=valid token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.GetValidationInternalServerError(t, ctx, service, ctrl, 123)
	exp = "[EROR] unable to send validation email error=mail error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.GetValidationNoContent(t, ctx, service, ctrl, 123)
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}

func TestValidationValidate(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	tmock := NewMockTokenHelper(mctrl)
	mamock := NewMockMailSender(mctrl)
	gomock.InOrder(
		tmock.EXPECT().ValidateValidationToken("foo").Return(int64(0), "", errors.New("validation token error")),
		tmock.EXPECT().ValidateValidationToken("foo").Return(int64(123), "foo@bar.com", nil),
		tmock.EXPECT().ValidateValidationToken("foo").Return(int64(123), "foo@bar.com", nil),
		mmock.EXPECT().UpdateUserActivation(123, true).Return(errors.New("get error")),
		mmock.EXPECT().Close(),
		tmock.EXPECT().ValidateValidationToken("foo").Return(int64(123), "foo@bar.com", nil),
		mmock.EXPECT().UpdateUserActivation(123, true).Return(model.ErrNotFound),
		mmock.EXPECT().Close(),
		tmock.EXPECT().ValidateValidationToken("foo").Return(int64(123), "foo@bar.com", nil),
		mmock.EXPECT().UpdateUserActivation(123, true).Return(nil),
		mamock.EXPECT().SendUserActivatedMail("foo@bar.com").Return(errors.New("mail error")),
		mmock.EXPECT().Close(),
		tmock.EXPECT().ValidateValidationToken("foo").Return(int64(123), "foo@bar.com", nil),
		mmock.EXPECT().UpdateUserActivation(123, true).Return(nil),
		mamock.EXPECT().SendUserActivatedMail("foo@bar.com").Return(nil),
		mmock.EXPECT().Close(),
	)

	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewValidationController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}), tmock, mamock)

	logbuf.Reset()
	test.ValidateValidationUnprocessableEntity(t, ctx, service, ctrl, &app.ValidateValidationPayload{Token: "foo"})
	exp := "[EROR] invalid validation token error=validation token error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ValidateValidationServiceUnavailable(t, ctx, service, ctrl, &app.ValidateValidationPayload{Token: "foo"})
	exp = "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewValidationController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}), tmock, mamock)

	logbuf.Reset()
	test.ValidateValidationInternalServerError(t, ctx, service, ctrl, &app.ValidateValidationPayload{Token: "foo"})
	exp = "[EROR] unable to activate user error=get error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ValidateValidationNotFound(t, ctx, service, ctrl, &app.ValidateValidationPayload{Token: "foo"})
	exp = "[EROR] unable to activate user error=not found\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ValidateValidationInternalServerError(t, ctx, service, ctrl, &app.ValidateValidationPayload{Token: "foo"})
	exp = "[EROR] unable to send activated email error=mail error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	logbuf.Reset()
	test.ValidateValidationNoContent(t, ctx, service, ctrl, &app.ValidateValidationPayload{Token: "foo"})
	exp = ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
