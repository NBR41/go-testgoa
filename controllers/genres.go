package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// GenresController implements the genres resource.
type GenresController struct {
	*goa.Controller
}

// NewGenresController creates a genres controller.
func NewGenresController(service *goa.Service) *GenresController {
	return &GenresController{Controller: service.NewController("GenresController")}
}

// Create runs the create action.
func (c *GenresController) Create(ctx *app.CreateGenresContext) error {
	// GenresController_Create: start_implement

	// Put your logic here

	return nil
	// GenresController_Create: end_implement
}

// Delete runs the delete action.
func (c *GenresController) Delete(ctx *app.DeleteGenresContext) error {
	// GenresController_Delete: start_implement

	// Put your logic here

	return nil
	// GenresController_Delete: end_implement
}

// List runs the list action.
func (c *GenresController) List(ctx *app.ListGenresContext) error {
	// GenresController_List: start_implement

	// Put your logic here

	res := app.GenreCollection{}
	return ctx.OK(res)
	// GenresController_List: end_implement
}

// Show runs the show action.
func (c *GenresController) Show(ctx *app.ShowGenresContext) error {
	// GenresController_Show: start_implement

	// Put your logic here

	res := &app.Genre{}
	return ctx.OK(res)
	// GenresController_Show: end_implement
}

// Update runs the update action.
func (c *GenresController) Update(ctx *app.UpdateGenresContext) error {
	// GenresController_Update: start_implement

	// Put your logic here

	return nil
	// GenresController_Update: end_implement
}
