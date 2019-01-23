package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationPrintsEditorsController implements the relationPrintsEditors resource.
type RelationPrintsEditorsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationPrintsEditorsController creates a relationPrintsEditors controller.
func NewRelationPrintsEditorsController(service *goa.Service, fm Fmodeler, l Lister) *RelationPrintsEditorsController {
	return &RelationPrintsEditorsController{Controller: service.NewController("RelationPrintsEditorsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationPrintsEditorsController) ListBooks(ctx *app.ListBooksRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID, nil)
	// RelationPrintsEditorsController_ListBooks: end_implement
}

// ListBooksByCollection runs the listBooksByCollection action.
func (c *RelationPrintsEditorsController) ListBooksByCollection(ctx *app.ListBooksByCollectionRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListBooksByCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, nil)
	// RelationPrintsEditorsController_ListBooksByCollection: end_implement
}

// ListBooksByCollectionSeries runs the listBooksByCollectionSeries action.
func (c *RelationPrintsEditorsController) ListBooksByCollectionSeries(ctx *app.ListBooksByCollectionSeriesRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListBooksByCollectionSeries: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsEditorsController_ListBooksByCollectionSeries: end_implement
}

// ListBooksBySeries runs the listBooksBySeries action.
func (c *RelationPrintsEditorsController) ListBooksBySeries(ctx *app.ListBooksBySeriesRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListBooksBySeries: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsEditorsController_ListBooksBySeries: end_implement
}

// ListBooksBySeriesCollection runs the listBooksBySeriesCollection action.
func (c *RelationPrintsEditorsController) ListBooksBySeriesCollection(ctx *app.ListBooksBySeriesCollectionRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListBooksBySeriesCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsEditorsController_ListBooksBySeriesCollection: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationPrintsEditorsController) ListCollections(ctx *app.ListCollectionsRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, &ctx.PrintID, nil)
	// RelationPrintsEditorsController_ListCollections: end_implement
}

// ListCollectionsBySeries runs the listCollectionsBySeries action.
func (c *RelationPrintsEditorsController) ListCollectionsBySeries(ctx *app.ListCollectionsBySeriesRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListCollectionsBySeries: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsEditorsController_ListCollectionsBySeries: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationPrintsEditorsController) ListSeries(ctx *app.ListSeriesRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListSeries: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID)
	// RelationPrintsEditorsController_ListSeries: end_implement
}

// ListSeriesByCollection runs the listSeriesByCollection action.
func (c *RelationPrintsEditorsController) ListSeriesByCollection(ctx *app.ListSeriesByCollectionRelationPrintsEditorsContext) error {
	// RelationPrintsEditorsController_ListSeriesByCollection: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID)
	// RelationPrintsEditorsController_ListSeriesByCollection: end_implement
}
