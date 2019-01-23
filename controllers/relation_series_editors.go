package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationSeriesEditorsController implements the relationSeriesEditors resource.
type RelationSeriesEditorsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationSeriesEditorsController creates a relationSeriesEditors controller.
func NewRelationSeriesEditorsController(service *goa.Service, fm Fmodeler, l Lister) *RelationSeriesEditorsController {
	return &RelationSeriesEditorsController{Controller: service.NewController("RelationSeriesEditorsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationSeriesEditorsController) ListBooks(ctx *app.ListBooksRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, nil, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListBooks: end_implement
}

// ListBooksByCollection runs the listBooksByCollection action.
func (c *RelationSeriesEditorsController) ListBooksByCollection(ctx *app.ListBooksByCollectionRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListBooksByCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, nil, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListBooksByCollection: end_implement
}

// ListBooksByCollectionPrint runs the listBooksByCollectionPrint action.
func (c *RelationSeriesEditorsController) ListBooksByCollectionPrint(ctx *app.ListBooksByCollectionPrintRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListBooksByCollectionPrint: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListBooksByCollectionPrint: end_implement
}

// ListBooksByPrint runs the listBooksByPrint action.
func (c *RelationSeriesEditorsController) ListBooksByPrint(ctx *app.ListBooksByPrintRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListBooksByPrint: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListBooksByPrint: end_implement
}

// ListBooksByPrintCollection runs the listBooksByPrintCollection action.
func (c *RelationSeriesEditorsController) ListBooksByPrintCollection(ctx *app.ListBooksByPrintCollectionRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListBooksByPrintCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListBooksByPrintCollection: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationSeriesEditorsController) ListCollections(ctx *app.ListCollectionsRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, nil, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListCollections: end_implement
}

// ListCollectionsByPrint runs the listCollectionsByPrint action.
func (c *RelationSeriesEditorsController) ListCollectionsByPrint(ctx *app.ListCollectionsByPrintRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListCollectionsByPrint: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListCollectionsByPrint: end_implement
}

// ListPrints runs the listPrints action.
func (c *RelationSeriesEditorsController) ListPrints(ctx *app.ListPrintsRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListPrints: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListPrints: end_implement
}

// ListPrintsByCollection runs the listPrintsByCollection action.
func (c *RelationSeriesEditorsController) ListPrintsByCollection(ctx *app.ListPrintsByCollectionRelationSeriesEditorsContext) error {
	// RelationSeriesEditorsController_ListPrintsByCollection: start_implement
	return c.l.ListPrints(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.SeriesID)
	// RelationSeriesEditorsController_ListPrintsByCollection: end_implement
}
