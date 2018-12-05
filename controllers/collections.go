package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// CollectionsController implements the collections resource.
type CollectionsController struct {
	*goa.Controller
}

// NewCollectionsController creates a collections controller.
func NewCollectionsController(service *goa.Service) *CollectionsController {
	return &CollectionsController{Controller: service.NewController("CollectionsController")}
}

// Create runs the create action.
func (c *CollectionsController) Create(ctx *app.CreateCollectionsContext) error {
	// CollectionsController_Create: start_implement

	// Put your logic here

	return nil
	// CollectionsController_Create: end_implement
}

// Delete runs the delete action.
func (c *CollectionsController) Delete(ctx *app.DeleteCollectionsContext) error {
	// CollectionsController_Delete: start_implement

	// Put your logic here

	return nil
	// CollectionsController_Delete: end_implement
}

// List runs the list action.
func (c *CollectionsController) List(ctx *app.ListCollectionsContext) error {
	// CollectionsController_List: start_implement

	// Put your logic here

	res := app.CollectionCollection{}
	return ctx.OK(res)
	// CollectionsController_List: end_implement
}

// Show runs the show action.
func (c *CollectionsController) Show(ctx *app.ShowCollectionsContext) error {
	// CollectionsController_Show: start_implement

	// Put your logic here

	res := &app.Collection{}
	return ctx.OK(res)
	// CollectionsController_Show: end_implement
}

// Update runs the update action.
func (c *CollectionsController) Update(ctx *app.UpdateCollectionsContext) error {
	// CollectionsController_Update: start_implement

	// Put your logic here

	return nil
	// CollectionsController_Update: end_implement
}
