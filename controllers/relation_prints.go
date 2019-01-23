package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationPrintsController implements the relationPrints resource.
type RelationPrintsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationPrintsController creates a relationPrints controller.
func NewRelationPrintsController(service *goa.Service, fm Fmodeler, l Lister) *RelationPrintsController {
	return &RelationPrintsController{Controller: service.NewController("RelationPrintsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationPrintsController) ListBooks(ctx *app.ListBooksRelationPrintsContext) error {
	// RelationPrintsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, nil, &ctx.PrintID, nil)
	// RelationPrintsController_ListBooks: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationPrintsController) ListCollections(ctx *app.ListCollectionsRelationPrintsContext) error {
	// RelationPrintsController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, nil, &ctx.PrintID, nil)
	// RelationPrintsController_ListCollections: end_implement
}

// ListEditors runs the listEditors action.
func (c *RelationPrintsController) ListEditors(ctx *app.ListEditorsRelationPrintsContext) error {
	// RelationPrintsController_ListEditors: start_implement
	return c.l.ListEditors(ctx, c.fm, ctx, &ctx.PrintID, nil)
	// RelationPrintsController_ListEditors: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationPrintsController) ListSeries(ctx *app.ListSeriesRelationPrintsContext) error {
	// RelationPrintsController_ListSeries: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, nil, nil, &ctx.PrintID)
	// RelationPrintsController_ListSeries: end_implement
}
