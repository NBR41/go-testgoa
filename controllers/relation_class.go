package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// RelationClassController implements the relationClass resource.
type RelationClassController struct {
	*goa.Controller
	fm Fmodeler
}

// NewRelationClassController creates a relationClass controller.
func NewRelationClassController(service *goa.Service, fm Fmodeler) *RelationClassController {
	return &RelationClassController{Controller: service.NewController("RelationClassController"), fm: fm}
}

// ListCategories runs the listCategories action.
func (c *RelationClassController) ListCategories(ctx *app.ListCategoriesRelationClassContext) error {
	// RelationClassController_ListCategories: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	_, err = m.GetClassByID(ctx.ClassID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get class`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	list, err := m.ListCategoriesByClassID(ctx.ClassID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get category list`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	bs := make(app.CategoryCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToCategoryMedia(bk)
	}
	return ctx.OK(bs)
	// RelationClassController_ListCategories: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationClassController) ListSeries(ctx *app.ListSeriesRelationClassContext) error {
	// RelationClassController_ListSeries: start_implement
	return listSeries(ctx, c.fm, ctx, nil, nil, nil, &ctx.ClassID)
	// RelationClassController_ListSeries: end_implement
}

// ListSeriesByCategory runs the listSeriesByCategory action.
func (c *RelationClassController) ListSeriesByCategory(ctx *app.ListSeriesByCategoryRelationClassContext) error {
	// RelationClassController_ListSeriesByCategory: start_implement
	return listSeries(ctx, c.fm, ctx, nil, nil, &ctx.CategoryID, &ctx.ClassID)
	// RelationClassController_ListSeriesByCategory: end_implement
}
