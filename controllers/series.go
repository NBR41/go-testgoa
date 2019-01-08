package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// SeriesController implements the series resource.
type SeriesController struct {
	*goa.Controller
	fm Fmodeler
}

// NewSeriesController creates a series controller.
func NewSeriesController(service *goa.Service, fm Fmodeler) *SeriesController {
	return &SeriesController{
		Controller: service.NewController("SeriesController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *SeriesController) Create(ctx *app.CreateSeriesContext) error {
	// SeriesController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	v, err := m.InsertSeries(ctx.Payload.SeriesName, ctx.Payload.CategoryID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert series`, `error`, err.Error())
		if err == model.ErrDuplicateKey || err == model.ErrNotFound {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.SeriesHref(v.ID))
	return ctx.Created()
	// SeriesController_Create: end_implement
}

// Delete runs the delete action.
func (c *SeriesController) Delete(ctx *app.DeleteSeriesContext) error {
	// SeriesController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteSeries(ctx.SeriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete series`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// SeriesController_Delete: end_implement
}

// List runs the list action.
func (c *SeriesController) List(ctx *app.ListSeriesContext) error {
	// SeriesController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListSeries()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get series list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.SeriesCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToSeriesMedia(bk)
	}
	return ctx.OK(bs)
	// SeriesController_List: end_implement
}

// Show runs the show action.
func (c *SeriesController) Show(ctx *app.ShowSeriesContext) error {
	// SeriesController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetSeriesByID(ctx.SeriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get series`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToSeriesMedia(b))
	// SeriesController_Show: end_implement
}

// Update runs the update action.
func (c *SeriesController) Update(ctx *app.UpdateSeriesContext) error {
	// SeriesController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateSeries(ctx.SeriesID, ctx.Payload.SeriesName, ctx.Payload.CategoryID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update series`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// SeriesController_Update: end_implement
}
