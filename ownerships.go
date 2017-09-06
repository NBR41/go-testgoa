package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appapi"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/goadesign/goa"
)

// ToOwnershipMedia converts a book model into a book media type
func ToOwnershipMedia(a *appmodel.Ownership) *app.Ownership {
	return &app.Ownership{
		Book:   ToBookMedia(a.Book),
		BookID: int(a.BookID),
		Href:   app.OwnershipsHref(a.UserID, a.BookID),
		UserID: int(a.UserID),
	}
}

// OwnershipsController implements the ownerships resource.
type OwnershipsController struct {
	*goa.Controller
}

// NewOwnershipsController creates a ownerships controller.
func NewOwnershipsController(service *goa.Service) *OwnershipsController {
	return &OwnershipsController{Controller: service.NewController("OwnershipsController")}
}

// Create runs the create action.
func (c *OwnershipsController) Create(ctx *app.CreateOwnershipsContext) error {
	// OwnershipsController_Create: start_implement

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	o, err := m.InsertOwnership(ctx.UserID, ctx.Payload.BookID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to insert ownership`, `error`, err)
		if err == appmodel.ErrNotFound {
			return ctx.NotFound()
		}
		if err == appmodel.ErrDuplicateKey || err == appmodel.ErrInvalidID {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.OwnershipsHref(ctx.UserID, o.BookID))
	return ctx.Created()
	// OwnershipsController_Create: end_implement
}

// Add runs the create action.
func (c *OwnershipsController) Add(ctx *app.AddOwnershipsContext) error {
	// OwnershipsController_Create: start_implement
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	book, err := m.GetBookByISBN(ctx.Payload.Isbn)
	if err != nil {
		if err == appmodel.ErrNotFound {
			// Get the book name
			var bookName string
			bookName, err = appapi.GetBookName(ctx.Payload.Isbn)
			if err != nil {
				goa.ContextLogger(ctx).Error(`unable to get book name`, `error`, err)
				return ctx.InternalServerError()
			}

			// insert the new book
			book, err = m.InsertBook(ctx.Payload.Isbn, bookName)
			if err != nil {
				goa.ContextLogger(ctx).Error(`unable to insert book`, `error`, err)
				if err == appmodel.ErrDuplicateKey {
					return ctx.UnprocessableEntity()
				}
				return ctx.InternalServerError()
			}
		} else {
			goa.ContextLogger(ctx).Error(`unable to get book by isbn`, `error`, err)
			return ctx.InternalServerError()
		}
	}

	o, err := m.InsertOwnership(ctx.UserID, int(book.ID))
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to insert ownership`, `error`, err)
		if err == appmodel.ErrNotFound {
			return ctx.NotFound()
		}
		if err == appmodel.ErrDuplicateKey || err == appmodel.ErrInvalidID {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	ctx.ResponseData.Header().Set("Location", app.OwnershipsHref(ctx.UserID, o.BookID))
	return ctx.Created()
	// OwnershipsController_Create: end_implement
}

// Delete runs the delete action.
func (c *OwnershipsController) Delete(ctx *app.DeleteOwnershipsContext) error {
	// OwnershipsController_Delete: start_implement
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteOwnership(ctx.UserID, ctx.BookID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to delete ownership`, `error`, err)
		if err == appmodel.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// OwnershipsController_Delete: end_implement
}

// List runs the list action.
func (c *OwnershipsController) List(ctx *app.ListOwnershipsContext) error {
	// OwnershipsController_List: start_implement
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	books, err := m.GetOwnershipList(ctx.UserID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get ownership list`, `error`, err)
		return ctx.InternalServerError()
	}

	bs := make(app.OwnershipCollection, len(books))
	for i, bk := range books {
		bs[i] = ToOwnershipMedia(&bk)
	}
	return ctx.OK(bs)
	// OwnershipsController_List: end_implement
}

// Show runs the show action.
func (c *OwnershipsController) Show(ctx *app.ShowOwnershipsContext) error {
	// OwnershipsController_Show: start_implement
	m, err := appmodel.GetModeler()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err)
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	o, err := m.GetOwnership(ctx.UserID, ctx.BookID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get ownership`, `error`, err)
		if err == appmodel.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(ToOwnershipMedia(o))
	// OwnershipsController_Show: end_implement
}
