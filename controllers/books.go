package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// BooksController implements the books resource.
type BooksController struct {
	*goa.Controller
	fm  Fmodeler
	api APIHelper
}

// NewBooksController creates a books controller.
func NewBooksController(service *goa.Service, fm Fmodeler, api APIHelper) *BooksController {
	return &BooksController{
		Controller: service.NewController("BooksController"),
		fm:         fm,
		api:        api,
	}
}

// Create runs the create action.
func (c *BooksController) Create(ctx *app.CreateBooksContext) error {
	// BooksController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertBook(ctx.Payload.BookIsbn, ctx.Payload.BookName, ctx.Payload.SeriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to insert book`, `error`, err.Error())
		switch err {
		case model.ErrDuplicateKey:
			return ctx.UnprocessableEntity()
		case model.ErrInvalidID:
			return ctx.UnprocessableEntity()
		default:
			return ctx.InternalServerError()
		}
	}

	ctx.ResponseData.Header().Set("Location", app.BooksHref(b.ID))
	return ctx.Created()
	// BooksController_Create: end_implement
}

// Delete runs the delete action.
func (c *BooksController) Delete(ctx *app.DeleteBooksContext) error {
	// BooksController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteBook(ctx.BookID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to delete book`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// BooksController_Delete: end_implement
}

// List runs the list action.
func (c *BooksController) List(ctx *app.ListBooksContext) error {
	// BooksController_List: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	books, err := m.ListBooksByIDs(nil, nil, nil, nil)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get book list`, `error`, err.Error())
		return ctx.InternalServerError()
	}

	bs := make(app.BookCollection, len(books))
	for i, bk := range books {
		bs[i] = convert.ToBookMedia(bk)
	}
	return ctx.OK(bs)
	// BooksController_List: end_implement
}

// Show runs the show action.
func (c *BooksController) Show(ctx *app.ShowBooksContext) error {
	// BooksController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetBookByID(ctx.BookID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get book`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToBookMedia(b))
	// BooksController_Show: end_implement
}

// Update runs the update action.
func (c *BooksController) Update(ctx *app.UpdateBooksContext) error {
	// BooksController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateBook(ctx.BookID, ctx.Payload.BookName, ctx.Payload.SeriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to update book`, `error`, err.Error())
		switch err {
		case model.ErrNotFound:
			return ctx.NotFound()
		case model.ErrInvalidID:
			return ctx.UnprocessableEntity()
		case model.ErrDuplicateKey:
			return ctx.UnprocessableEntity()
		default:
			return ctx.InternalServerError()
		}
	}

	return ctx.NoContent()
	// BooksController_Update: end_implement
}

// IsbnSearch runs the isbn search action.
func (c *BooksController) IsbnSearch(ctx *app.IsbnSearchBooksContext) error {
	// BooksController_IsbnSearch: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetBookDetail(ctx.BookIsbn)
	if err != nil {
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}

		ab, err := c.api.GetBookDetail(ctx.BookIsbn)
		if err != nil {
			return ctx.InternalServerError()
		}

		b, err = m.InsertBookDetail(ctx.BookIsbn, ab)
		if err != nil {
			return ctx.InternalServerError()
		}
	}
	return ctx.OK(convert.ToBookDetailMedia(b))
	// BooksController_IsbnSearch: end_implement
}
