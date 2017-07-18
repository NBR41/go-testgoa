package main

import (
	"github.com/NBR41/go-testgoa/app"
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

	// ValidationController_Get: end_implement
	return nil
}

// Validate runs the validate action.
func (c *ValidationController) Validate(ctx *app.ValidateValidationContext) error {
	// ValidationController_Validate: start_implement

	// Put your logic here

	// ValidationController_Validate: end_implement
	return nil
}
