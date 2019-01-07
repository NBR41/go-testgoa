package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// AuthorshipsController implements the authorships resource.
type AuthorshipsController struct {
	*goa.Controller
}

// NewAuthorshipsController creates a authorships controller.
func NewAuthorshipsController(service *goa.Service) *AuthorshipsController {
	return &AuthorshipsController{Controller: service.NewController("AuthorshipsController")}
}

// Create runs the create action.
func (c *AuthorshipsController) Create(ctx *app.CreateAuthorshipsContext) error {
	// AuthorshipsController_Create: start_implement

	// Put your logic here

	return nil
	// AuthorshipsController_Create: end_implement
}

// Delete runs the delete action.
func (c *AuthorshipsController) Delete(ctx *app.DeleteAuthorshipsContext) error {
	// AuthorshipsController_Delete: start_implement

	// Put your logic here

	return nil
	// AuthorshipsController_Delete: end_implement
}

// List runs the list action.
func (c *AuthorshipsController) List(ctx *app.ListAuthorshipsContext) error {
	// AuthorshipsController_List: start_implement

	// Put your logic here

	res := app.AuthorshipCollection{}
	return ctx.OK(res)
	// AuthorshipsController_List: end_implement
}

// Show runs the show action.
func (c *AuthorshipsController) Show(ctx *app.ShowAuthorshipsContext) error {
	// AuthorshipsController_Show: start_implement

	// Put your logic here

	res := &app.Authorship{}
	return ctx.OK(res)
	// AuthorshipsController_Show: end_implement
}
