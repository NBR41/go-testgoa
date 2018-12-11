package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// AuthorsController implements the authors resource.
type AuthorsController struct {
	*goa.Controller
	fm Fmodeler
}

// NewAuthorsController creates a authors controller.
func NewAuthorsController(service *goa.Service, fm Fmodeler) *AuthorsController {
	return &AuthorsController{
		Controller: service.NewController("AuthorsController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *AuthorsController) Create(ctx *app.CreateAuthorsContext) error {
	// AuthorsController_Create: start_implement

	// Put your logic here
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertAuthor(ctx.Payload.AuthorName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert author`, `error`, err.Error())
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.AuthorsHref(b.ID))
	return ctx.Created()
	// AuthorsController_Create: end_implement
}

// Delete runs the delete action.
func (c *AuthorsController) Delete(ctx *app.DeleteAuthorsContext) error {
	// AuthorsController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteAuthor(ctx.AuthorID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete author`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// AuthorsController_Delete: end_implement
}

// List runs the list action.
func (c *AuthorsController) List(ctx *app.ListAuthorsContext) error {
	// AuthorsController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.GetAuthorList()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get author list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.AuthorCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToAuthorMedia(bk)
	}
	return ctx.OK(bs)
	// AuthorsController_List: end_implement
}

// Show runs the show action.
func (c *AuthorsController) Show(ctx *app.ShowAuthorsContext) error {
	// AuthorsController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetAuthorByID(ctx.AuthorID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get author`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToAuthorMedia(b))
	// AuthorsController_Show: end_implement
}

// Update runs the update action.
func (c *AuthorsController) Update(ctx *app.UpdateAuthorsContext) error {
	// AuthorsController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateAuthor(ctx.AuthorID, ctx.Payload.AuthorName)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update author`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// AuthorsController_Update: end_implement
}
