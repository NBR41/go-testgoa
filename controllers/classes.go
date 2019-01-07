package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// ClassesController implements the genres resource.
type ClassesController struct {
	*goa.Controller
	fm Fmodeler
}

// NewClassesController creates a genres controller.
func NewClassesController(service *goa.Service, fm Fmodeler) *ClassesController {
	return &ClassesController{
		Controller: service.NewController("ClassesController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *ClassesController) Create(ctx *app.CreateClassesContext) error {
	// ClassesController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertClass(ctx.Payload.ClassName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert genre`, `error`, err.Error())
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.ClassesHref(b.ID))
	return ctx.Created()
	// ClassesController_Create: end_implement
}

// Delete runs the delete action.
func (c *ClassesController) Delete(ctx *app.DeleteClassesContext) error {
	// ClassesController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteClass(ctx.ClassID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete genre`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// ClassesController_Delete: end_implement
}

// List runs the list action.
func (c *ClassesController) List(ctx *app.ListClassesContext) error {
	// ClassesController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListClasses()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get genre list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.ClassCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToClassMedia(bk)
	}
	return ctx.OK(bs)
	// ClassesController_List: end_implement
}

// Show runs the show action.
func (c *ClassesController) Show(ctx *app.ShowClassesContext) error {
	// ClassesController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetClassByID(ctx.ClassID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get genre`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToClassMedia(b))
	// ClassesController_Show: end_implement
}

// Update runs the update action.
func (c *ClassesController) Update(ctx *app.UpdateClassesContext) error {
	// ClassesController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateClass(ctx.ClassID, ctx.Payload.ClassName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update genre`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// ClassesController_Update: end_implement
}
