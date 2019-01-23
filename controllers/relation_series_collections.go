package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationSeriesCollectionsController implements the relationSeriesCollections resource.
type RelationSeriesCollectionsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationSeriesCollectionsController creates a relationSeriesCollections controller.
func NewRelationSeriesCollectionsController(service *goa.Service, fm Fmodeler, l Lister) *RelationSeriesCollectionsController {
	return &RelationSeriesCollectionsController{Controller: service.NewController("RelationSeriesCollectionsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationSeriesCollectionsController) ListBooks(ctx *app.ListBooksRelationSeriesCollectionsContext) error {
	// RelationSeriesCollectionsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, nil, nil, &ctx.SeriesID)
	// RelationSeriesCollectionsController_ListBooks: end_implement
}

// ListBooksByPrint runs the listBooksByPrint action.
func (c *RelationSeriesCollectionsController) ListBooksByPrint(ctx *app.ListBooksByPrintRelationSeriesCollectionsContext) error {
	// RelationSeriesCollectionsController_ListBooksByPrint: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, nil, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesCollectionsController_ListBooksByPrint: end_implement
}

// ListPrints runs the listPrints action.
func (c *RelationSeriesCollectionsController) ListPrints(ctx *app.ListPrintsRelationSeriesCollectionsContext) error {
	// RelationSeriesCollectionsController_ListPrints: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, &ctx.CollectionID, nil, &ctx.SeriesID)
	// RelationSeriesCollectionsController_ListPrints: end_implement
}
