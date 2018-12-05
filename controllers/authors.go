package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// AuthorsController implements the authors resource.
type AuthorsController struct {
	*goa.Controller
}

// NewAuthorsController creates a authors controller.
func NewAuthorsController(service *goa.Service) *AuthorsController {
	return &AuthorsController{Controller: service.NewController("AuthorsController")}
}

// Create runs the create action.
func (c *AuthorsController) Create(ctx *app.CreateAuthorsContext) error {
	// AuthorsController_Create: start_implement

	// Put your logic here

	return nil
	// AuthorsController_Create: end_implement
}

// Delete runs the delete action.
func (c *AuthorsController) Delete(ctx *app.DeleteAuthorsContext) error {
	// AuthorsController_Delete: start_implement

	// Put your logic here

	return nil
	// AuthorsController_Delete: end_implement
}

// List runs the list action.
func (c *AuthorsController) List(ctx *app.ListAuthorsContext) error {
	// AuthorsController_List: start_implement

	// Put your logic here

	res := app.AuthorCollection{}
	return ctx.OK(res)
	// AuthorsController_List: end_implement
}

// Show runs the show action.
func (c *AuthorsController) Show(ctx *app.ShowAuthorsContext) error {
	// AuthorsController_Show: start_implement

	// Put your logic here

	res := &app.Author{}
	return ctx.OK(res)
	// AuthorsController_Show: end_implement
}

// Update runs the update action.
func (c *AuthorsController) Update(ctx *app.UpdateAuthorsContext) error {
	// AuthorsController_Update: start_implement

	// Put your logic here

	return nil
	// AuthorsController_Update: end_implement
}
