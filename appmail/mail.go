package appmail

import (
	"github.com/NBR41/go-testgoa/appmodel"
)

// SendResetPasswordMail send reset password link mail
func SendResetPasswordMail(email, token string) error {
	return nil
}

// SendPasswordUpdatedMail send reset password  notification mail
func SendPasswordUpdatedMail(email string) error {
	return nil
}

// SendNewUserMail send user creation mail
func SendNewUserMail(u *appmodel.User, token string) error {
	return nil
}

// SendActivationMail send user activation mail
func SendActivationMail(u *appmodel.User, token string) error {
	return nil
}

// SendUserActivatedMail send activated user notification mail
func SendUserActivatedMail(email string) error {
	return nil
}
