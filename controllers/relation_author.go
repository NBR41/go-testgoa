package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// RelationAuthorController implements the relationAuthor resource.
type RelationAuthorController struct {
	*goa.Controller
	fm Fmodeler
	l  Lister
}

// NewRelationAuthorController creates a relationAuthor controller.
func NewRelationAuthorController(service *goa.Service, fm Fmodeler, l Lister) *RelationAuthorController {
	return &RelationAuthorController{Controller: service.NewController("RelationAuthorController"), fm: fm, l: l}
}

// ListCategories runs the listCategories action.
func (c *RelationAuthorController) ListCategories(ctx *app.ListCategoriesRelationAuthorContext) error {
	// RelationAuthorController_ListCategories: start_implement
	return c.l.ListCategories(ctx, c.fm, ctx, &ctx.AuthorID, nil)
	// RelationAuthorController_ListCategories: end_implement
}

// ListClasses runs the listClasses action.
func (c *RelationAuthorController) ListClasses(ctx *app.ListClassesRelationAuthorContext) error {
	// RelationAuthorController_ListClasses: start_implement
	return c.l.ListClasses(ctx, c.fm, ctx, &ctx.AuthorID, nil, nil)
	// RelationAuthorController_ListClasses: end_implement
}

// ListRoles runs the listRoles action.
func (c *RelationAuthorController) ListRoles(ctx *app.ListRolesRelationAuthorContext) error {
	// RelationAuthorController_ListRoles: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	_, err = m.GetAuthorByID(ctx.AuthorID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get author`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	list, err := m.ListRolesByIDs(&ctx.AuthorID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get role list`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	bs := make(app.RoleCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToRoleMedia(bk)
	}
	return ctx.OK(bs)
	// RelationAuthorController_ListRoles: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationAuthorController) ListSeries(ctx *app.ListSeriesRelationAuthorContext) error {
	// RelationAuthorController_ListSeries: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, &ctx.AuthorID, nil, nil, nil)
	// RelationAuthorController_ListSeries: end_implement
}

// ListSeriesByCategory runs the listSeriesByCategory action.
func (c *RelationAuthorController) ListSeriesByCategory(ctx *app.ListSeriesByCategoryRelationAuthorContext) error {
	// RelationAuthorController_ListSeriesByCategory: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, &ctx.AuthorID, &ctx.CategoryID, nil, nil)
	// RelationAuthorController_ListSeriesByCategory: end_implement
}

// ListSeriesByClass runs the listSeriesByClass action.
func (c *RelationAuthorController) ListSeriesByClass(ctx *app.ListSeriesByClassRelationAuthorContext) error {
	// RelationAuthorController_ListSeriesByClass: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, &ctx.AuthorID, nil, &ctx.ClassID, nil)
	// RelationAuthorController_ListSeriesByClass: end_implement
}

// ListSeriesByRole runs the listSeriesByRole action.
func (c *RelationAuthorController) ListSeriesByRole(ctx *app.ListSeriesByRoleRelationAuthorContext) error {
	// RelationAuthorController_ListSeriesByRole: start_implement
	return c.l.ListSeries(ctx, c.fm, ctx, &ctx.AuthorID, nil, nil, &ctx.RoleID)
	// RelationAuthorController_ListSeriesByRole: end_implement
}
