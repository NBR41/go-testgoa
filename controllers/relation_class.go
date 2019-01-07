package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationClassController implements the relationClass resource.
type RelationClassController struct {
	*goa.Controller
}

// NewRelationClassController creates a relationClass controller.
func NewRelationClassController(service *goa.Service) *RelationClassController {
	return &RelationClassController{Controller: service.NewController("RelationClassController")}
}

// ListCategories runs the listCategories action.
func (c *RelationClassController) ListCategories(ctx *app.ListCategoriesRelationClassContext) error {
	// RelationClassController_ListCategories: start_implement

	// Put your logic here

	res := app.CategoryCollection{}
	return ctx.OK(res)
	// RelationClassController_ListCategories: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationClassController) ListSeries(ctx *app.ListSeriesRelationClassContext) error {
	// RelationClassController_ListSeries: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationClassController_ListSeries: end_implement
}

// ListSeriesByCategory runs the listSeriesByCategory action.
func (c *RelationClassController) ListSeriesByCategory(ctx *app.ListSeriesByCategoryRelationClassContext) error {
	// RelationClassController_ListSeriesByCategory: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationClassController_ListSeriesByCategory: end_implement
}
