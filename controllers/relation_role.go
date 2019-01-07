package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationRoleController implements the relationRole resource.
type RelationRoleController struct {
	*goa.Controller
}

// NewRelationRoleController creates a relationRole controller.
func NewRelationRoleController(service *goa.Service) *RelationRoleController {
	return &RelationRoleController{Controller: service.NewController("RelationRoleController")}
}

// ListAuthors runs the listAuthors action.
func (c *RelationRoleController) ListAuthors(ctx *app.ListAuthorsRelationRoleContext) error {
	// RelationRoleController_ListAuthors: start_implement

	// Put your logic here

	res := app.AuthorCollection{}
	return ctx.OK(res)
	// RelationRoleController_ListAuthors: end_implement
}

// ListSeriesByAuthors runs the listSeriesByAuthors action.
func (c *RelationRoleController) ListSeriesByAuthors(ctx *app.ListSeriesByAuthorsRelationRoleContext) error {
	// RelationRoleController_ListSeriesByAuthors: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationRoleController_ListSeriesByAuthors: end_implement
}
