package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationEditorsPrintsController implements the relationEditorsPrints resource.
type RelationEditorsPrintsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationEditorsPrintsController creates a relationEditorsPrints controller.
func NewRelationEditorsPrintsController(service *goa.Service, fm Fmodeler, l Lister) *RelationEditorsPrintsController {
	return &RelationEditorsPrintsController{Controller: service.NewController("RelationEditorsPrintsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationEditorsPrintsController) ListBooks(ctx *app.ListBooksRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID, nil)
	// RelationEditorsPrintsController_ListBooks: end_implement
}

// ListBooksByCollection runs the listBooksByCollection action.
func (c *RelationEditorsPrintsController) ListBooksByCollection(ctx *app.ListBooksByCollectionRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListBooksByCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, nil)
	// RelationEditorsPrintsController_ListBooksByCollection: end_implement
}

// ListBooksByCollectionSeries runs the listBooksByCollectionSeries action.
func (c *RelationEditorsPrintsController) ListBooksByCollectionSeries(ctx *app.ListBooksByCollectionSeriesRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListBooksByCollectionSeries: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsPrintsController_ListBooksByCollectionSeries: end_implement
}

// ListBooksBySeries runs the listBooksBySeries action.
func (c *RelationEditorsPrintsController) ListBooksBySeries(ctx *app.ListBooksBySeriesRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListBooksBySeries: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsPrintsController_ListBooksBySeries: end_implement
}

// ListBooksBySeriesCollection runs the listBooksBySeriesCollection action.
func (c *RelationEditorsPrintsController) ListBooksBySeriesCollection(ctx *app.ListBooksBySeriesCollectionRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListBooksBySeriesCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsPrintsController_ListBooksBySeriesCollection: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationEditorsPrintsController) ListCollections(ctx *app.ListCollectionsRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, &ctx.PrintID, nil)
	// RelationEditorsPrintsController_ListCollections: end_implement
}

// ListCollectionsBySeries runs the listCollectionsBySeries action.
func (c *RelationEditorsPrintsController) ListCollectionsBySeries(ctx *app.ListCollectionsBySeriesRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListCollectionsBySeries: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationEditorsPrintsController_ListCollectionsBySeries: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationEditorsPrintsController) ListSeries(ctx *app.ListSeriesRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListSeries: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID)
	// RelationEditorsPrintsController_ListSeries: end_implement
}

// ListSeriesByCollection runs the listSeriesByCollection action.
func (c *RelationEditorsPrintsController) ListSeriesByCollection(ctx *app.ListSeriesByCollectionRelationEditorsPrintsContext) error {
	// RelationEditorsPrintsController_ListSeriesByCollection: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID)
	// RelationEditorsPrintsController_ListSeriesByCollection: end_implement
}
