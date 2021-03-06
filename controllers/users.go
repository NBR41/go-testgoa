package controllers

import (
	"github.com/NBR41/go-testgoa/app"
	"github.com/NBR41/go-testgoa/internal/convert"
	"github.com/NBR41/go-testgoa/internal/model"
	"github.com/goadesign/goa"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	fm   Fmodeler
	tok  TokenHelper
	mail MailSender
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, fm Fmodeler, tok TokenHelper, mail MailSender) *UsersController {
	return &UsersController{
		Controller: service.NewController("UsersController"),
		fm:         fm,
		tok:        tok,
		mail:       mail,
	}
}

// Create runs the create action.
func (c *UsersController) Create(ctx *app.CreateUsersContext) error {
	// UsersController_Create: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.InsertUser(ctx.Payload.Email, ctx.Payload.Nickname, ctx.Payload.Password)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to insert user`, `error`, err.Error())
		if err == model.ErrDuplicateKey || err == model.ErrDuplicateEmail || err == model.ErrDuplicateNickname {
			return ctx.UnprocessableEntity(ErrUnprocessableEntity(err))
		}
		return ctx.InternalServerError()
	}

	token, err := c.tok.GetValidationToken(u.ID, u.Email)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get validation token`, `error`, err.Error())
	} else {
		err = c.mail.SendNewUserMail(u, token)
		if err != nil {
			goa.ContextLogger(ctx).Error(`unable to send user creation email`, `error`, err.Error())
		}
	}

	ctx.ResponseData.Header().Set("Location", app.UsersHref(u.ID))
	return ctx.Created()
	// UsersController_Create: end_implement
}

// Delete runs the delete action.
func (c *UsersController) Delete(ctx *app.DeleteUsersContext) error {
	// UsersController_Delete: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.DeleteUser(ctx.UserID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to delete user`, `error`, err.Error())
		if err == model.ErrNotFound {
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
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	var users []*model.User
	if ctx.Email != nil || ctx.Nickname != nil {
		var user *model.User
		switch {
		case ctx.Email != nil && ctx.Nickname == nil:
			user, err = m.GetUserByEmail(*ctx.Email)
		case ctx.Email == nil && ctx.Nickname != nil:
			user, err = m.GetUserByNickname(*ctx.Nickname)
		default:
			user, err = m.GetUserByEmailOrNickname(*ctx.Email, *ctx.Nickname)
		}
		if err != nil && err != model.ErrNotFound {
			goa.ContextLogger(ctx).Error(`unable to get user`, `error`, err.Error())
			return ctx.InternalServerError()
		}
		if err == nil {
			users = append(users, user)
		}
	} else {
		users, err = m.ListUsers()
		if err != nil {
			goa.ContextLogger(ctx).Error(`unable to get user list`, `error`, err.Error())
			return ctx.InternalServerError()
		}
	}

	us := make(app.UserCollection, len(users))
	for i, u := range users {
		us[i] = convert.ToUserMedia(u)
	}
	return ctx.OK(us)
	// UsersController_List: end_implement
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	// UsersController_Show: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	u, err := m.GetUserByID(ctx.UserID)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get user`, `error`, err.Error())
		if err == model.ErrNotFound {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(convert.ToUserMedia(u))
	// UsersController_Show: end_implement
}

// Update runs the update action.
func (c *UsersController) Update(ctx *app.UpdateUsersContext) error {
	// UsersController_Update: start_implement
	m, err := c.fm()
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to get model`, `error`, err.Error())
		return ctx.ServiceUnavailable()
	}
	defer func() { m.Close() }()

	err = m.UpdateUserNickname(ctx.UserID, ctx.Payload.Nickname)
	if err != nil {
		goa.ContextLogger(ctx).Error(`unable to update user`, `error`, err.Error())
		switch {
		case err == model.ErrDuplicateNickname:
			return ctx.UnprocessableEntity(ErrUnprocessableEntity(err))
		case err == model.ErrNotFound:
			return ctx.NotFound()
		default:
			return ctx.InternalServerError()
		}
	}
	return ctx.NoContent()
	// UsersController_Update: end_implement
}
