package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// EditionsController implements the editions resource.
type EditionsController struct {
	*goa.Controller
}

// NewEditionsController creates a editions controller.
func NewEditionsController(service *goa.Service) *EditionsController {
	return &EditionsController{Controller: service.NewController("EditionsController")}
}

// Create runs the create action.
func (c *EditionsController) Create(ctx *app.CreateEditionsContext) error {
	// EditionsController_Create: start_implement

	// Put your logic here

	return nil
	// EditionsController_Create: end_implement
}

// Delete runs the delete action.
func (c *EditionsController) Delete(ctx *app.DeleteEditionsContext) error {
	// EditionsController_Delete: start_implement

	// Put your logic here

	return nil
	// EditionsController_Delete: end_implement
}

// List runs the list action.
func (c *EditionsController) List(ctx *app.ListEditionsContext) error {
	// EditionsController_List: start_implement

	// Put your logic here

	res := app.EditionCollection{}
	return ctx.OK(res)
	// EditionsController_List: end_implement
}

// Show runs the show action.
func (c *EditionsController) Show(ctx *app.ShowEditionsContext) error {
	// EditionsController_Show: start_implement

	// Put your logic here

	res := &app.Edition{}
	return ctx.OK(res)
	// EditionsController_Show: end_implement
}
