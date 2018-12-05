package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// EditorsController implements the editors resource.
type EditorsController struct {
	*goa.Controller
}

// NewEditorsController creates a editors controller.
func NewEditorsController(service *goa.Service) *EditorsController {
	return &EditorsController{Controller: service.NewController("EditorsController")}
}

// Create runs the create action.
func (c *EditorsController) Create(ctx *app.CreateEditorsContext) error {
	// EditorsController_Create: start_implement

	// Put your logic here

	return nil
	// EditorsController_Create: end_implement
}

// Delete runs the delete action.
func (c *EditorsController) Delete(ctx *app.DeleteEditorsContext) error {
	// EditorsController_Delete: start_implement

	// Put your logic here

	return nil
	// EditorsController_Delete: end_implement
}

// List runs the list action.
func (c *EditorsController) List(ctx *app.ListEditorsContext) error {
	// EditorsController_List: start_implement

	// Put your logic here

	res := app.EditorCollection{}
	return ctx.OK(res)
	// EditorsController_List: end_implement
}

// Show runs the show action.
func (c *EditorsController) Show(ctx *app.ShowEditorsContext) error {
	// EditorsController_Show: start_implement

	// Put your logic here

	res := &app.Editor{}
	return ctx.OK(res)
	// EditorsController_Show: end_implement
}

// Update runs the update action.
func (c *EditorsController) Update(ctx *app.UpdateEditorsContext) error {
	// EditorsController_Update: start_implement

	// Put your logic here

	return nil
	// EditorsController_Update: end_implement
}
