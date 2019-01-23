package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationSeriesController implements the relationSeries resource.
type RelationSeriesController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationSeriesController creates a relationSeries controller.
func NewRelationSeriesController(service *goa.Service, fm Fmodeler, l Lister) *RelationSeriesController {
	return &RelationSeriesController{Controller: service.NewController("RelationSeriesController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationSeriesController) ListBooks(ctx *app.ListBooksRelationSeriesContext) error {
	// RelationSeriesController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, nil, nil, &ctx.SeriesID)
	// RelationSeriesController_ListBooks: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationSeriesController) ListCollections(ctx *app.ListCollectionsRelationSeriesContext) error {
	// RelationSeriesController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, nil, nil, &ctx.SeriesID)
	// RelationSeriesController_ListCollections: end_implement
}

// ListEditors runs the listEditors action.
func (c *RelationSeriesController) ListEditors(ctx *app.ListEditorsRelationSeriesContext) error {
	// RelationSeriesController_ListEditors: start_implement
	return c.l.ListEditors(ctx, c.fm, ctx, nil, &ctx.SeriesID)
	// RelationSeriesController_ListEditors: end_implement
}

// ListPrints runs the listPrints action.
func (c *RelationSeriesController) ListPrints(ctx *app.ListPrintsRelationSeriesContext) error {
	// RelationSeriesController_ListPrints: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, nil, nil, &ctx.SeriesID)
	// RelationSeriesController_ListPrints: end_implement
}
