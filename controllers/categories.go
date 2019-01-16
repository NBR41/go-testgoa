package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// CategoriesController implements the categories resource.
type CategoriesController struct {
	*goa.Controller
	fm Fmodeler
}

// NewCategoriesController creates a categories controller.
func NewCategoriesController(service *goa.Service, fm Fmodeler) *CategoriesController {
	return &CategoriesController{
		Controller: service.NewController("CategoriesController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *CategoriesController) Create(ctx *app.CreateCategoriesContext) error {
	// CategoriesController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertCategory(ctx.Payload.CategoryName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert category`, `error`, err.Error())
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.CategoriesHref(b.ID))
	return ctx.Created()
	// CategoriesController_Create: end_implement
}

// Delete runs the delete action.
func (c *CategoriesController) Delete(ctx *app.DeleteCategoriesContext) error {
	// CategoriesController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteCategory(ctx.CategoryID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete category`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// CategoriesController_Delete: end_implement
}

// List runs the list action.
func (c *CategoriesController) List(ctx *app.ListCategoriesContext) error {
	// CategoriesController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListCategories()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get category list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.CategoryCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToCategoryMedia(bk)
	}
	return ctx.OK(bs)
	// CategoriesController_List: end_implement
}

// Show runs the show action.
func (c *CategoriesController) Show(ctx *app.ShowCategoriesContext) error {
	// CategoriesController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetCategoryByID(ctx.CategoryID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get category`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToCategoryMedia(b))
	// CategoriesController_Show: end_implement
}

// Update runs the update action.
func (c *CategoriesController) Update(ctx *app.UpdateCategoriesContext) error {
	// CategoriesController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateCategory(ctx.CategoryID, ctx.Payload.CategoryName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update category`, `error`, err.Error())
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
	// CategoriesController_Update: end_implement
}
