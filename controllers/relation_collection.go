package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// RelationCollectionController implements the relationCollection resource.
type RelationCollectionController struct {
	*goa.Controller
	fm Fmodeler
}

// NewRelationCollectionController creates a relationCollection controller.
func NewRelationCollectionController(service *goa.Service, fm Fmodeler) *RelationCollectionController {
	return &RelationCollectionController{Controller: service.NewController("RelationCollectionController"), fm: fm}
}

// ListBooks runs the listBooks action.
func (c *RelationCollectionController) ListBooks(ctx *app.ListBooksRelationCollectionContext) error {
	// RelationCollectionController_ListBooks: start_implement
	return listBooks(ctx, c.fm, ctx, &ctx.CollectionID, nil, nil)
	// RelationCollectionController_ListBooks: end_implement
}

// ListBooksByPrint runs the listBooksByPrint action.
func (c *RelationCollectionController) ListBooksByPrint(ctx *app.ListBooksByPrintRelationCollectionContext) error {
	// RelationCollectionController_ListBooksByPrint: start_implement
	return listBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.PrintID, nil)
	// RelationCollectionController_ListBooksByPrint: end_implement
}

// ListBooksByPrintSeries runs the listBooksByPrintSeries action.
func (c *RelationCollectionController) ListBooksByPrintSeries(ctx *app.ListBooksByPrintSeriesRelationCollectionContext) error {
	// RelationCollectionController_ListBooksByPrintSeries: start_implement
	return listBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.PrintID, &ctx.SeriesID)
	// RelationCollectionController_ListBooksByPrintSeries: end_implement
}

// ListBooksBySeries runs the listBooksBySeries action.
func (c *RelationCollectionController) ListBooksBySeries(ctx *app.ListBooksBySeriesRelationCollectionContext) error {
	// RelationCollectionController_ListBooksBySeries: start_implement
	return listBooks(ctx, c.fm, ctx, &ctx.CollectionID, nil, &ctx.SeriesID)
	// RelationCollectionController_ListBooksBySeries: end_implement
}

// ListBooksBySeriesPrint runs the listBooksBySeriesPrint action.
func (c *RelationCollectionController) ListBooksBySeriesPrint(ctx *app.ListBooksBySeriesPrintRelationCollectionContext) error {
	// RelationCollectionController_ListBooksBySeriesPrint: start_implement
	return listBooks(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.PrintID, &ctx.SeriesID)
	// RelationCollectionController_ListBooksBySeriesPrint: end_implement
}

// ListPrints runs the listPrints action.
func (c *RelationCollectionController) ListPrints(ctx *app.ListPrintsRelationCollectionContext) error {
	// RelationCollectionController_ListPrints: start_implement
	return listPrints(ctx, c.fm, ctx, &ctx.CollectionID, nil)
	// RelationCollectionController_ListPrints: end_implement
}

// ListPrintsBySeries runs the listPrintsBySeries action.
func (c *RelationCollectionController) ListPrintsBySeries(ctx *app.ListPrintsBySeriesRelationCollectionContext) error {
	// RelationCollectionController_ListPrintsBySeries: start_implement
	return listPrints(ctx, c.fm, ctx, &ctx.CollectionID, &ctx.SeriesID)
	// RelationCollectionController_ListPrintsBySeries: end_implement
}

// ListSeries runs the listSeries action.
func (c *RelationCollectionController) ListSeries(ctx *app.ListSeriesRelationCollectionContext) error {
	// RelationCollectionController_ListSeries: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	_, err = m.GetCollectionByID(ctx.CollectionID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get collection`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	list, err := m.ListSeriesByCollectionID(ctx.CollectionID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get series list`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	bs := make(app.SeriesCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToSeriesMedia(bk)
	}
	return ctx.OK(bs)
	// RelationCollectionController_ListSeries: end_implement
}

// ListSeriesByPrint runs the listSeriesByPrint action.
func (c *RelationCollectionController) ListSeriesByPrint(ctx *app.ListSeriesByPrintRelationCollectionContext) error {
	// RelationCollectionController_ListSeriesByPrint: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	_, err = m.GetCollectionByID(ctx.CollectionID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get collection`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	_, err = m.GetPrintByID(ctx.PrintID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get print`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	list, err := m.ListSeriesByCollectionIDPrintID(ctx.CollectionID, ctx.PrintID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get series list`, `error`, err.Error())
		return ctx.InternalServerError()
	}
	bs := make(app.SeriesCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToSeriesMedia(bk)
	}
	return ctx.OK(bs)
	// RelationCollectionController_ListSeriesByPrint: end_implement
}
