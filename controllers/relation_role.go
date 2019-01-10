package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// RelationRoleController implements the relationRole resource.
type RelationRoleController struct {
	*goa.Controller
	fm Fmodeler
}

// NewRelationRoleController creates a relationRole controller.
func NewRelationRoleController(service *goa.Service, fm Fmodeler) *RelationRoleController {
	return &RelationRoleController{Controller: service.NewController("RelationRoleController"), fm: fm}
}

// ListAuthors runs the listAuthors action.
func (c *RelationRoleController) ListAuthors(ctx *app.ListAuthorsRelationRoleContext) error {
	// RelationRoleController_ListAuthors: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListAuthorsByRoleID(ctx.RoleID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get author list`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}
	bs := make(app.AuthorCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToAuthorMedia(bk)
	}
	return ctx.OK(bs)
	// RelationRoleController_ListAuthors: end_implement
}

// ListSeriesByAuthor runs the listSeriesByAuthor action.
func (c *RelationRoleController) ListSeriesByAuthor(ctx *app.ListSeriesByAuthorRelationRoleContext) error {
	// RelationRoleController_ListSeriesByAuthor: start_implement
	return listSeries(ctx, c.fm, ctx, &ctx.AuthorID, &ctx.RoleID, nil, nil)
	// RelationRoleController_ListSeriesByAuthor: end_implement
}
