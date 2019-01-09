package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// ValidationController implements the validation resource.
type ValidationController struct {
	*goa.Controller
	fm   Fmodeler
	tok  TokenHelper
	mail MailSender
}

// NewValidationController creates a validation controller.
func NewValidationController(service *goa.Service, fm Fmodeler, tok TokenHelper, mail MailSender) *ValidationController {
	return &ValidationController{
		Controller: service.NewController("ValidationController"),
		fm:         fm,
		tok:        tok,
		mail:       mail,
	}
}

// Get runs the get action.
func (c *ValidationController) Get(ctx *app.GetValidationContext) error {
	// ValidationController_Get: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByID(int(ctx.UserID))
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get user`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	token, err := c.tok.GetValidationToken(u.ID, u.Email)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get validation token`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	err = c.mail.SendActivationMail(u, token)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to send validation email`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	return ctx.NoContent()
	// ValidationController_Get: end_implement
}

// Validate runs the validate action.
func (c *ValidationController) Validate(ctx *app.ValidateValidationContext) error {
	// ValidationController_Validate: start_implement
	uID, uEmail, err := c.tok.ValidateValidationToken(ctx.Payload.Token)
	if err != nil {
		goa.ContextLogger(ctx).Error(`invalid validation token`, `error`, err.Error())
		return ctx.UnprocessableEntity()
	}

	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateUserActivation(int(uID), true)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to activate user`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	err = c.mail.SendUserActivatedMail(uEmail)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to send activated email`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	return ctx.NoContent()
	// ValidationController_Validate: end_implement
}
