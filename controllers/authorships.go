package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// AuthorshipsController implements the authorships resource.
type AuthorshipsController struct {
	*goa.Controller
	fm Fmodeler
}

// NewAuthorshipsController creates a authorships controller.
func NewAuthorshipsController(service *goa.Service, fm Fmodeler) *AuthorshipsController {
	return &AuthorshipsController{
		Controller: service.NewController("AuthorshipsController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *AuthorshipsController) Create(ctx *app.CreateAuthorshipsContext) error {
	// AuthorshipsController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	v, err := m.InsertAuthorship(ctx.Payload.AuthorID, ctx.Payload.BookID, ctx.Payload.RoleID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert authorship`, `error`, err.Error())
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.AuthorshipsHref(v.ID))
	return ctx.Created()
	// AuthorshipsController_Create: end_implement
}

// Delete runs the delete action.
func (c *AuthorshipsController) Delete(ctx *app.DeleteAuthorshipsContext) error {
	// AuthorshipsController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteAuthorship(ctx.AuthorshipID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete authorship`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// AuthorshipsController_Delete: end_implement
}

// List runs the list action.
func (c *AuthorshipsController) List(ctx *app.ListAuthorshipsContext) error {
	// AuthorshipsController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListAuthorships()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get authorship list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.AuthorshipCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToAuthorshipMedia(bk)
	}
	return ctx.OK(bs)
	// AuthorshipsController_List: end_implement
}

// Show runs the show action.
func (c *AuthorshipsController) Show(ctx *app.ShowAuthorshipsContext) error {
	// AuthorshipsController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	v, err := m.GetAuthorshipByID(ctx.AuthorshipID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get authorship`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToAuthorshipMedia(v))
	// AuthorshipsController_Show: end_implement
}
