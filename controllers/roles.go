package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// RolesController implements the roles resource.
type RolesController struct {
	*goa.Controller
	fm Fmodeler
}

// NewRolesController creates a roles controller.
func NewRolesController(service *goa.Service, fm Fmodeler) *RolesController {
	return &RolesController{
		Controller: service.NewController("RolesController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *RolesController) Create(ctx *app.CreateRolesContext) error {
	// RolesController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertRole(ctx.Payload.Name)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert role`, `error`, err)
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.RolesHref(b.ID))
	return ctx.Created()
	// RolesController_Create: end_implement
}

// Delete runs the delete action.
func (c *RolesController) Delete(ctx *app.DeleteRolesContext) error {
	// RolesController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteRole(ctx.RoleID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete role`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// RolesController_Delete: end_implement
}

// List runs the list action.
func (c *RolesController) List(ctx *app.ListRolesContext) error {
	// RolesController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.GetRoleList()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get role list`, `error`, err)
		return ctx.InternalServerError()
	}

	bs := make(app.RoleCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToRoleMedia(bk)
	}
	return ctx.OK(bs)
	// RolesController_List: end_implement
}

// Show runs the show action.
func (c *RolesController) Show(ctx *app.ShowRolesContext) error {
	// RolesController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetRoleByID(ctx.RoleID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get role`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToRoleMedia(b))
	// RolesController_Show: end_implement
}

// Update runs the update action.
func (c *RolesController) Update(ctx *app.UpdateRolesContext) error {
	// RolesController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateRole(ctx.RoleID, ctx.Payload.Name)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update role`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// RolesController_Update: end_implement
}
