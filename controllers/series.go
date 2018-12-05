package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// SeriesController implements the series resource.
type SeriesController struct {
	*goa.Controller
}

// NewSeriesController creates a series controller.
func NewSeriesController(service *goa.Service) *SeriesController {
	return &SeriesController{Controller: service.NewController("SeriesController")}
}

// Create runs the create action.
func (c *SeriesController) Create(ctx *app.CreateSeriesContext) error {
	// SeriesController_Create: start_implement

	// Put your logic here

	return nil
	// SeriesController_Create: end_implement
}

// Delete runs the delete action.
func (c *SeriesController) Delete(ctx *app.DeleteSeriesContext) error {
	// SeriesController_Delete: start_implement

	// Put your logic here

	return nil
	// SeriesController_Delete: end_implement
}

// List runs the list action.
func (c *SeriesController) List(ctx *app.ListSeriesContext) error {
	// SeriesController_List: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// SeriesController_List: end_implement
}

// Show runs the show action.
func (c *SeriesController) Show(ctx *app.ShowSeriesContext) error {
	// SeriesController_Show: start_implement

	// Put your logic here

	res := &app.Series{}
	return ctx.OK(res)
	// SeriesController_Show: end_implement
}

// Update runs the update action.
func (c *SeriesController) Update(ctx *app.UpdateSeriesContext) error {
	// SeriesController_Update: start_implement

	// Put your logic here

	return nil
	// SeriesController_Update: end_implement
}
