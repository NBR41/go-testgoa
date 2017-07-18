package main

import (
	"github.com/NBR41/go-testgoa/app"
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

	// Put your logic here

	// PasswordController_Get: end_implement
	return nil
}

// Update runs the update action.
func (c *PasswordController) Update(ctx *app.UpdatePasswordContext) error {
	// PasswordController_Update: start_implement

	// Put your logic here

	// PasswordController_Update: end_implement
	return nil
}
