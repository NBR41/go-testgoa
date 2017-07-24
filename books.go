package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/goadesign/goa"
)

// ToBookMedia converts a book model into a book media type
func ToBookMedia(a *appmodel.Book) *app.Book {
	return &app.Book{
		Href: app.BooksHref(a.ID),
		ID:   int(a.ID),
		Name: a.Name,
	}
}

// ToBookLinkMedia converts a book model into a book link media type
func ToBookLinkMedia(a *appmodel.Book) *app.BookLink {
	return &app.BookLink{
		Href: app.BooksHref(a.ID),
		ID:   int(a.ID),
		Name: a.Name,
	}
}

// BooksController implements the books resource.
type BooksController struct {
	*goa.Controller
}

// NewBooksController creates a books controller.
func NewBooksController(service *goa.Service) *BooksController {
	return &BooksController{Controller: service.NewController("BooksController")}
}

// Create runs the create action.
func (c *BooksController) Create(ctx *app.CreateBooksContext) error {
	// BooksController_Create: start_implement

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.InsertBook(ctx.Payload.Name)
	if err != nil {
		if err == appmodel.ErrDuplicateKey {
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

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteBook(ctx.BookID)
	if err != nil {
		if err == appmodel.ErrNotFound {
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

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	books, err := m.GetBookList()
	if err != nil {
		return ctx.InternalServerError()
	}

	bs := make(app.BookCollection, len(books))
	for i, bk := range books {
		bs[i] = ToBookMedia(&bk)
	}
	return ctx.OK(bs)
	// BooksController_List: end_implement
}

// Show runs the show action.
func (c *BooksController) Show(ctx *app.ShowBooksContext) error {
	// BooksController_Show: start_implement

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	b, err := m.GetBookByID(ctx.BookID)
	if err != nil {
		if err == appmodel.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(ToBookMedia(b))
	// BooksController_Show: end_implement
}

// Update runs the update action.
func (c *BooksController) Update(ctx *app.UpdateBooksContext) error {
	// BooksController_Update: start_implement

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateBook(ctx.BookID, ctx.Payload.Name)
	if err != nil {
		if err == appmodel.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// BooksController_Update: end_implement
}
