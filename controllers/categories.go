package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// CategoriesController implements the categories resource.
type CategoriesController struct {
	*goa.Controller
}

// NewCategoriesController creates a categories controller.
func NewCategoriesController(service *goa.Service) *CategoriesController {
	return &CategoriesController{Controller: service.NewController("CategoriesController")}
}

// Create runs the create action.
func (c *CategoriesController) Create(ctx *app.CreateCategoriesContext) error {
	// CategoriesController_Create: start_implement

	// Put your logic here

	return nil
	// CategoriesController_Create: end_implement
}

// Delete runs the delete action.
func (c *CategoriesController) Delete(ctx *app.DeleteCategoriesContext) error {
	// CategoriesController_Delete: start_implement

	// Put your logic here

	return nil
	// CategoriesController_Delete: end_implement
}

// List runs the list action.
func (c *CategoriesController) List(ctx *app.ListCategoriesContext) error {
	// CategoriesController_List: start_implement

	// Put your logic here

	res := app.CategoryCollection{}
	return ctx.OK(res)
	// CategoriesController_List: end_implement
}

// Show runs the show action.
func (c *CategoriesController) Show(ctx *app.ShowCategoriesContext) error {
	// CategoriesController_Show: start_implement

	// Put your logic here

	res := &app.Category{}
	return ctx.OK(res)
	// CategoriesController_Show: end_implement
}

// Update runs the update action.
func (c *CategoriesController) Update(ctx *app.UpdateCategoriesContext) error {
	// CategoriesController_Update: start_implement

	// Put your logic here

	return nil
	// CategoriesController_Update: end_implement
}
