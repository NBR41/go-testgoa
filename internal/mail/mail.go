package mail

import (
	"github.com/NBR41/go-testgoa/internal/model"

	"github.com/NBR41/go-testgoa/internal/mail/content"
)

type Actionner interface {
	Do(email string, cnt *content.Mail) error
}

type Contenter interface {
	GetResetPasswordMail(token string) *content.Mail
	GetPasswordUpdatedMail() *content.Mail
	GetNewUserMail(nickname, token string) *content.Mail
	GetActivationMail(nickname, token string) *content.Mail
	GetUserActivatedMail() *content.Mail
}

type Mailer struct {
	cnt Contenter
	act Actionner
}

func New(cnt Contenter, act Actionner) *Mailer {
	return &Mailer{cnt: cnt, act: act}
}

// SendResetPasswordMail send reset password link mail
func (m *Mailer) SendResetPasswordMail(email, token string) error {
	return m.act.Do(email, m.cnt.GetResetPasswordMail(token))
}

// SendPasswordUpdatedMail send reset password notification mail
func (m *Mailer) SendPasswordUpdatedMail(email string) error {
	return m.act.Do(email, m.cnt.GetPasswordUpdatedMail())
}

// SendNewUserMail send user creation mail
func (m *Mailer) SendNewUserMail(u *model.User, token string) error {
	return m.act.Do(u.Email, m.cnt.GetNewUserMail(u.Nickname, token))
}

// SendActivationMail send user activation mail
func (m *Mailer) SendActivationMail(u *model.User, token string) error {
	return m.act.Do(u.Email, m.cnt.GetActivationMail(u.Nickname, token))
}

// SendUserActivatedMail send activated user notification mail
func (m *Mailer) SendUserActivatedMail(email string) error {
	return m.act.Do(email, m.cnt.GetUserActivatedMail())
}
