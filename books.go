package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/store"
	"github.com/goadesign/goa"
)

// BooksController implements the books resource.
type BooksController struct {
	*goa.Controller
}

// ToBookMedia converts a book model into a book media type
func ToBookMedia(a *store.Book) *app.Book {
	return &app.Book{
		Href: app.BooksHref(a.ID),
		ID:   a.ID,
		Name: a.Name,
	}
}

// NewBooksController creates a books controller.
func NewBooksController(service *goa.Service) *BooksController {
	return &BooksController{Controller: service.NewController("BooksController")}
}

// Create runs the create action.
func (c *BooksController) Create(ctx *app.CreateBooksContext) error {
	// BooksController_Create: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {

	}

	b, err := m.InsertBook(*ctx.Payload.Name.BookName)
	if err != nil {

	}

	ctx.ResponseData.Header().Set("Location", app.BooksHref(b.ID))
	return ctx.Created()
	// BooksController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *BooksController) Delete(ctx *app.DeleteBooksContext) error {
	// BooksController_Delete: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {

	}

	err = m.DeleteBook(ctx.BookID)
	if err != nil {

	}

	return ctx.NoContent()
	// BooksController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *BooksController) List(ctx *app.ListBooksContext) error {
	// BooksController_List: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {

	}

	books, err := m.GetBookList()
	if err != nil {

	}

	bs := make(app.BookCollection, len(books))
	for i, bk := range books {
		bs[i] = ToBookMedia(&bk)
	}
	return ctx.OK(bs)

	// BooksController_List: end_implement
	res := app.BookCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *BooksController) Show(ctx *app.ShowBooksContext) error {
	// BooksController_Show: start_implement

	// Put your logic here

	// BooksController_Show: end_implement
	res := &app.Book{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *BooksController) Update(ctx *app.UpdateBooksContext) error {
	// BooksController_Update: start_implement

	// Put your logic here

	// BooksController_Update: end_implement
	return nil
}
