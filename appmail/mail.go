package appmail

import (
	"github.com/NBR41/go-testgoa/store"
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
func SendNewUserMail(u *store.User, token string) error {
	return nil
}

// SendActivationMail send user activation mail
func SendActivationMail(u *store.User, token string) error {
	return nil
}

// SendUserActivatedMail send activated user notification mail
func SendUserActivatedMail(email string) error {
	return nil
}
