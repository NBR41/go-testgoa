package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationSeriesPrintsController implements the relationSeriesPrints resource.
type RelationSeriesPrintsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationSeriesPrintsController creates a relationSeriesPrints controller.
func NewRelationSeriesPrintsController(service *goa.Service, fm Fmodeler, l Lister) *RelationSeriesPrintsController {
	return &RelationSeriesPrintsController{Controller: service.NewController("RelationSeriesPrintsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationSeriesPrintsController) ListBooks(ctx *app.ListBooksRelationSeriesPrintsContext) error {
	// RelationSeriesPrintsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, nil, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesPrintsController_ListBooks: end_implement
}

// ListBooksByEditor runs the listBooksByEditor action.
func (c *RelationSeriesPrintsController) ListBooksByEditor(ctx *app.ListBooksByEditorRelationSeriesPrintsContext) error {
	// RelationSeriesPrintsController_ListBooksByEditor: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, nil, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesPrintsController_ListBooksByEditor: end_implement
}

// ListBooksByEditorCollection runs the listBooksByEditorCollection action.
func (c *RelationSeriesPrintsController) ListBooksByEditorCollection(ctx *app.ListBooksByEditorCollectionRelationSeriesPrintsContext) error {
	// RelationSeriesPrintsController_ListBooksByEditorCollection: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesPrintsController_ListBooksByEditorCollection: end_implement
}

// ListCollections runs the listCollections action.
func (c *RelationSeriesPrintsController) ListCollections(ctx *app.ListCollectionsRelationSeriesPrintsContext) error {
	// RelationSeriesPrintsController_ListCollections: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, nil, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesPrintsController_ListCollections: end_implement
}

// ListCollectionsByEditor runs the listCollectionsByEditor action.
func (c *RelationSeriesPrintsController) ListCollectionsByEditor(ctx *app.ListCollectionsByEditorRelationSeriesPrintsContext) error {
	// RelationSeriesPrintsController_ListCollectionsByEditor: start_implement
	return c.l.ListCollections(ctx, c.fm, ctx, &ctx.EditorID, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesPrintsController_ListCollectionsByEditor: end_implement
}

// ListEditors runs the listEditors action.
func (c *RelationSeriesPrintsController) ListEditors(ctx *app.ListEditorsRelationSeriesPrintsContext) error {
	// RelationSeriesPrintsController_ListEditors: start_implement
	return c.l.ListEditors(ctx, c.fm, ctx, &ctx.PrintID, &ctx.SeriesID)
	// RelationSeriesPrintsController_ListEditors: end_implement
}
