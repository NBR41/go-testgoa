package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// OwnershipsController implements the ownerships resource.
type OwnershipsController struct {
	*goa.Controller
}

// NewOwnershipsController creates a ownerships controller.
func NewOwnershipsController(service *goa.Service) *OwnershipsController {
	return &OwnershipsController{Controller: service.NewController("OwnershipsController")}
}

// Create runs the create action.
func (c *OwnershipsController) Create(ctx *app.CreateOwnershipsContext) error {
	// OwnershipsController_Create: start_implement

	// Put your logic here

	// OwnershipsController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *OwnershipsController) Delete(ctx *app.DeleteOwnershipsContext) error {
	// OwnershipsController_Delete: start_implement

	// Put your logic here

	// OwnershipsController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *OwnershipsController) List(ctx *app.ListOwnershipsContext) error {
	// OwnershipsController_List: start_implement

	// Put your logic here

	// OwnershipsController_List: end_implement
	res := app.BookCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *OwnershipsController) Show(ctx *app.ShowOwnershipsContext) error {
	// OwnershipsController_Show: start_implement

	// Put your logic here

	// OwnershipsController_Show: end_implement
	res := &app.Ownership{}
	return ctx.OK(res)
}
