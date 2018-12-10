package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// EditionTypesController implements the edition_types resource.
type EditionTypesController struct {
	*goa.Controller
	fm Fmodeler
}

// NewEditionTypesController creates a edition_types controller.
func NewEditionTypesController(service *goa.Service, fm Fmodeler) *EditionTypesController {
	return &EditionTypesController{
		Controller: service.NewController("EditionTypesController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *EditionTypesController) Create(ctx *app.CreateEditionTypesContext) error {
	// EditionTypesController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertEditionType(ctx.Payload.Name)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert edition type`, `error`, err)
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.EditionTypesHref(b.ID))
	return ctx.Created()
	// EditionTypesController_Create: end_implement
}

// Delete runs the delete action.
func (c *EditionTypesController) Delete(ctx *app.DeleteEditionTypesContext) error {
	// EditionTypesController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteEditionType(ctx.EditionTypeID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete edition type`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// EditionTypesController_Delete: end_implement
}

// List runs the list action.
func (c *EditionTypesController) List(ctx *app.ListEditionTypesContext) error {
	// EditionTypesController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.GetEditionTypeList()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get edition type list`, `error`, err)
		return ctx.InternalServerError()
	}

	bs := make(app.EditiontypeCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToEditionTypeMedia(bk)
	}
	return ctx.OK(bs)
	// EditionTypesController_List: end_implement
}

// Show runs the show action.
func (c *EditionTypesController) Show(ctx *app.ShowEditionTypesContext) error {
	// EditionTypesController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetEditionTypeByID(ctx.EditionTypeID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get edition type`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToEditionTypeMedia(b))
	// EditionTypesController_Show: end_implement
}

// Update runs the update action.
func (c *EditionTypesController) Update(ctx *app.UpdateEditionTypesContext) error {
	// EditionTypesController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateEditionType(ctx.EditionTypeID, ctx.Payload.Name)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update edition type`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// EditionTypesController_Update: end_implement
}
