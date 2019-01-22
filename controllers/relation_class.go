package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationClassController implements the relationClass resource.
type RelationClassController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationClassController creates a relationClass controller.
func NewRelationClassController(service *goa.Service, fm Fmodeler, l Lister) *RelationClassController {
	return &RelationClassController{Controller: service.NewController("RelationClassController"), fm: fm, l: l}
}

// ListCategories runs the listCategories action.
func (c *RelationClassController) ListCategories(ctx *app.ListCategoriesRelationClassContext) error {
	// RelationClassController_ListCategories: start_implement
	return c.l.ListCategories(ctx, c.fm, ctx, nil, &ctx.ClassID)
	// RelationClassController_ListCategories: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationClassController) ListSeries(ctx *app.ListSeriesRelationClassContext) error {
	// RelationClassController_ListSeries: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, nil, nil, &ctx.ClassID, nil)
	// RelationClassController_ListSeries: end_implement
}

// ListSeriesByCategory runs the listSeriesByCategory action.
func (c *RelationClassController) ListSeriesByCategory(ctx *app.ListSeriesByCategoryRelationClassContext) error {
	// RelationClassController_ListSeriesByCategory: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, nil, &ctx.CategoryID, &ctx.ClassID, nil)
	// RelationClassController_ListSeriesByCategory: end_implement
}
