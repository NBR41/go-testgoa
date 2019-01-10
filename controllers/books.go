package controllers

import (
	"context"
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// BooksController implements the books resource.
type BooksController struct {
	*goa.Controller
	fm Fmodeler
}

// NewBooksController creates a books controller.
func NewBooksController(service *goa.Service, fm Fmodeler) *BooksController {
	return &BooksController{
		Controller: service.NewController("BooksController"),
		fm:         fm,
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
		if err == model.ErrDuplicateKey || err == model.ErrNotFound {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
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

	books, err := m.ListBooks()
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
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// BooksController_Update: end_implement
}

type booksResponse interface {
	OK(r app.BookCollection) error
	NotFound() error
	InternalServerError() error
	ServiceUnavailable() error
}

func listBooks(ctx context.Context, fm Fmodeler, rCtx booksResponse, collectionID, printID, seriesID *int) error {
	m, err := fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return rCtx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	list, err := m.ListBooksByIDs(collectionID, printID, seriesID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`failed to get book list`, `error`, err.Error())
		if err == model.ErrNotFound {
			return rCtx.NotFound()
		}
		return rCtx.InternalServerError()
	}
	bs := make(app.BookCollection, len(list))
	for i, bk := range list {
		bs[i] = convert.ToBookMedia(bk)
	}
	return rCtx.OK(bs)
}
