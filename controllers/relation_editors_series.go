package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationEditorsSeriesController implements the relationEditorsSeries resource.
type RelationEditorsSeriesController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationEditorsSeriesController creates a relationEditorsSeries controller.
func NewRelationEditorsSeriesController(service *goa.Service, fm Fmodeler, l Lister) *RelationEditorsSeriesController {
	return &RelationEditorsSeriesController{Controller: service.NewController("RelationEditorsSeriesController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationEditorsSeriesController) ListBooks(ctx *app.ListBooksRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, nil, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListBooks: end_implement
}

// ListBooksByCollection runs the listBooksByCollection action.
func (c *RelationEditorsSeriesController) ListBooksByCollection(ctx *app.ListBooksByCollectionRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListBooksByCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, nil, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListBooksByCollection: end_implement
}

// ListBooksByCollectionPrint runs the listBooksByCollectionPrint action.
func (c *RelationEditorsSeriesController) ListBooksByCollectionPrint(ctx *app.ListBooksByCollectionPrintRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListBooksByCollectionPrint: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListBooksByCollectionPrint: end_implement
}

// ListBooksByPrint runs the listBooksByPrint action.
func (c *RelationEditorsSeriesController) ListBooksByPrint(ctx *app.ListBooksByPrintRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListBooksByPrint: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListBooksByPrint: end_implement
}

// ListBooksByPrintCollection runs the listBooksByPrintCollection action.
func (c *RelationEditorsSeriesController) ListBooksByPrintCollection(ctx *app.ListBooksByPrintCollectionRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListBooksByPrintCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListBooksByPrintCollection: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationEditorsSeriesController) ListCollections(ctx *app.ListCollectionsRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, nil, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListCollections: end_implement
}

// ListCollectionsByPrint runs the listCollectionsByPrint action.
func (c *RelationEditorsSeriesController) ListCollectionsByPrint(ctx *app.ListCollectionsByPrintRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListCollectionsByPrint: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListCollectionsByPrint: end_implement
}

// ListPrints runs the listPrints action.
func (c *RelationEditorsSeriesController) ListPrints(ctx *app.ListPrintsRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListPrints: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListPrints: end_implement
}

// ListPrintsByCollection runs the listPrintsByCollection action.
func (c *RelationEditorsSeriesController) ListPrintsByCollection(ctx *app.ListPrintsByCollectionRelationEditorsSeriesContext) error {
	// RelationEditorsSeriesController_ListPrintsByCollection: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.SeriesID)
	// RelationEditorsSeriesController_ListPrintsByCollection: end_implement
}
