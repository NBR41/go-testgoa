package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// GenresController implements the genres resource.
type GenresController struct {
	*goa.Controller
	fm Fmodeler
}

// NewGenresController creates a genres controller.
func NewGenresController(service *goa.Service, fm Fmodeler) *GenresController {
	return &GenresController{
		Controller: service.NewController("GenresController"),
		fm:         fm,
	}
}

// Create runs the create action.
func (c *GenresController) Create(ctx *app.CreateGenresContext) error {
	// GenresController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertGenre(ctx.Payload.Name)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert genre`, `error`, err)
		if err == model.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.GenresHref(b.ID))
	return ctx.Created()
	// GenresController_Create: end_implement
}

// Delete runs the delete action.
func (c *GenresController) Delete(ctx *app.DeleteGenresContext) error {
	// GenresController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteGenre(ctx.GenreID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete genre`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// GenresController_Delete: end_implement
}

// List runs the list action.
func (c *GenresController) List(ctx *app.ListGenresContext) error {
	// GenresController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.GetGenreList()
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get genre list`, `error`, err)
		return ctx.InternalServerError()
	}

	bs := make(app.GenreCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToGenreMedia(bk)
	}
	return ctx.OK(bs)
	// GenresController_List: end_implement
}

// Show runs the show action.
func (c *GenresController) Show(ctx *app.ShowGenresContext) error {
	// GenresController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetGenreByID(ctx.GenreID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get genre`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToGenreMedia(b))
	// GenresController_Show: end_implement
}

// Update runs the update action.
func (c *GenresController) Update(ctx *app.UpdateGenresContext) error {
	// GenresController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateGenre(ctx.GenreID, ctx.Payload.Name)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update genre`, `error`, err)
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// GenresController_Update: end_implement
}
