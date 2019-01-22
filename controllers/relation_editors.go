package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationEditorsController implements the relationEditors resource.
type RelationEditorsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationEditorsController creates a relationEditors controller.
func NewRelationEditorsController(service *goa.Service, fm Fmodeler, l Lister) *RelationEditorsController {
	return &RelationEditorsController{Controller: service.NewController("RelationEditorsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationEditorsController) ListBooks(ctx *app.ListBooksRelationEditorsContext) error {
	// RelationEditorsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, nil, nil)
	// RelationEditorsController_ListBooks: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationEditorsController) ListCollections(ctx *app.ListCollectionsRelationEditorsContext) error {
	// RelationEditorsController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, nil, nil)
	// RelationEditorsController_ListCollections: end_implement
}

// ListPrints runs the listPrints action.
func (c *RelationEditorsController) ListPrints(ctx *app.ListPrintsRelationEditorsContext) error {
	// RelationEditorsController_ListPrints: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, nil, &ctx.EditorID, nil)
	// RelationEditorsController_ListPrints: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationEditorsController) ListSeries(ctx *app.ListSeriesRelationEditorsContext) error {
	// RelationEditorsController_ListSeries: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, nil, &ctx.EditorID, nil)
	// RelationEditorsController_ListSeries: end_implement
}
