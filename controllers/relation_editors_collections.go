package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationEditorsCollectionsController implements the relationEditorsCollections resource.
type RelationEditorsCollectionsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationEditorsCollectionsController creates a relationEditorsCollections controller.
func NewRelationEditorsCollectionsController(service *goa.Service, fm Fmodeler, l Lister) *RelationEditorsCollectionsController {
	return &RelationEditorsCollectionsController{Controller: service.NewController("RelationEditorsCollectionsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationEditorsCollectionsController) ListBooks(ctx *app.ListBooksRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, nil, nil)
	// RelationEditorsCollectionsController_ListBooks: end_implement
}

// ListBooksByPrint runs the listBooksByPrint action.
func (c *RelationEditorsCollectionsController) ListBooksByPrint(ctx *app.ListBooksByPrintRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListBooksByPrint: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, nil)
	// RelationEditorsCollectionsController_ListBooksByPrint: end_implement
}

// ListBooksByPrintSeries runs the listBooksByPrintSeries action.
func (c *RelationEditorsCollectionsController) ListBooksByPrintSeries(ctx *app.ListBooksByPrintSeriesRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListBooksByPrintSeries: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsCollectionsController_ListBooksByPrintSeries: end_implement
}

// ListBooksBySeries runs the listBooksBySeries action.
func (c *RelationEditorsCollectionsController) ListBooksBySeries(ctx *app.ListBooksBySeriesRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListBooksBySeries: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, nil, &ctx.SeriesID)
	// RelationEditorsCollectionsController_ListBooksBySeries: end_implement
}

// ListBooksBySeriesPrint runs the listBooksBySeriesPrint action.
func (c *RelationEditorsCollectionsController) ListBooksBySeriesPrint(ctx *app.ListBooksBySeriesPrintRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListBooksBySeriesPrint: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsCollectionsController_ListBooksBySeriesPrint: end_implement
}

// ListPrints runs the listPrints action.
func (c *RelationEditorsCollectionsController) ListPrints(ctx *app.ListPrintsRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListPrints: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, nil)
	// RelationEditorsCollectionsController_ListPrints: end_implement
}

// ListPrintsBySeries runs the listPrintsBySeries action.
func (c *RelationEditorsCollectionsController) ListPrintsBySeries(ctx *app.ListPrintsBySeriesRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListPrintsBySeries: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.SeriesID)
	// RelationEditorsCollectionsController_ListPrintsBySeries: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationEditorsCollectionsController) ListSeries(ctx *app.ListSeriesRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListSeries: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, nil)
	// RelationEditorsCollectionsController_ListSeries: end_implement
}

// ListSeriesByPrint runs the listSeriesByPrint action.
func (c *RelationEditorsCollectionsController) ListSeriesByPrint(ctx *app.ListSeriesByPrintRelationEditorsCollectionsContext) error {
	// RelationEditorsCollectionsController_ListSeriesByPrint: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID)
	// RelationEditorsCollectionsController_ListSeriesByPrint: end_implement
}
