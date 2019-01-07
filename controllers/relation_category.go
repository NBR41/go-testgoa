package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationCategoryController implements the relationCategory resource.
type RelationCategoryController struct {
	*goa.Controller
}

// NewRelationCategoryController creates a relationCategory controller.
func NewRelationCategoryController(service *goa.Service) *RelationCategoryController {
	return &RelationCategoryController{Controller: service.NewController("RelationCategoryController")}
}

// ListAuthors runs the listAuthors action.
func (c *RelationCategoryController) ListAuthors(ctx *app.ListAuthorsRelationCategoryContext) error {
	// RelationCategoryController_ListAuthors: start_implement

	// Put your logic here

	res := app.AuthorCollection{}
	return ctx.OK(res)
	// RelationCategoryController_ListAuthors: end_implement
}

// ListClasses runs the listClasses action.
func (c *RelationCategoryController) ListClasses(ctx *app.ListClassesRelationCategoryContext) error {
	// RelationCategoryController_ListClasses: start_implement

	// Put your logic here

	res := app.ClassCollection{}
	return ctx.OK(res)
	// RelationCategoryController_ListClasses: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationCategoryController) ListSeries(ctx *app.ListSeriesRelationCategoryContext) error {
	// RelationCategoryController_ListSeries: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationCategoryController_ListSeries: end_implement
}

// ListSeriesByClass runs the listSeriesByClass action.
func (c *RelationCategoryController) ListSeriesByClass(ctx *app.ListSeriesByClassRelationCategoryContext) error {
	// RelationCategoryController_ListSeriesByClass: start_implement

	// Put your logic here

	res := app.SeriesCollection{}
	return ctx.OK(res)
	// RelationCategoryController_ListSeriesByClass: end_implement
}
