package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RolesController implements the roles resource.
type RolesController struct {
	*goa.Controller
}

// NewRolesController creates a roles controller.
func NewRolesController(service *goa.Service) *RolesController {
	return &RolesController{Controller: service.NewController("RolesController")}
}

// Create runs the create action.
func (c *RolesController) Create(ctx *app.CreateRolesContext) error {
	// RolesController_Create: start_implement

	// Put your logic here

	return nil
	// RolesController_Create: end_implement
}

// Delete runs the delete action.
func (c *RolesController) Delete(ctx *app.DeleteRolesContext) error {
	// RolesController_Delete: start_implement

	// Put your logic here

	return nil
	// RolesController_Delete: end_implement
}

// List runs the list action.
func (c *RolesController) List(ctx *app.ListRolesContext) error {
	// RolesController_List: start_implement

	// Put your logic here

	res := app.RoleCollection{}
	return ctx.OK(res)
	// RolesController_List: end_implement
}

// Show runs the show action.
func (c *RolesController) Show(ctx *app.ShowRolesContext) error {
	// RolesController_Show: start_implement

	// Put your logic here

	res := &app.Role{}
	return ctx.OK(res)
	// RolesController_Show: end_implement
}

// Update runs the update action.
func (c *RolesController) Update(ctx *app.UpdateRolesContext) error {
	// RolesController_Update: start_implement

	// Put your logic here

	return nil
	// RolesController_Update: end_implement
}
