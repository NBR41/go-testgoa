package content

import (
	"fmt"
)

var (
	fmtResetPasswordBody    = "You have forgotten your password ?\r\nCopy the link below in your favorite browser:\r\n\r\n%s/password/reset?t=%s"
	fmtModifiedPasswordBody = "Your password has been successfully updated.\r\nVisit: %s"
	fmtNewUserBody          = "Hello %s,\r\nYour account has been successfully created.\r\nVisit to activate your account: %s/user/validate?t=%s"
	fmtValidateUserBody     = "Hello %s,\r\nYour account need to be validated.\r\nCopy the link below in your favorite browser:\r\n\r\n%s/user/validate?t=%s"
	fmtValidatedUserBody    = "Your account has been successfully validated.\r\nVisit: %s"
)

type Mail struct {
	Subject string
	Body    string
}

type Factory struct {
	url string
}

func NewFactory(url string) *Factory {
	return &Factory{url: url}
}

// GetResetPasswordMail get reset password link mail
func (f *Factory) GetResetPasswordMail(token string) *Mail {
	return &Mail{Subject: "MyInventory: Reset Password", Body: fmt.Sprintf(fmtResetPasswordBody, f.url, token)}
}

// GetPasswordUpdatedMail send reset password notification mail
func (f *Factory) GetPasswordUpdatedMail() *Mail {
	return &Mail{Subject: "MyInventory: Password update", Body: fmt.Sprintf(fmtModifiedPasswordBody, f.url)}
}

// SendNewUserMail send user creation mail
func (f *Factory) GetNewUserMail(nickname, token string) *Mail {
	return &Mail{Subject: "MyInventory: New Account", Body: fmt.Sprintf(fmtNewUserBody, nickname, f.url, token)}
}

// SendActivationMail send user activation mail
func (f *Factory) GetActivationMail(nickname, token string) *Mail {
	return &Mail{Subject: "MyInventory: Validate your Account", Body: fmt.Sprintf(fmtValidateUserBody, nickname, f.url, token)}
}

// SendUserActivatedMail send activated user notification mail
func (f *Factory) GetUserActivatedMail() *Mail {
	return &Mail{Subject: "MyInventory: Your account is validated", Body: fmt.Sprintf(fmtValidatedUserBody, f.url)}
}
