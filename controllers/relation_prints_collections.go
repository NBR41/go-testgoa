package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationPrintsCollectionsController implements the relationPrintsCollections resource.
type RelationPrintsCollectionsController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationPrintsCollectionsController creates a relationPrintsCollections controller.
func NewRelationPrintsCollectionsController(service *goa.Service, fm Fmodeler, l Lister) *RelationPrintsCollectionsController {
	return &RelationPrintsCollectionsController{Controller: service.NewController("RelationPrintsCollectionsController"), fm: fm, l: l}
}

// ListBooks runs the listBooks action.
func (c *RelationPrintsCollectionsController) ListBooks(ctx *app.ListBooksRelationPrintsCollectionsContext) error {
	// RelationPrintsCollectionsController_ListBooks: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, nil, &ctx.PrintID, nil)
	// RelationPrintsCollectionsController_ListBooks: end_implement
}

// ListBooksBySeries runs the listBooksBySeries action.
func (c *RelationPrintsCollectionsController) ListBooksBySeries(ctx *app.ListBooksBySeriesRelationPrintsCollectionsContext) error {
	// RelationPrintsCollectionsController_ListBooksBySeries: start_implement
	return c.l.ListBooks(ctx, c.fm, ctx, &ctx.CollectionID, nil, &ctx.PrintID, &ctx.SeriesID)
	// RelationPrintsCollectionsController_ListBooksBySeries: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationPrintsCollectionsController) ListSeries(ctx *app.ListSeriesRelationPrintsCollectionsContext) error {
	// RelationPrintsCollectionsController_ListSeries: start_implement
	return c.l.ListSeriesByEditionIDs(ctx, c.fm, ctx, &ctx.CollectionID, nil, &ctx.PrintID)
	// RelationPrintsCollectionsController_ListSeries: end_implement
}
