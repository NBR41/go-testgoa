package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// PasswordController implements the password resource.
type PasswordController struct {
	*goa.Controller
	fm   Fmodeler
	tok  TokenHelper
	mail MailSender
}

// NewPasswordController creates a password controller.
func NewPasswordController(service *goa.Service, fm Fmodeler, tok TokenHelper, mail MailSender) *PasswordController {
	return &PasswordController{
		Controller: service.NewController("PasswordController"),
		fm:         fm,
		tok:        tok,
		mail:       mail,
	}
}

// Get runs the get action.
func (c *PasswordController) Get(ctx *app.GetPasswordContext) error {
	// PasswordController_Get: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByEmail(ctx.Email)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get user`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	token, err := c.tok.GetPasswordToken(u.ID, u.Email)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get password token`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	err = c.mail.SendResetPasswordMail(u.Email, token)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to send password email`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// PasswordController_Get: end_implement
}

// Update runs the update action.
func (c *PasswordController) Update(ctx *app.UpdatePasswordContext) error {
	// PasswordController_Update: start_implement
	uID, uEmail, err := c.tok.ValidatePasswordToken(ctx.Payload.Token)
	if err != nil {
		goa.ContextLogger(ctx).Error(`invalid password token`, `error`, err.Error())
		return ctx.UnprocessableEntity()
	}

	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateUserPassword(int(uID), ctx.Payload.Password)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to update user password`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	err = c.mail.SendPasswordUpdatedMail(uEmail)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to send password update email`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	return ctx.NoContent()
	// PasswordController_Update: end_implement
}
