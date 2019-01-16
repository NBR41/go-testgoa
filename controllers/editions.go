package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// EditionsController implements the editions resource.
type EditionsController struct {
	*goa.Controller
	fm Fmodeler
}

// NewEditionsController creates a editions controller.
func NewEditionsController(service *goa.Service, fm Fmodeler) *EditionsController {
	return &EditionsController{
		Controller: service.NewController("EditionsController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *EditionsController) Create(ctx *app.CreateEditionsContext) error {
	// EditionsController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	v, err := m.InsertEdition(ctx.Payload.BookID, ctx.Payload.CollectionID, ctx.Payload.PrintID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert edition`, `error`, err.Error())
		switch err {
		case model.ErrDuplicateKey:
			return ctx.UnprocessableEntity()
		case model.ErrInvalidID:
			return ctx.UnprocessableEntity()
		default:
			return ctx.InternalServerError()
		}
	}

	ctx.ResponseData.Header().Set("Location", app.EditionsHref(v.ID))
	return ctx.Created()
	// EditionsController_Create: end_implement
}

// Delete runs the delete action.
func (c *EditionsController) Delete(ctx *app.DeleteEditionsContext) error {
	// EditionsController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteEdition(ctx.EditionID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete edition`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// EditionsController_Delete: end_implement
}

// List runs the list action.
func (c *EditionsController) List(ctx *app.ListEditionsContext) error {
	// EditionsController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListEditions()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get edition list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.EditionCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToEditionMedia(bk)
	}
	return ctx.OK(bs)
	// EditionsController_List: end_implement
}

// Show runs the show action.
func (c *EditionsController) Show(ctx *app.ShowEditionsContext) error {
	// EditionsController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	v, err := m.GetEditionByID(ctx.EditionID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get edition`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToEditionMedia(v))
	// EditionsController_Show: end_implement
}
