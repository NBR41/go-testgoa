package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationCategoryController implements the relationCategory resource.
type RelationCategoryController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationCategoryController creates a relationCategory controller.
func NewRelationCategoryController(service *goa.Service, fm Fmodeler, l Lister) *RelationCategoryController {
	return &RelationCategoryController{Controller: service.NewController("RelationCategoryController"), fm: fm, l: l}
}

// ListAuthors runs the listAuthors action.
func (c *RelationCategoryController) ListAuthors(ctx *app.ListAuthorsRelationCategoryContext) error {
	// RelationCategoryController_ListAuthors: start_implement
	return c.l.ListAuthors(ctx, c.fm, ctx, &ctx.CategoryID, nil)
	// RelationCategoryController_ListAuthors: end_implement
}

// ListClasses runs the listClasses action.
func (c *RelationCategoryController) ListClasses(ctx *app.ListClassesRelationCategoryContext) error {
	// RelationCategoryController_ListClasses: start_implement
	return c.l.ListClasses(ctx, c.fm, ctx, nil, &ctx.CategoryID, nil)
	// RelationCategoryController_ListClasses: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationCategoryController) ListSeries(ctx *app.ListSeriesRelationCategoryContext) error {
	// RelationCategoryController_ListSeries: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, nil, &ctx.CategoryID, nil, nil)
	// RelationCategoryController_ListSeries: end_implement
}

// ListSeriesByClass runs the listSeriesByClass action.
func (c *RelationCategoryController) ListSeriesByClass(ctx *app.ListSeriesByClassRelationCategoryContext) error {
	// RelationCategoryController_ListSeriesByClass: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, nil, &ctx.CategoryID, &ctx.ClassID, nil)
	// RelationCategoryController_ListSeriesByClass: end_implement
}
