package main

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/appmail"
	"github.com/NBR41/go-testgoa/appmodel"
	"github.com/NBR41/go-testgoa/appsec"
	"github.com/goadesign/goa"
)

// ToUserMedia converts a user model into a user media type
func ToUserMedia(a *appmodel.User) *app.User {
	return &app.User{
		Email:       a.Email,
		Href:        app.UsersHref(a.ID),
		ID:          int(a.ID),
		Nickname:    a.Nickname,
		IsAdmin:     a.IsAdmin,
		IsValidated: a.IsValidated,
	}
}

// ToUserTinyMedia converts a user model into a user media type
func ToUserTinyMedia(a *appmodel.User) *app.UserTiny {
	return &app.UserTiny{
		Href:     app.UsersHref(a.ID),
		ID:       int(a.ID),
		Nickname: a.Nickname,
	}
}

var errDuplicateKey = goa.NewErrorClass("duplicate_key", 422)

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
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.InsertUser(ctx.Payload.Email, ctx.Payload.Nickname, ctx.Payload.Password)
	if err != nil {
		switch err {
		case appmodel.ErrDuplicateKey:
			return errDuplicateKey(err)
		case appmodel.ErrDuplicateEmail:
			return errDuplicateKey(err)
		case appmodel.ErrDuplicateNickname:
			return errDuplicateKey(err)
		default:
			return ctx.InternalServerError()
		}
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
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteUser(ctx.UserID)
	if err != nil {
		if err == appmodel.ErrNotFound {
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
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	var users []appmodel.User
	if ctx.Email != nil || ctx.Nickname != nil {
		var user *appmodel.User
		switch {
		case ctx.Email != nil && ctx.Nickname == nil:
			user, err = m.GetUserByEmail(*ctx.Email)
			if err != nil && err != appmodel.ErrNotFound {
				return ctx.InternalServerError()
			}
		case ctx.Email == nil && ctx.Nickname != nil:
			user, err = m.GetUserByNickname(*ctx.Nickname)
			if err != nil && err != appmodel.ErrNotFound {
				return ctx.InternalServerError()
			}
		default:
			user, err = m.GetUserByEmailOrNickname(*ctx.Email, *ctx.Nickname)
			if err != nil && err != appmodel.ErrNotFound {
				return ctx.InternalServerError()
			}
		case ctx.Email != nil && ctx.Nickname != nil:
		}
		if err == nil {
			users = append(users, *user)
		}
	} else {
		users, err = m.GetUserList()
		if err != nil {
			return ctx.InternalServerError()
		}
	}

	us := make(app.UserTinyCollection, len(users))
	for i, u := range users {
		us[i] = ToUserTinyMedia(&u)
	}
	return ctx.OKTiny(us)
	// UsersController_List: end_implement
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	// UsersController_Show: start_implement

	// Put your logic here
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByID(ctx.UserID)
	if err != nil {
		if err == appmodel.ErrNotFound {
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
	m, err := appmodel.GetModeler()
	if err != nil {
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()
	err = m.UpdateUserNickname(ctx.UserID, ctx.Payload.Nickname)
	if err != nil {
		switch {
		case err == appmodel.ErrDuplicateKey:
			return ctx.UnprocessableEntity()
		case err == appmodel.ErrNotFound:
			return ctx.NotFound()
		default:
			return ctx.InternalServerError()
		}
	}
	return ctx.NoContent()
	// UsersController_Update: end_implement
}
