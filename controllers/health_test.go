package controllers

import (
	"context"
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/NBR41/go-testgoa/app/test"
	"github.com/goadesign/goa"
	"github.com/golang/mock/gomock"
)

func TestHealthHealth(t *testing.T) {
	mctrl := gomock.NewController(t)
	defer mctrl.Finish()
	mmock := NewMockModeler(mctrl)
	gomock.InOrder(
		mmock.EXPECT().Close(),
	)
	service := goa.New("my-inventory-test")
	logbuf := &strings.Builder{}
	ctx := goa.WithLogger(context.Background(), goa.NewLogger(log.New(logbuf, "", 0)))

	ctrl := NewHealthController(service, Fmodeler(func() (Modeler, error) {
		return mmock, nil
	}))
	logbuf.Reset()
	test.HealthHealthOK(t, ctx, service, ctrl)
	exp := ""
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}

	ctrl = NewHealthController(service, Fmodeler(func() (Modeler, error) {
		return nil, errors.New("model error")
	}))
	logbuf.Reset()
	test.HealthHealthInternalServerError(t, ctx, service, ctrl)
	exp = "[EROR] unable to get model error=model error\n"
	if exp != logbuf.String() {
		t.Errorf("unexpected log\n exp [%s]\ngot [%s]", exp, logbuf.String())
	}
}
