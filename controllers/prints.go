package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// PrintsController implements the prints resource.
type PrintsController struct {
	*goa.Controller
	fm Fmodeler
}

// NewPrintsController creates a prints controller.
func NewPrintsController(service *goa.Service, fm Fmodeler) *PrintsController {
	return &PrintsController{
		Controller: service.NewController("PrintsController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *PrintsController) Create(ctx *app.CreatePrintsContext) error {
	// PrintsController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertPrint(ctx.Payload.PrintName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert print`, `error`, err.Error())
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.PrintsHref(b.ID))
	return ctx.Created()
	// PrintsController_Create: end_implement
}

// Delete runs the delete action.
func (c *PrintsController) Delete(ctx *app.DeletePrintsContext) error {
	// PrintsController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeletePrint(ctx.PrintID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete print`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// PrintsController_Delete: end_implement
}

// List runs the list action.
func (c *PrintsController) List(ctx *app.ListPrintsContext) error {
	// PrintsController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListPrintsByIDs(nil, nil, nil)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get print list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.PrintCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToPrintMedia(bk)
	}
	return ctx.OK(bs)
	// PrintsController_List: end_implement
}

// Show runs the show action.
func (c *PrintsController) Show(ctx *app.ShowPrintsContext) error {
	// PrintsController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetPrintByID(ctx.PrintID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get print`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToPrintMedia(b))
	// PrintsController_Show: end_implement
}

// Update runs the update action.
func (c *PrintsController) Update(ctx *app.UpdatePrintsContext) error {
	// PrintsController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdatePrint(ctx.PrintID, ctx.Payload.PrintName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update print`, `error`, err.Error())
		switch err {
		case model.ErrNotFound:
			return ctx.NotFound()
		case model.ErrDuplicateKey:
			return ctx.UnprocessableEntity()
		default:
			return ctx.InternalServerError()
		}
	}

	return ctx.NoContent()
	// PrintsController_Update: end_implement
}
