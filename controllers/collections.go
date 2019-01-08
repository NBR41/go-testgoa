package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// CollectionsController implements the collections resource.
type CollectionsController struct {
	*goa.Controller
	fm Fmodeler
}

// NewCollectionsController creates a collections controller.
func NewCollectionsController(service *goa.Service, fm Fmodeler) *CollectionsController {
	return &CollectionsController{
		Controller: service.NewController("CollectionsController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *CollectionsController) Create(ctx *app.CreateCollectionsContext) error {
	// CollectionsController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	v, err := m.InsertCollection(ctx.Payload.CollectionName, ctx.Payload.EditorID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert collection`, `error`, err.Error())
		if err == model.ErrDuplicateKey || err == model.ErrNotFound {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.CollectionsHref(v.ID))
	return ctx.Created()
	// CollectionsController_Create: end_implement
}

// Delete runs the delete action.
func (c *CollectionsController) Delete(ctx *app.DeleteCollectionsContext) error {
	// CollectionsController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteCollection(ctx.CollectionID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete collection`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// CollectionsController_Delete: end_implement
}

// List runs the list action.
func (c *CollectionsController) List(ctx *app.ListCollectionsContext) error {
	// CollectionsController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListCollections()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get collection list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.CollectionCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToCollectionMedia(bk)
	}
	return ctx.OK(bs)
	// CollectionsController_List: end_implement
}

// Show runs the show action.
func (c *CollectionsController) Show(ctx *app.ShowCollectionsContext) error {
	// CollectionsController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetCollectionByID(ctx.CollectionID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get collection`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToCollectionMedia(b))
	// CollectionsController_Show: end_implement
}

// Update runs the update action.
func (c *CollectionsController) Update(ctx *app.UpdateCollectionsContext) error {
	// CollectionsController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateCollection(ctx.CollectionID, ctx.Payload.CollectionName, ctx.Payload.EditorID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update collection`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// CollectionsController_Update: end_implement
}
