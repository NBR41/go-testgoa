package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// EditorsController implements the editors resource.
type EditorsController struct {
	*goa.Controller
	fm Fmodeler
}

// NewEditorsController creates a editors controller.
func NewEditorsController(service *goa.Service, fm Fmodeler) *EditorsController {
	return &EditorsController{
		Controller: service.NewController("EditorsController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *EditorsController) Create(ctx *app.CreateEditorsContext) error {
	// EditorsController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertEditor(ctx.Payload.EditorName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert editor`, `error`, err.Error())
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.EditorsHref(b.ID))
	return ctx.Created()
	// EditorsController_Create: end_implement
}

// Delete runs the delete action.
func (c *EditorsController) Delete(ctx *app.DeleteEditorsContext) error {
	// EditorsController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteEditor(ctx.EditorID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete editor`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// EditorsController_Delete: end_implement
}

// List runs the list action.
func (c *EditorsController) List(ctx *app.ListEditorsContext) error {
	// EditorsController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.GetEditorList()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get editor list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.EditorCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToEditorMedia(bk)
	}
	return ctx.OK(bs)
	// EditorsController_List: end_implement
}

// Show runs the show action.
func (c *EditorsController) Show(ctx *app.ShowEditorsContext) error {
	// EditorsController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetEditorByID(ctx.EditorID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get editor`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToEditorMedia(b))
	// EditorsController_Show: end_implement
}

// Update runs the update action.
func (c *EditorsController) Update(ctx *app.UpdateEditorsContext) error {
	// EditorsController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateEditor(ctx.EditorID, ctx.Payload.EditorName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update editor`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// EditorsController_Update: end_implement
}
