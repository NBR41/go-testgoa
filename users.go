package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmail"
	"github.com/NBR41/go-testgoa/appsec"
	"github.com/NBR41/go-testgoa/store"
	"github.com/goadesign/goa"
)

// ToUserMedia converts a user model into a user media type
func ToUserMedia(a *store.User) *app.User {
	return &app.User{
		Email:    a.Email,
		Href:     app.UsersHref(a.ID),
		ID:       int(a.ID),
		Nickname: a.Nickname,
	}
}

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service) *UsersController {
	return &UsersController{Controller: service.NewController("UsersController")}
}

// Create runs the create action.
func (c *UsersController) Create(ctx *app.CreateUsersContext) error {
	// UsersController_Create: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.InsertUser(ctx.Payload.Email, ctx.Payload.Nickname, ctx.Payload.Password)
	if err != nil {
		if err == store.ErrDuplicateKey {
			return ctx.UnprocessableEntity()
		}
		return ctx.InternalServerError()
	}

	token, err := appsec.GetValidationToken(u.ID, u.Email)
	if err == nil {
		_ = appmail.SendNewUserMail(u, token)
	}

	ctx.ResponseData.Header().Set("Location", app.UsersHref(u.ID))
	return ctx.Created()
	// UsersController_Create: end_implement
}

// Delete runs the delete action.
func (c *UsersController) Delete(ctx *app.DeleteUsersContext) error {
	// UsersController_Delete: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteUser(ctx.UserID)
	if err != nil {
		if err == store.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// UsersController_Delete: end_implement
}

// List runs the list action.
func (c *UsersController) List(ctx *app.ListUsersContext) error {
	// UsersController_List: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	users, err := m.GetUserList()
	if err != nil {
		return ctx.InternalServerError()
	}

	us := make(app.UserCollection, len(users))
	for i, u := range users {
		us[i] = ToUserMedia(&u)
	}
	return ctx.OK(us)
	// UsersController_List: end_implement
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	// UsersController_Show: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByID(ctx.UserID)
	if err != nil {
		if err == store.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(ToUserMedia(u))
	// UsersController_Show: end_implement
}

// Update runs the update action.
func (c *UsersController) Update(ctx *app.UpdateUsersContext) error {
	// UsersController_Update: start_implement

	// Put your logic here
	m, err := store.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateBook(ctx.UserID, ctx.Payload.Nickname)
	if err != nil {
		if err == store.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
	// UsersController_Update: end_implement
}
