package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// AuthenticateController implements the authenticate resource.
type AuthenticateController struct {
	*goa.Controller
}

// NewAuthenticateController creates a authenticate controller.
func NewAuthenticateController(service *goa.Service) *AuthenticateController {
	return &AuthenticateController{Controller: service.NewController("AuthenticateController")}
}

// Auth runs the auth action.
func (c *AuthenticateController) Auth(ctx *app.AuthAuthenticateContext) error {
	// AuthenticateController_Auth: start_implement

	// Put your logic here

	// AuthenticateController_Auth: end_implement
	res := &app.Token{}
	return ctx.OK(res)
}
