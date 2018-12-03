package mail

import (
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/internal/mail/content"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/golang/mock/gomock"
)

var (
	exp  = errors.New("send error")
	mail = &content.Mail{Subject: "corge", Body: "grault"}
)

func TestSendResetPasswordMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmock := NewMockContenter(ctrl)
	amock := NewMockActionner(ctrl)
	gomock.InOrder(
		cmock.EXPECT().GetResetPasswordMail("baz").Return(mail),
		amock.EXPECT().Do("foo@bar.com", mail).Return(exp),
	)
	m := New(cmock, amock)
	err := m.SendResetPasswordMail("foo@bar.com", "baz")
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != exp {
			t.Errorf("unexpected error, exp [%v] got [%v]", exp, err)
		}
	}
}

func TestSendPasswordUpdatedMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmock := NewMockContenter(ctrl)
	amock := NewMockActionner(ctrl)
	gomock.InOrder(
		cmock.EXPECT().GetPasswordUpdatedMail().Return(mail),
		amock.EXPECT().Do("foo@bar.com", mail).Return(exp),
	)
	m := New(cmock, amock)
	err := m.SendPasswordUpdatedMail("foo@bar.com")
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != exp {
			t.Errorf("unexpected error, exp [%v] got [%v]", exp, err)
		}
	}
}

func TestSendNewUserMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmock := NewMockContenter(ctrl)
	amock := NewMockActionner(ctrl)
	gomock.InOrder(
		cmock.EXPECT().GetNewUserMail("baz", "qux").Return(mail),
		amock.EXPECT().Do("foo@bar.com", mail).Return(exp),
	)
	m := New(cmock, amock)
	err := m.SendNewUserMail(&model.User{Email: "foo@bar.com", Nickname: "baz"}, "qux")
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != exp {
			t.Errorf("unexpected error, exp [%v] got [%v]", exp, err)
		}
	}
}

func TestSendActivationMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmock := NewMockContenter(ctrl)
	amock := NewMockActionner(ctrl)
	gomock.InOrder(
		cmock.EXPECT().GetActivationMail("baz", "qux").Return(mail),
		amock.EXPECT().Do("foo@bar.com", mail).Return(exp),
	)
	m := New(cmock, amock)
	err := m.SendActivationMail(&model.User{Email: "foo@bar.com", Nickname: "baz"}, "qux")
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != exp {
			t.Errorf("unexpected error, exp [%v] got [%v]", exp, err)
		}
	}
}

func TestSendUserActivatedMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmock := NewMockContenter(ctrl)
	amock := NewMockActionner(ctrl)
	gomock.InOrder(
		cmock.EXPECT().GetUserActivatedMail().Return(mail),
		amock.EXPECT().Do("foo@bar.com", mail).Return(exp),
	)
	m := New(cmock, amock)
	err := m.SendUserActivatedMail("foo@bar.com")
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != exp {
			t.Errorf("unexpected error, exp [%v] got [%v]", exp, err)
		}
	}
}
