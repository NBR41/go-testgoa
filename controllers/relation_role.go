package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/goadesign/goa"
)

// RelationRoleController implements the relationRole resource.
type RelationRoleController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationRoleController creates a relationRole controller.
func NewRelationRoleController(service *goa.Service, fm Fmodeler, l Lister) *RelationRoleController {
	return &RelationRoleController{Controller: service.NewController("RelationRoleController"), fm: fm, l: l}
}

// ListAuthors runs the listAuthors action.
func (c *RelationRoleController) ListAuthors(ctx *app.ListAuthorsRelationRoleContext) error {
	// RelationRoleController_ListAuthors: start_implement
	return c.l.ListAuthors(ctx, c.fm, ctx, nil, &ctx.RoleID)
	// RelationRoleController_ListAuthors: end_implement
}

// ListSeriesByAuthor runs the listSeriesByAuthor action.
func (c *RelationRoleController) ListSeriesByAuthor(ctx *app.ListSeriesByAuthorRelationRoleContext) error {
	// RelationRoleController_ListSeriesByAuthor: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, &ctx.AuthorID, nil, nil, &ctx.RoleID)
	// RelationRoleController_ListSeriesByAuthor: end_implement
}
