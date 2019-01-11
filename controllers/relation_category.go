package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// RelationCategoryController implements the relationCategory resource.
type RelationCategoryController struct {
	*goa.Controller
	fm Fmodeler
}

// NewRelationCategoryController creates a relationCategory controller.
func NewRelationCategoryController(service *goa.Service, fm Fmodeler) *RelationCategoryController {
	return &RelationCategoryController{Controller: service.NewController("RelationCategoryController"), fm: fm}
}

// ListAuthors runs the listAuthors action.
func (c *RelationCategoryController) ListAuthors(ctx *app.ListAuthorsRelationCategoryContext) error {
	// RelationCategoryController_ListAuthors: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	_, err = m.GetCategoryByID(ctx.CategoryID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get category`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	list, err := m.ListAuthorsByCategoryID(ctx.CategoryID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get author list`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	bs := make(app.AuthorCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToAuthorMedia(bk)
	}
	return ctx.OK(bs)
	// RelationCategoryController_ListAuthors: end_implement
}

// ListClasses runs the listClasses action.
func (c *RelationCategoryController) ListClasses(ctx *app.ListClassesRelationCategoryContext) error {
	// RelationCategoryController_ListClasses: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	_, err = m.GetCategoryByID(ctx.CategoryID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get category`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	list, err := m.ListClassesByCategoryID(ctx.CategoryID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get class list`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	bs := make(app.ClassCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToClassMedia(bk)
	}
	return ctx.OK(bs)
	// RelationCategoryController_ListClasses: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationCategoryController) ListSeries(ctx *app.ListSeriesRelationCategoryContext) error {
	// RelationCategoryController_ListSeries: start_implement
	return listSeries(ctx, c.fm, ctx, nil, nil, &ctx.CategoryID, nil)
	// RelationCategoryController_ListSeries: end_implement
}

// ListSeriesByClass runs the listSeriesByClass action.
func (c *RelationCategoryController) ListSeriesByClass(ctx *app.ListSeriesByClassRelationCategoryContext) error {
	// RelationCategoryController_ListSeriesByClass: start_implement
	return listSeries(ctx, c.fm, ctx, nil, nil, &ctx.CategoryID, &ctx.ClassID)
	// RelationCategoryController_ListSeriesByClass: end_implement
}
