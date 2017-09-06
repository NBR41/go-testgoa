package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmail"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/NBR41/go-testgoa/appsec"
	"github.com/goadesign/goa"
)

// ValidationController implements the validation resource.
type ValidationController struct {
	*goa.Controller
}

// NewValidationController creates a validation controller.
func NewValidationController(service *goa.Service) *ValidationController {
	return &ValidationController{Controller: service.NewController("ValidationController")}
}

// Get runs the get action.
func (c *ValidationController) Get(ctx *app.GetValidationContext) error {
	// ValidationController_Get: start_implement

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByID(int(ctx.UserID))
	if err != nil {
		if err == appmodel.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}
	token, err := appsec.GetValidationToken(u.ID, u.Email)
	if err != nil {
		return ctx.InternalServerError()
	}

	err = appmail.SendActivationMail(u, token)
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.NoContent()
	// ValidationController_Get: end_implement
}

// Validate runs the validate action.
func (c *ValidationController) Validate(ctx *app.ValidateValidationContext) error {
	// ValidationController_Validate: start_implement

	// Put your logic here
	uID, uEmail, err := appsec.ValidateValidationToken(ctx.Payload.Token)
	if err != nil {
		return ctx.UnprocessableEntity()
	}

	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateUserActivation(int(uID), true)
	if err != nil {
		if err == appmodel.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}
	_ = appmail.SendUserActivatedMail(uEmail)
	return ctx.NoContent()
	// ValidationController_Validate: end_implement
}
