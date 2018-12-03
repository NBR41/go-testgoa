package content

import (
	"reflect"
	"testing"
)

func TestGetResetPasswordMail(t *testing.T) {
	exp := &Mail{Subject: "MyInventory: Reset Password", Body: "You have forgotten your password ?\r\nCopy the link below in your favorite browser:\r\n\r\nhttp://foo.com/password/reset?t=bar"}
	f := NewFactory("http://foo.com")
	v := f.GetResetPasswordMail("bar")
	if !reflect.DeepEqual(exp, v) {
		t.Errorf("unexpected value, exp [%+v] got [%+v]", exp, v)
	}
}

func TestGetPasswordUpdatedMail(t *testing.T) {
	exp := &Mail{Subject: "MyInventory: Password update", Body: "Your password has been successfully updated.\r\nVisit: http://foo.com"}
	f := NewFactory("http://foo.com")
	v := f.GetPasswordUpdatedMail()
	if !reflect.DeepEqual(exp, v) {
		t.Errorf("unexpected value, exp [%+v] got [%+v]", exp, v)
	}
}

func TestGetNewUserMail(t *testing.T) {
	exp := &Mail{Subject: "MyInventory: New Account", Body: "Hello bar,\r\nYour account has been successfully created.\r\nVisit to activate your account: http://foo.com/user/validate?t=baz"}
	f := NewFactory("http://foo.com")
	v := f.GetNewUserMail("bar", "baz")
	if !reflect.DeepEqual(exp, v) {
		t.Errorf("unexpected value, exp [%+v] got [%+v]", exp, v)
	}
}

func TestGetActivationMail(t *testing.T) {
	exp := &Mail{Subject: "MyInventory: Validate your Account", Body: "Hello bar,\r\nYour account need to be validated.\r\nCopy the link below in your favorite browser:\r\n\r\nhttp://foo.com/user/validate?t=baz"}
	f := NewFactory("http://foo.com")
	v := f.GetActivationMail("bar", "baz")
	if !reflect.DeepEqual(exp, v) {
		t.Errorf("unexpected value, exp [%+v] got [%+v]", exp, v)
	}
}

func TestGetUserActivatedMail(t *testing.T) {
	exp := &Mail{Subject: "MyInventory: Your account is validated", Body: "Your account has been successfully validated.\r\nVisit: http://foo.com"}
	f := NewFactory("http://foo.com")
	v := f.GetUserActivatedMail()
	if !reflect.DeepEqual(exp, v) {
		t.Errorf("unexpected value, exp [%+v] got [%+v]", exp, v)
	}
}
