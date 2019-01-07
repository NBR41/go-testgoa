package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// ClassificationsController implements the classifications resource.
type ClassificationsController struct {
	*goa.Controller
}

// NewClassificationsController creates a classifications controller.
func NewClassificationsController(service *goa.Service) *ClassificationsController {
	return &ClassificationsController{Controller: service.NewController("ClassificationsController")}
}

// Create runs the create action.
func (c *ClassificationsController) Create(ctx *app.CreateClassificationsContext) error {
	// ClassificationsController_Create: start_implement

	// Put your logic here

	return nil
	// ClassificationsController_Create: end_implement
}

// Delete runs the delete action.
func (c *ClassificationsController) Delete(ctx *app.DeleteClassificationsContext) error {
	// ClassificationsController_Delete: start_implement

	// Put your logic here

	return nil
	// ClassificationsController_Delete: end_implement
}

// List runs the list action.
func (c *ClassificationsController) List(ctx *app.ListClassificationsContext) error {
	// ClassificationsController_List: start_implement

	// Put your logic here

	res := app.ClassCollection{}
	return ctx.OK(res)
	// ClassificationsController_List: end_implement
}

// Show runs the show action.
func (c *ClassificationsController) Show(ctx *app.ShowClassificationsContext) error {
	// ClassificationsController_Show: start_implement

	// Put your logic here

	res := &app.Class{}
	return ctx.OK(res)
	// ClassificationsController_Show: end_implement
}
