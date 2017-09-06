package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmail"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/NBR41/go-testgoa/appsec"
	"github.com/goadesign/goa"
)

// PasswordController implements the password resource.
type PasswordController struct {
	*goa.Controller
}

// NewPasswordController creates a password controller.
func NewPasswordController(service *goa.Service) *PasswordController {
	return &PasswordController{Controller: service.NewController("PasswordController")}
}

// Get runs the get action.
func (c *PasswordController) Get(ctx *app.GetPasswordContext) error {
	// PasswordController_Get: start_implement
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByEmail(ctx.Email)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get user`, `error`, err)
		if err == appmodel.ErrNotFound {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	token, err := appsec.GetPasswordToken(u.ID, u.Email)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get password token`, `error`, err)
		return ctx.InternalServerError()
	}

	err = appmail.SendResetPasswordMail(u.Email, token)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to send password email`, `error`, err)
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// PasswordController_Get: end_implement
}

// Update runs the update action.
func (c *PasswordController) Update(ctx *app.UpdatePasswordContext) error {
	// PasswordController_Update: start_implement
	uID, uEmail, err := appsec.ValidatePasswordToken(ctx.Payload.Token)
	if err != nil {
		goa.ContextLogger(ctx).Error(`invalid password token`, `error`, err)
		return ctx.UnprocessableEntity()
	}

	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateUserPassword(int(uID), ctx.Payload.Password)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to update user password`, `error`, err)
		if err == appmodel.ErrNotFound {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	_ = appmail.SendPasswordUpdatedMail(uEmail)
	return ctx.NoContent()
	// PasswordController_Update: end_implement
}
