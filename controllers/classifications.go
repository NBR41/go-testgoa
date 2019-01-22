package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// ClassificationsController implements the classifications resource.
type ClassificationsController struct {
	*goa.Controller
	fm Fmodeler
}

// NewClassificationsController creates a classifications controller.
func NewClassificationsController(service *goa.Service, fm Fmodeler) *ClassificationsController {
	return &ClassificationsController{
		Controller: service.NewController("ClassificationsController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *ClassificationsController) Create(ctx *app.CreateClassificationsContext) error {
	// ClassificationsController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	_, err = m.InsertClassification(ctx.SeriesID, ctx.Payload.ClassID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert classification`, `error`, err.Error())
		switch err {
		case model.ErrDuplicateKey:
			return ctx.UnprocessableEntity()
		case model.ErrInvalidID:
			return ctx.UnprocessableEntity()
		default:
			return ctx.InternalServerError()
		}
	}

	ctx.ResponseData.Header().Set("Location", app.ClassificationsHref(ctx.SeriesID, ctx.Payload.ClassID))
	return ctx.Created()
	// ClassificationsController_Create: end_implement
}

// Delete runs the delete action.
func (c *ClassificationsController) Delete(ctx *app.DeleteClassificationsContext) error {
	// ClassificationsController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteClassification(ctx.SeriesID, ctx.ClassID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete classification`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// ClassificationsController_Delete: end_implement
}

// List runs the list action.
func (c *ClassificationsController) List(ctx *app.ListClassificationsContext) error {
	// ClassificationsController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListClassesByIDs(nil, nil, &ctx.SeriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get classification list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.ClassificationCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToClassificationMedia(ctx.SeriesID, bk)
	}
	return ctx.OK(bs)
	// ClassificationsController_List: end_implement
}

// Show runs the show action.
func (c *ClassificationsController) Show(ctx *app.ShowClassificationsContext) error {
	// ClassificationsController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	v, err := m.GetClassification(ctx.SeriesID, ctx.ClassID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get classification`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToClassificationMedia(ctx.SeriesID, v))
	// ClassificationsController_Show: end_implement
}
