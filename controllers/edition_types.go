package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// EditionTypesController implements the edition_types resource.
type EditionTypesController struct {
	*goa.Controller
}

// NewEditionTypesController creates a edition_types controller.
func NewEditionTypesController(service *goa.Service) *EditionTypesController {
	return &EditionTypesController{Controller: service.NewController("EditionTypesController")}
}

// Create runs the create action.
func (c *EditionTypesController) Create(ctx *app.CreateEditionTypesContext) error {
	// EditionTypesController_Create: start_implement

	// Put your logic here

	return nil
	// EditionTypesController_Create: end_implement
}

// Delete runs the delete action.
func (c *EditionTypesController) Delete(ctx *app.DeleteEditionTypesContext) error {
	// EditionTypesController_Delete: start_implement

	// Put your logic here

	return nil
	// EditionTypesController_Delete: end_implement
}

// List runs the list action.
func (c *EditionTypesController) List(ctx *app.ListEditionTypesContext) error {
	// EditionTypesController_List: start_implement

	// Put your logic here

	res := app.EditiontypeCollection{}
	return ctx.OK(res)
	// EditionTypesController_List: end_implement
}

// Show runs the show action.
func (c *EditionTypesController) Show(ctx *app.ShowEditionTypesContext) error {
	// EditionTypesController_Show: start_implement

	// Put your logic here

	res := &app.Editiontype{}
	return ctx.OK(res)
	// EditionTypesController_Show: end_implement
}

// Update runs the update action.
func (c *EditionTypesController) Update(ctx *app.UpdateEditionTypesContext) error {
	// EditionTypesController_Update: start_implement

	// Put your logic here

	return nil
	// EditionTypesController_Update: end_implement
}
