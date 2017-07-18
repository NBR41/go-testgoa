package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service) *UsersController {
	return &UsersController{Controller: service.NewController("UsersController")}
}

// Create runs the create action.
func (c *UsersController) Create(ctx *app.CreateUsersContext) error {
	// UsersController_Create: start_implement

	// Put your logic here

	// UsersController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *UsersController) Delete(ctx *app.DeleteUsersContext) error {
	// UsersController_Delete: start_implement

	// Put your logic here

	// UsersController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *UsersController) List(ctx *app.ListUsersContext) error {
	// UsersController_List: start_implement

	// Put your logic here

	// UsersController_List: end_implement
	res := app.UserCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	// UsersController_Show: start_implement

	// Put your logic here

	// UsersController_Show: end_implement
	res := &app.User{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UsersController) Update(ctx *app.UpdateUsersContext) error {
	// UsersController_Update: start_implement

	// Put your logic here

	// UsersController_Update: end_implement
	return nil
}
