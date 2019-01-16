package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/api"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// OwnershipsController implements the ownerships resource.
type OwnershipsController struct {
	*goa.Controller
	fm  Fmodeler
	api APIHelper
}

// NewOwnershipsController creates a ownerships controller.
func NewOwnershipsController(service *goa.Service, fm Fmodeler, api APIHelper) *OwnershipsController {
	return &OwnershipsController{
		Controller: service.NewController("OwnershipsController"),
		fm:         fm,
		api:        api,
	}
}

// Create runs the create action.
func (c *OwnershipsController) Create(ctx *app.CreateOwnershipsContext) error {
	// OwnershipsController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	o, err := m.InsertOwnership(ctx.UserID, ctx.Payload.BookID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to insert ownership`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		switch err {
		case model.ErrDuplicateKey:
			return ctx.UnprocessableEntity()
		case model.ErrInvalidID:
			return ctx.UnprocessableEntity()
		default:
			return ctx.InternalServerError()
		}
	}

	ctx.ResponseData.Header().Set("Location", app.OwnershipsHref(ctx.UserID, o.BookID))
	return ctx.Created()
	// OwnershipsController_Create: end_implement
}

// Add runs the add action.
func (c *OwnershipsController) Add(ctx *app.AddOwnershipsContext) error {
	// OwnershipsController_Add: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	book, err := m.GetBookByISBN(ctx.Payload.BookIsbn)
	if err != nil {
		if err == model.ErrNotFound {
			// Get the book name
			var bookName string
			bookName, err = c.api.GetBookName(ctx.Payload.BookIsbn)
			if err != nil {
				goa.ContextLogger(ctx).Error(`unable to get book name`, `error`, err.Error())
				if err == api.ErrNoResult {
					return ctx.UnprocessableEntity(err)
				}
				return ctx.InternalServerError()
			}

			//TODO set series id
			var seriesID int

			// insert the new book
			book, err = m.InsertBook(ctx.Payload.BookIsbn, bookName, seriesID)
			if err != nil {
				goa.ContextLogger(ctx).Error(`unable to insert book`, `error`, err.Error())
				switch err {
				case model.ErrInvalidID:
					return ctx.UnprocessableEntity(err)
				case model.ErrDuplicateKey:
					return ctx.UnprocessableEntity(err)
				default:
					return ctx.InternalServerError()
				}
			}
		} else {
			goa.ContextLogger(ctx).Error(`unable to get book by isbn`, `error`, err.Error())
			return ctx.InternalServerError()
		}
	}

	_, err = m.InsertOwnership(ctx.UserID, int(book.ID))
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to insert ownership`, `error`, err.Error())
		switch err {
		case model.ErrNotFound:
			return ctx.NotFound()
		case model.ErrInvalidID:
			return ctx.UnprocessableEntity(err)
		case model.ErrDuplicateKey:
			return ctx.UnprocessableEntity(err)
		default:
			return ctx.InternalServerError()
		}
	}

	ctx.ResponseData.Header().Set("Location", app.OwnershipsHref(ctx.UserID, book.ID))
	return ctx.Created()
	// OwnershipsController_Add: end_implement
}

// Delete runs the delete action.
func (c *OwnershipsController) Delete(ctx *app.DeleteOwnershipsContext) error {
	// OwnershipsController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteOwnership(ctx.UserID, ctx.BookID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to delete ownership`, `error`, err.Error())
		if err == model.ErrNotFound {
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
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	books, err := m.ListOwnershipsByUserID(ctx.UserID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get ownership list`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	bs := make(app.OwnershipCollection, len(books))
	for i, bk := range books {
		bs[i] = convert.ToOwnershipMedia(bk)
	}
	return ctx.OK(bs)
	// OwnershipsController_List: end_implement
}

// Show runs the show action.
func (c *OwnershipsController) Show(ctx *app.ShowOwnershipsContext) error {
	// OwnershipsController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	o, err := m.GetOwnership(ctx.UserID, ctx.BookID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get ownership`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToOwnershipMedia(o))
	// OwnershipsController_Show: end_implement
}
