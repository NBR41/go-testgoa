package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationPrintsSeriesController implements the relationPrintsSeries resource.
type RelationPrintsSeriesController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationPrintsSeriesController creates a relationPrintsSeries controller.
func NewRelationPrintsSeriesController(service *goa.Service, fm Fmodeler, l Lister) *RelationPrintsSeriesController {
	return &RelationPrintsSeriesController{Controller: service.NewController("RelationPrintsSeriesController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationPrintsSeriesController) ListBooks(ctx *app.ListBooksRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, nil, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListBooks: end_implement
}

// ListBooksByCollection runs the listBooksByCollection action.
func (c *RelationPrintsSeriesController) ListBooksByCollection(ctx *app.ListBooksByCollectionRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListBooksByCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, nil, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListBooksByCollection: end_implement
}

// ListBooksByCollectionEditor runs the listBooksByCollectionEditor action.
func (c *RelationPrintsSeriesController) ListBooksByCollectionEditor(ctx *app.ListBooksByCollectionEditorRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListBooksByCollectionEditor: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListBooksByCollectionEditor: end_implement
}

// ListBooksByEditor runs the listBooksByEditor action.
func (c *RelationPrintsSeriesController) ListBooksByEditor(ctx *app.ListBooksByEditorRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListBooksByEditor: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListBooksByEditor: end_implement
}

// ListBooksByEditorCollection runs the listBooksByEditorCollection action.
func (c *RelationPrintsSeriesController) ListBooksByEditorCollection(ctx *app.ListBooksByEditorCollectionRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListBooksByEditorCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListBooksByEditorCollection: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationPrintsSeriesController) ListCollections(ctx *app.ListCollectionsRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, nil, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListCollections: end_implement
}

// ListCollectionsByEditor runs the listCollectionsByEditor action.
func (c *RelationPrintsSeriesController) ListCollectionsByEditor(ctx *app.ListCollectionsByEditorRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListCollectionsByEditor: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListCollectionsByEditor: end_implement
}

// ListEditors runs the listEditors action.
func (c *RelationPrintsSeriesController) ListEditors(ctx *app.ListEditorsRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListEditors: start_implement
	return c.l.ListEditors(ctx, c.fm, ctx, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListEditors: end_implement
}

// ListEditorsByCollection runs the listEditorsByCollection action.
func (c *RelationPrintsSeriesController) ListEditorsByCollection(ctx *app.ListEditorsByCollectionRelationPrintsSeriesContext) error {
	// RelationPrintsSeriesController_ListEditorsByCollection: start_implement
	return c.l.ListEditors(ctx, c.fm, ctx, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsSeriesController_ListEditorsByCollection: end_implement
}
